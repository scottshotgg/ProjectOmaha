/*
 *
 *  Author: Matthew Kramer
 *  Revision 1.0 (8 February 2016)
 *
 *  THIS CODE IS CONFIDENTIAL AND FOR THE EYES OF A SELECT FEW
 *
 * 	main.c version 1.0
 *
 *	main source code that also includes standalone intialization functions
 *
 *
 *
 *
*/

#include <math.h>
#include "whiteNoise.h"		  //whiteNoise generation
#include "timers.h"			  //configure timers specific to this application
#include "communication.h"
#include "Cla_Shared.h"


//lfsrarray, needs to be in scope of whitenoise.c
extern 	float 			lfsr_Array		[ARRAYSIZE];

//The programs behaves much better when these variables are includes in the scope of main and not the scopes of the source that uses them.
//Espeically when the variables are used in several scopes
//variables used after cla is done filtering
		int				count			= 0;			//count is used for the fad function
		int				paging    		= 0;			//used as a boolean to determin if a page has been sent
		int				hasFaded		= 0;			//keeps track of the fade status
		float			initialS        = 1;
		int				ADCReading_music;
		int				ADCReading_paging;
		int 			unitID			= 0;



		uint16_t 		currentAddress	= 0;
		float			filteredArray	[ARRAYSIZE];
		float			Sfactor         = 1;
		float			Mfactor 		= 0;
		float			Vfactor			= .1;
		float 			seconds			= 1;
		float			fadeTo 			= .5;
		int				fadeConstant	= 0;
		FIR_FP			firFP			= FIR_FP_DEFAULTS;
		FIR_FP_Handle 	hnd_firFP		= &firFP;
		float			FIRBuffer		[FIRORDER + 1];
		int				foundNorm1 		= 0;
		char 			receivedChar	[BUFFERSIZE];
		//300 tap FIR filter used. These coeffceints were generated in matlab
		float			coeffcients1	[FIRORDER + 1]       		   ={ 0.031416,0.031903,0.032397,0.032899,0.033409,0.033926,0.034452,0.034986,0.035528,0.036079,0.036638,0.037205,
																0.037782,0.038367,0.038962,0.039566,0.040179,0.040801,0.041433,0.042075,0.042727,0.043389,0.044062,0.044744,
																0.045438,0.046142,0.046857,0.047583,0.048320,0.049069,0.049829,0.050601,0.051385,0.052182,0.052990,0.053811,
																0.054645,0.055492,0.056352,0.057225,0.058111,0.059012,0.059926,0.060855,0.061798,0.062755,0.063728,0.064715,
																0.065718,0.066736,0.067770,0.068820,0.069887,0.070970,0.072069,0.073186,0.074320,0.075472,0.076641,0.077829,
																0.079035,0.080259,0.081503,0.082766,0.084048,0.085351,0.086673,0.088016,0.089380,0.090765,0.092171,0.093599,
																0.095050,0.096523,0.098018,0.099537,0.101079,0.102646,0.104236,0.105851,0.107491,0.109157,0.110848,0.112566,
																0.114310,0.116081,0.117880,0.119707,0.121561,0.123445,0.125358,0.127300,0.129273,0.131276,0.133310,0.135376,
																0.137473,0.139603,0.141767,0.143963,0.146194,0.148459,0.150760,0.153096,0.155468,0.157877,0.160323,0.162807,
																0.165330,0.167892,0.170493,0.173135,0.175818,0.178542,0.181309,0.184118,0.186971,0.189868,0.192810,0.195798,
																0.198832,0.201913,0.205041,0.208218,0.211445,0.214721,0.218048,0.221427,0.224858,0.228342,0.231880,0.235473,
																0.239122,0.242827,0.246590,0.250410,0.254291,0.258231,0.262232,0.266295,0.270422,0.274612,0.278867,0.283188,
																0.287576,0.292032,0.296557,0.301152,0.305818,0.310557,0.315369,0.320256,0.325218,0.330258,0.335375,0.340572,
																0.345849,0.351208,0.356650,0.362176,0.367788,0.373487,0.379274,0.385151,0.391119,0.397179,0.403333,0.409583,
																0.415929,0.422374,0.428919,0.435565,0.442314,0.449168,0.456128,0.463195,0.470373,0.477661,0.485062,0.492578,
																0.500211,0.507962,0.515833,0.523825,0.531942,0.540185,0.548555,0.557055,0.565686,0.574452,0.583353,0.592392,
																0.601571,0.610892,0.620358,0.629970,0.639732,0.649644,0.659711,0.669933,0.680314,0.690855,0.701560,0.712431,
																0.723470,0.734680,0.746064,0.757624,0.769363,0.781285,0.793391,0.805684,0.818168,0.830846,0.843720,0.856793,
																0.870069,0.883551,0.897242,0.911145,0.925263,0.939600,0.954159,0.968943,0.983957,0.999204,1.014686,1.030409,
																1.046375,1.062589,1.079054,1.095774,1.112753,1.129995,1.147504,1.165285,1.183341,1.201677,1.220297,1.239205,
																1.258407,1.277906,1.297707,1.317815,1.338234,1.358970,1.380028,1.401411,1.423126,1.445177,1.467570,1.490310,
																1.513403,1.536853,1.560667,1.584849,1.609406,1.634344,1.659668,1.685385,1.711500,1.738020,1.764950,1.792298,
																1.820070,1.848272,1.876911,1.905994,1.935527,1.965518,1.995974,2.026902,2.058309,2.090202,2.122590,2.155479,
																2.188879,2.222795,2.257237,2.292213,2.327731,2.363800,2.400427,2.437621,2.475392,2.513749,2.552699,2.592253,
																2.632420,2.673210,2.714631,2.756694,2.799409,2.842786,2.886835,2.931567,2.976991,3.023120,3.069963,3.117532 };






