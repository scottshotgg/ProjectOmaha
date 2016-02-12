#ifndef _CONSTANTS_H_
#define _CONSTANTS_H_

#include "fpu_filter.h"


#define ADCSAMPLERATE 	1000
#define SAMPLERATE	    12000
#define AUDSAMPLERATE 	24000
#define ARRAYSIZE		512
#define BLOCKSIZE		128
#define	ADCARRAYSIZE	128
#define	PERIOD			40
#define BUFFERSIZE		3
#define FIRORDER		299
#define	NUM_BAND	  	21

#define CLA_PROG_ENABLE      0x0001
#define CLARAM0_ENABLE       0x0010
#define CLARAM1_ENABLE       0x0020
#define CLARAM2_ENABLE       0x0040
#define CLA_RAM0CPUE         0x0100
#define CLA_RAM1CPUE         0x0200
#define CLA_RAM2CPUE         0x0400


#endif
