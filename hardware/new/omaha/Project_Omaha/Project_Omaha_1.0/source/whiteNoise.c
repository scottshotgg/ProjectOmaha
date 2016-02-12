
/*

whitenoise.c

Version 1.0 First stable release.

This is confidential code and is written by Matthew Kramer of the University of Texas at Dallas Applied Research Center for Speech Privacy Systems

This source code is used for creating what seems like randoms numbers. The unpredictable nature of the fibbocacci sequence yields pseudo random numbers that seem
more random than tradional methods. This source code uses 16 of them. shifting them each in sequence and then loading them it up it decreases correlation vs. used one single one (proven in testing)




*/
#include "whiteNoise.h"


int i;

uint32_t     			lfsr_1 									= 0xF8621234;
uint32_t     			lfsr_2 									= 0x12784321;
uint32_t				lfsr_3 									= 0x57621465;
uint32_t				lfsr_4 									= 0x21345676;
uint32_t				lfsr_5 									= 0x85672468;
uint32_t				lfsr_6 									= 0x01011011;
uint32_t				lfsr_7 									= 0x11065431;
uint32_t				lfsr_8 									= 0xA1AF3214;
uint32_t     			lfsr_9 									= 0x52F11F68; //initialize linear feedback shift registers
uint32_t     			lfsr_10 								= 0xF158F620;
uint32_t				lfsr_11 								= 0xF16C8F13;
uint32_t				lfsr_12 								= 0x65F21561;
uint32_t				lfsr_13									= 0xF86F1235;
uint32_t				lfsr_14 								= 0x7788221F;
uint32_t				lfsr_15 								= 0xEF123FEE;
uint32_t				lfsr_16 								= 0xF0668742;
uint32_t				carrybit								= 0;
float 					lfsr_Array								[ARRAYSIZE];


/*

		32 bit Fibonacci Linear Feedback Shift Register

	 	 bit 0		bit 1	   bit 3				 bit 25		bit 26	   bit 27     bit 28     bit 29     bit 30     bit 31    bit 32
 	 	 _____		_____	   _____				 _____      _____      _____      _____      _____      _____      _____      _____
		|	  |	   |	 |	  |	    |				|	  |	   |	 |	  |	    |    |	   |    |	  |    |	 |    |	    |    |	   |
     ___|  1  |--->|  0  |--->|  0  |---		--->|  0  |--->|  1  |--->|  0  |--->|  1  |--->|  1  |--->|  0  |--->|  1  |--->|  0  |
	/\	|_____|	   |_____|	  |_____|	........	|_____|    |_____|    |_____|    |_____|    |_____|    |_____|    |_____|    |_____|
	||												  |		       |										  |		                |
	||												  |   ______   |                                          |        ______       |
	||												  |  |		|  |                                          |       |		 |      |
	||												  |__| xor	|__|                                          |_______| xor	 |______|
	||													 |______|                                                     |______|
	||					 _______						    ||                                                           ||
	||                  |       |___________________________\/                                                           ||
    ||__________________|  xor  |                                                                                        ||
                        |       |________________________________________________________________________________________\/
						|_______|
*/


