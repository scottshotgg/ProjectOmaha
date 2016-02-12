/*
 *
 *  Author: Matthew Kramer
 *  Revision 1.0 (8 February 2016)
 *
 *  THIS CODE IS CONFIDENTIAL AND FOR THE EYES OF A SELECT FEW
 *
 * 	CLA_Funcs.c version 1.0
 *
 *	Source code which contains CLA initialization functions as well as the ISR handlers for when the CLA completes a task
 *
 *	The CPU will tell the CLA to execute a task (via interrupt), the CLA will then execute the task while the CPU continues to work, once the CLA is finished, it throws an interrupt back at the cpu
 *
 *	Cla1Task* 		functions contain the actual CLA program
 *	cla1_task*_isr	functions contain the program that the cpu will execute after the cla is complete with its task
 *
 *	The CLA programs are located in eq.cla
 *
 *	Three Tasks are written for the CLA. Their priority decreases as the Task number increases (Task 1 highest priority)
 *	Task 1 is used when the user wants to modify the equalizer. It adjusts the BIquad filter structure that corresponds to the appropiate band
 *	Task 2 is used for executing the filter.
 *	Task 3 is used for initializing CLA variables. All biquad structures have a gain of 0db initially. Cla variables cannot be initialized at compile time.
 *
 *	Task 1's CPU ISR does nothing
 *	Task 2's CPU ISR does volume adjustment, mixes the soundmasking with the music/paging, Muxes the music and paging signal and calls fade time functions
 *	Task 3's CPU ISR does nothing
 *
 *	fade and unfade essentially bring the soundmasking volume down to a specified percentage of current volume and back up to the volume level before fading
 *
 *
 *
 *	Cla_init initializes CLA
 *
 *
*/

#include "Cla_Shared.h"
#include "constants.h"
#include "timers.h"
extern uint16_t Cla1funcsRunStart, Cla1funcsLoadStart, Cla1funcsLoadSize;
extern uint16_t Cla1mathTablesRunStart, Cla1mathTablesLoadStart, Cla1mathTablesLoadSize;
extern uint16_t Cla1Prog_Start;
//external variables
extern int 		scalar;						// used for normalizing pwm		(declared in main.c)
extern int		ADCReading_music;			// adc sample from music channel (declared in main.c)
extern int 		ADCReading_paging;			// adc sample from paging channel (decaled in main.c)
extern float 	Sfactor;					// Soundmasking percentage (maximum value of 1) (declared in main.c) (initial value of 1)
extern float	Mfactor;					// Music percentage (maximum value of 1) (declared in main.c) 		 (initial value of 0)
extern float 	Vfactor;					// over all volume  (maximum value of 1) (declared in main.c)		 (initial value of .1)
extern float	initialS;					// used for fading
extern float 	fadeTo;						// used for fading
extern int		count;						// used for fading
extern int		paging;						// used as a boolean. If true (!= 0), the paging is mixed with the sound masking. if false (== 0), the music is mixed. Also causes fading to be called
extern int		hasFaded;					// used for fading
extern int		fadeConstant;				// used for fading


//internal variables
int		   		hrpwm;						// variable that is send the hrpwm duty cycle register. By carfully adjusting the values that go to it, a very
											// linear sweep of the duty cycle is achievable when increasing masking from 0 to 4000
int		   		pwm;						// variable that is sent to pwm duty cycle register. It the mixed signal normalized by the scalar.
int				music;					    // used for mixing (argument 1)
int				masking;					// used for mixing (argument 2)
int				mixed;						// used for mixing (final output)




void fade(float toValue, int rate);			//these functions are used locally here and are used for fading and unfading the masking volume when a page is sent
void unFade(float toValue, int rate);



