//###########################################################################
//
// FILE:   F2806x_Examples.h
//
// TITLE:  F2806x Device Definitions.
//
//###########################################################################
// $TI Release: F2806x C/C++ Header Files and Peripheral Examples V150 $
// $Release Date: June 16, 2015 $
// $Copyright: Copyright (C) 2011-2015 Texas Instruments Incorporated -
//             http://www.ti.com/ ALL RIGHTS RESERVED $
//###########################################################################

#ifndef F2806x_PROTOTYPES_H
#define F2806x_PROTOTYPES_H

#ifdef __cplusplus
extern "C" {
#endif

/*-----------------------------------------------------------------------------
      Specify the PLL control register (PLLCR) and divide select (DIVSEL) value.
-----------------------------------------------------------------------------*/
#define DSP28_DIVSEL   2 // Enable /2 for SYSCLKOUT


#define DSP28_PLLCR   18  // Uncomment for 90 MHz devices [90 MHz = (10MHz * 18)/2]

/*-----------------------------------------------------------------------------*/
#define CPU_RATE   11.111L   // for a 90MHz CPU clock speed (SYSCLKOUT)

#define PLL2_PLLSRC			0x2		// PLL2 Input X1
#define PLL2_SYSCLK2DIV2DIS   	0 	// PLL2 Output /2
#define PLL2_PLLMULT    	6	// (CLKSOURCE*6) /2 = SYSCLKOUT2

// The following pointer to a function call calibrates the ADC and internal oscillators
#define Device_cal (void   (*)(void))0x3D7C80

//---------------------------------------------------------------------------
// Include Example Header Files:
//

#include "F2806x_GlobalPrototypes.h"         // Prototypes for global functions within the .c files.
#include "F2806x_EPwm_defines.h"             // Macros used for PWM examples.              
#include "F2806x_Cla_defines.h"              // Macros used for CLA examples.

// Include files not used with DSP/BIOS
#ifndef DSP28_BIOS
#include "F2806x_DefaultISR.h"
#endif

// DO NOT MODIFY THIS LINE.
#define DELAY_US(A)  DSP28x_usDelay(((((long double) A * 1000.0L) / (long double)CPU_RATE) - 9.0L) / 5.0L)

#ifdef __cplusplus
}
#endif /* extern "C" */

#endif  // end of F2806x_EXAMPLES_H definition

//===========================================================================
// End of file.
//===========================================================================