void 	HRPWM2_Config();
void 	initialize();
void 	initializeFilter();
void	initializeADC();
void 	configGPIO();

extern Uint16 RamfuncsLoadStart;
extern Uint16 RamfuncsLoadSize;
extern Uint16 RamfuncsRunStart;

#pragma DATA_SECTION(inputToCLA,"CpuToCla1MsgRAM")
float inputToCLA;
#pragma DATA_SECTION(outputFromCLA,"Cla1ToCpuMsgRAM")
float outputFromCLA;
#pragma DATA_SECTION(band,"CpuToCla1MsgRAM")
int band;
#pragma DATA_SECTION(decibel,"CpuToCla1MsgRAM")
float decibel;
#pragma DATA_SECTION(bandfreq,"CpuToCla1MsgRAM")
float bandfreq [NUM_BAND] = {80, 100, 125, 160, 200, 250, 315, 400, 500, 625, 800, 1000, 1250, 1600, 2000, 2500, 3150, 4000, 5000, 6300, 8000};

void main(void)
{

	memcpy(&RamfuncsRunStart, &RamfuncsLoadStart, (Uint32)&RamfuncsLoadSize);
	initialize();
	for(;;){

		loadWhiteNoiseSamples(currentAddress, BLOCKSIZE);
		currentAddress = currentAddress + BLOCKSIZE;
		if(currentAddress >= ARRAYSIZE){
			currentAddress = 0;
		}
	}
}


void initialize(){

	inputToCLA = 0;			//these variables need to be initalized by the CPU after compiletime. This is the nature of the CLA
	outputFromCLA = 0;
	decibel = 0;
	EALLOW;
	InitSysCtrl();			//set system up. Provided by TI. Located in ./source_TI/F2806x_SysCtrl.c
	InitAdc();				//initialize ADC. Provided by TI. Located in ./source_TI/F2806x_Adc.c
	EDIS;
	InitEPwmGpio();			//initialize Epwm. Provided by TI. Located in ./source_TI/F2806x_Epwm.c
	initializeFilter();		//set filter up. Located below
	DINT;
	InitPieCtrl();			//set up PIE block. PIE = Peripheral Interrupt Expansion
	EALLOW;
	fadeConstant = AUDSAMPLERATE/((initialS - fadeTo)*100); //one second by default
	IER = 0x0000;			//Make sure all interrupts are disabled by default
	IFR = 0x0000;
	InitPieVectTable();		//set up the PIE vector table (where interrupts point to)  located in ./source_TI/F2806x_PieVect.c
	EALLOW;
	PieCtrlRegs.PIECTRL.bit.ENPIE  = 1;   // Enable the PIE block
	IER |= 0x400;   					  // Enable CPU INT
	EALLOW;
	SysCtrlRegs.PCLKCR0.bit.TBCLKSYNC = 1;
	SysCtrlRegs.LOSPCP.bit.LSPCLK = 6; //increase low speed periphial clock divider to 12 to reduce power comsumption.
	EDIS;
	Cla_Init();
	HRPWM2_Config();					//set up HRPWM
	Cla1ForceTask3();					//initialize IIR biquad structure for Equalizer
	initializeTimers();					//setup timers. Located in ./timers.c
	initializeSCI();					//setup communication interface. Located in ./communication
	initializeADC();					//setup ADC
	EINT;   // Enable Global __interrupt INTM
	ERTM;   // Enable Global realtime __interrupt DBGM
	configGPIO();	//set all unused pins as outputs as they are left floating.
}