void Cla_Init(){

	EALLOW;
	memcpy(&Cla1funcsRunStart, &Cla1funcsLoadStart, (Uint32)&Cla1funcsLoadSize);					//CLA code is written in flash. Copy it over to RAM for execution
	memcpy(&Cla1mathTablesRunStart, &Cla1mathTablesLoadStart, (Uint32)&Cla1mathTablesLoadSize);		//Math tables the CLA needs for sin, cos and exp functions. Again needs to be copied from FLASH to RAM for execution
	Cla1Regs.MMEMCFG.bit.PROGE = 1;																	//let the CLA know everything has been copied over
	Cla1Regs.MVECT1 = (Uint16)((Uint32)&Cla1Task1 - (Uint32)&Cla1Prog_Start);						//Map out locations of CLA program space
	Cla1Regs.MVECT2 = (Uint16)((Uint32)&Cla1Task2 - (Uint32)&Cla1Prog_Start);
	Cla1Regs.MVECT3 = (Uint16)((Uint32)&Cla1Task3 - (Uint32)&Cla1Prog_Start);
	Cla1Regs.MPISRCSEL1.bit.PERINT1SEL = CLA_INT1_NONE;												//Currently Cla tasks are called by calling the function "CLA1ForceTask*"
	Cla1Regs.MPISRCSEL1.bit.PERINT2SEL = CLA_INT2_NONE;
	Cla1Regs.MPISRCSEL1.bit.PERINT3SEL = CLA_INT3_NONE;
	PieVectTable.CLA1_INT1 = &cla1_task1_isr;														//map out ISRs for the CLA
	PieVectTable.CLA1_INT2 = &cla1_task2_isr;
	PieVectTable.CLA1_INT3 = &cla1_task3_isr;
	Cla1Regs.MIER.all = 0x0007;
	PieCtrlRegs.PIEIER11.all = 0x0007;
	IER |= (M_INT11);
	Cla1Regs.MMEMCFG.all |= CLA_PROG_ENABLE | CLARAM0_ENABLE | CLARAM1_ENABLE | CLARAM2_ENABLE | CLA_RAM0CPUE;
	Cla1Regs.MCTL.bit.IACKE = 1;
	EDIS;

}

__interrupt void cla1_task1_isr(void){//After the CLA adjusts the appropiate biquad this will be called. Acknowledge

	PieCtrlRegs.PIEACK.all |= M_INT11;

}

__interrupt void cla1_task2_isr(void){ //After the CLA has finished adjusting filter this will be called
			//percentage variable percentge of sound masking vs music. .1 = 10% soundmasking
	masking = outputFromCLA*Sfactor;
//	GpioDataRegs.GPATOGGLE.bit.GPIO12 = 1; 		//used for timing interrupts. Will be taken out
	if(paging){									//if paging is on, the paging sampes will be be faded
		music = (ADCReading_paging - 2047);		//remove DC biasing for sampling
		if(!hasFaded)
			fade(fadeTo, fadeConstant);			//this is called initially to begin fading
	}

	else
			music = (ADCReading_music - 2047)*Mfactor;	//paging isnt on, mix the sound masking with the music. Again remove DC biasing

	if(!paging && hasFaded){							//time to bring volume back up.
		unFade(initialS, fadeConstant);
		music = (ADCReading_paging - 2047);
	}

	if ( music < 0 && masking < 0 ){								//mixing is done here. This algorithum maximizes dynamic range and prevents clipping
		mixed = (music + masking) - ((music * masking)/(-2000));
	}
	else if ( music > 0 && masking > 0 ){
		mixed = (music + masking) - ((music * masking)/(2000));
	}
	else{
		mixed = music + masking;
	}

	mixed = (mixed*Vfactor + 2000);									//add a DC bias to make signal all positive readings.
	pwm = (mixed)/scalar;											//Set pwm variable accordingly. It is course and not very precise.
																	//Precision is lost here but gained back in the line below fo the HRPWM
	hrpwm = ((( (float)mixed/scalar - pwm)*92));					//set variable Hrpwm accordingly. The loss of precision is added back here. The number was found while calibrating pwm.
	EPwm2Regs.CMPA.half.CMPA = pwm;         						//use pwm and hrpwm for the intended purposes
	EPwm2Regs.CMPA.half.CMPAHR = hrpwm << 8;
	PieCtrlRegs.PIEACK.all |= M_INT11;								//acknowledge

}

__interrupt void cla1_task3_isr(void){ //After the CLA initializes this will be called. Acknowledge

	PieCtrlRegs.PIEACK.all |= M_INT11;

}


void fade(float toValue, int rate){


	if(count == rate){
		count = 0;
		if(Sfactor > toValue)
			Sfactor = Sfactor - .01*initialS;
		else{
			hasFaded = 1;
			count = -1;
		}
	}
	else;
	count++;
}

void unFade(float toValue, int rate){

	if(count == rate){
		count = 0;
		if(Sfactor < toValue)
			Sfactor = Sfactor + .01*initialS;
		else{
			hasFaded = 0;
			count = -1;
			Sfactor = initialS;
		}
	}
	else;
	count++;
}