void loadWhiteNoiseSamples(int startingSample, int length){ //load lfsr_ array up with values, 16 at a time. Using 16 different registers decreasing correlation and increases pleasentness

	for(i = startingSample; i <= (startingSample + length);){

		carrybit = ((lfsr_1 ) ^ (lfsr_1 >> 2) ^ (lfsr_1 >> 3) ^ (lfsr_1 >> 5) & 1); //only the last bit of carrybit is used. This is key in understanding what is going on here
		lfsr_1 = (lfsr_1 >> 1)|(carrybit << 31);									//shift the whole register over, then or the first bit of LFSR (which will be zero) with the last bit of carry bit (shift by 31 bits)
		lfsr_Array[i++] = (float)(int16_t)((0x0000FFFF)&(lfsr_1));					//only using 16 bit of the 32 bit number. Also cast to float
		carrybit = ((lfsr_2 ) ^ (lfsr_2 >> 2) ^ (lfsr_2 >> 3) ^ (lfsr_2 >> 5) & 1);
		lfsr_2 = (lfsr_2 >> 1)|(carrybit << 31);
		lfsr_Array[i++] = (float)(int16_t)((0x0000FFFF)&(lfsr_2));
		carrybit = ((lfsr_3 ) ^ (lfsr_3 >> 2) ^ (lfsr_3 >> 3) ^ (lfsr_3 >> 5) & 1);
		lfsr_3 = (lfsr_3 >> 1)|(carrybit << 31);
		lfsr_Array[i++] = (float)(int16_t)((0x0000FFFF)&(lfsr_3));
		carrybit = ((lfsr_4 ) ^ (lfsr_4 >> 2) ^ (lfsr_4 >> 3) ^ (lfsr_4 >> 5) & 1);
		lfsr_4 = (lfsr_4 >> 1)|(carrybit << 31);
		lfsr_Array[i++] = (float)(int16_t)((0x0000FFFF)&(lfsr_4));
		carrybit = ((lfsr_5 ) ^ (lfsr_5 >> 2) ^ (lfsr_5 >> 3) ^ (lfsr_5 >> 5) & 1);
		lfsr_5 = (lfsr_5 >> 1)|(carrybit << 31);
		lfsr_Array[i++] = (float)(int16_t)((0x0000FFFF)&(lfsr_5));
		carrybit = ((lfsr_6 ) ^ (lfsr_6 >> 2) ^ (lfsr_6 >> 3) ^ (lfsr_6 >> 5) & 1);
		lfsr_6 = (lfsr_6 >> 1)|(carrybit << 31);
		lfsr_Array[i++] = (float)(int16_t)((0x0000FFFF)&(lfsr_6));
		carrybit = ((lfsr_7 ) ^ (lfsr_7 >> 2) ^ (lfsr_7 >> 3) ^ (lfsr_7 >> 5) & 1);
		lfsr_7 = (lfsr_7 >> 1)|(carrybit << 31);
		lfsr_Array[i++] = (float)(int16_t)((0x0000FFFF)&(lfsr_7));
		carrybit = ((lfsr_8 ) ^ (lfsr_8 >> 2) ^ (lfsr_8 >> 3) ^ (lfsr_8 >> 5) & 1);
		lfsr_8 = (lfsr_8 >> 1)|(carrybit << 31);
 		lfsr_Array[i++] = (float)(int16_t)((0x0000FFFF)&(lfsr_8));
		carrybit = ((lfsr_9 ) ^ (lfsr_9 >> 2) ^ (lfsr_9 >> 3) ^ (lfsr_9 >> 5) & 1);
		lfsr_9 = (lfsr_9 >> 1)|(carrybit << 31);
		lfsr_Array[i++] = (float)(int16_t)((0x0000FFFF)&(lfsr_9));
		carrybit = ((lfsr_10 ) ^ (lfsr_10 >> 2) ^ (lfsr_10 >> 3) ^ (lfsr_10 >> 5) & 1);
		lfsr_10 = (lfsr_10 >> 1)|(carrybit << 31);
		lfsr_Array[i++] = (float)(int16_t)((0x0000FFFF)&(lfsr_10));
		carrybit = ((lfsr_11 ) ^ (lfsr_11 >> 2) ^ (lfsr_11 >> 3) ^ (lfsr_11 >> 5) & 1);
		lfsr_11 = (lfsr_11 >> 1)|(carrybit << 31);
		lfsr_Array[i++] = (float)(int16_t)((0x0000FFFF)&(lfsr_11));
		carrybit = ((lfsr_12 ) ^ (lfsr_12 >> 2) ^ (lfsr_12 >> 3) ^ (lfsr_12 >> 5) & 1);
		lfsr_12 = (lfsr_12 >> 1)|(carrybit << 31);
		lfsr_Array[i++] = (float)(int16_t)((0x0000FFFF)&(lfsr_12));
		carrybit = ((lfsr_13 ) ^ (lfsr_13 >> 2) ^ (lfsr_13 >> 3) ^ (lfsr_13 >> 5) & 1);
		lfsr_13 = (lfsr_13 >> 1)|(carrybit << 31);
		lfsr_Array[i++] = (float)(int16_t)((0x0000FFFF)&(lfsr_13));
		carrybit = ((lfsr_14 ) ^ (lfsr_14 >> 2) ^ (lfsr_14 >> 3) ^ (lfsr_14 >> 5) & 1);
		lfsr_14 = (lfsr_14 >> 1)|(carrybit << 31);
		lfsr_Array[i++] = (float)(int16_t)((0x0000FFFF)&(lfsr_14));
		carrybit = ((lfsr_15 ) ^ (lfsr_15 >> 2) ^ (lfsr_15 >> 3) ^ (lfsr_15 >> 5) & 1);
		lfsr_15 = (lfsr_15 >> 1)|(carrybit << 31);
		lfsr_Array[i++] = (float)(int16_t)((0x0000FFFF)&(lfsr_15));
		carrybit = ((lfsr_16 ) ^ (lfsr_16 >> 2) ^ (lfsr_16 >> 3) ^ (lfsr_16 >> 5) & 1);
		lfsr_16 = (lfsr_16 >> 1)|(carrybit << 31);
		lfsr_Array[i++] = (float)(int16_t)((0x0000FFFF)&(lfsr_16));
	}
}