void HRPWM2_Config()
{

	EPwm2Regs.TBCTL.bit.PRDLD = TB_IMMEDIATE;	        // set Immediate load
	EPwm2Regs.TBPRD = PERIOD - 1;		                // PWM frequency = 1 / period
	EPwm2Regs.CMPA.half.CMPA = PERIOD / 2;              // set duty 50% initially
	EPwm2Regs.CMPA.half.CMPAHR = (1 << 8);              // initialize HRPWM extension	                    // set duty 50% initially
	EPwm2Regs.TBPHS.all = 0;
	EPwm2Regs.TBCTR = 0;
	EPwm2Regs.TBCTL.bit.CTRMODE = TB_COUNT_UP;
	EPwm2Regs.TBCTL.bit.PHSEN = TB_DISABLE;		         // ePWM2 is the Master
	EPwm2Regs.TBCTL.bit.SYNCOSEL = TB_SYNC_DISABLE;
	EPwm2Regs.TBCTL.bit.HSPCLKDIV = TB_DIV1;
	EPwm2Regs.TBCTL.bit.CLKDIV = TB_DIV1;
	EPwm2Regs.CMPCTL.bit.LOADAMODE = CC_CTR_ZERO;
	EPwm2Regs.CMPCTL.bit.LOADBMODE = CC_CTR_ZERO;
	EPwm2Regs.CMPCTL.bit.SHDWAMODE = CC_SHADOW;
	EPwm2Regs.CMPCTL.bit.SHDWBMODE = CC_SHADOW;
	EPwm2Regs.AQCTLA.bit.ZRO = AQ_CLEAR;                  // PWM toggle low/high
	EPwm2Regs.AQCTLA.bit.CAU = AQ_SET;
	EPwm2Regs.AQCTLB.bit.ZRO = AQ_CLEAR;
	EPwm2Regs.AQCTLB.bit.CBU = AQ_SET;
	EALLOW;
	EPwm2Regs.HRCNFG.all = 0x0;
	EPwm2Regs.HRCNFG.bit.EDGMODE = HR_REP;                //MEP control on Rising edge
	EPwm2Regs.HRCNFG.bit.CTLMODE = HR_CMP;
	EPwm2Regs.HRCNFG.bit.HRLOAD  = HR_CTR_ZERO;

	EDIS;
}


void initializeFilter(){			//initialize filter structure and normalize coeffceints so that the output doesnt overflow.

	hnd_firFP->order       = FIRORDER;		//setup FIR filter structure.
	hnd_firFP->dbuffer_ptr = FIRBuffer;
	hnd_firFP->init(hnd_firFP);				//let it initialize itself

	int i;
	float sum = 0;
	float mag;
	hnd_firFP->coeff_ptr   = (float *)coeffcients1;
	if(!foundNorm1){							//this function normalizes the coeffcients that are stored in memory. This is so the signal does not clip
		foundNorm1 = 1;
		for(i = 0; i < FIRORDER; i++){//add up squares of filter for magnitude calculation

			sum += coeffcients1[i]*coeffcients1[i]; //take squareroot
		}
		mag = sqrt(sum);
		for(i = 0; i < FIRORDER + 1; i++){

			coeffcients1[i] = coeffcients1[i]/(mag*64); //normalize coeffcients before filtering
		}

	}
}

void initializeADC(){

	EALLOW;
	PieCtrlRegs.PIEIER1.bit.INTx1  = 1;
	IER |= M_INT1;
	PieVectTable.ADCINT1 = &adc_isr; 		//located in timer.c since timer will fire ADC
	AdcRegs.ADCCTL2.bit.CLKDIV2EN = 1;
	AdcRegs.ADCCTL2.bit.CLKDIV2EN = 1;
	AdcRegs.ADCCTL1.bit.INTPULSEPOS = 1;
	AdcRegs.ADCCTL1.bit.ADCPWDN = 1;
	AdcRegs.ADCSOC0CTL.bit.CHSEL = 2;
	AdcRegs.ADCSOC0CTL.bit.ACQPS = 14;
	AdcRegs.INTSEL1N2.bit.INT1SEL = 0;
	AdcRegs.INTSEL1N2.bit.INT1E = 1;
	AdcRegs.ADCINTFLGCLR.bit.ADCINT1 = 1;
	AdcRegs.ADCSOC0CTL.bit.TRIGSEL 	= 0x01;
	AdcRegs.ADCSOC1CTL.bit.CHSEL = 3;
	AdcRegs.ADCSOC1CTL.bit.ACQPS = 14;
	AdcRegs.ADCSOC1CTL.bit.TRIGSEL 	= 0x01;
}

void configGPIO(){



}
