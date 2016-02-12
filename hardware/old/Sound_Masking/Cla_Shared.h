#ifndef _CLA_SHARED_H_
#define _CLA_SHARED_H_
/*#############################################################################
 * This is the header that links the variables that need to be shared by both the CPU as well as the CLA
 *
 * SOUND_MASKING program written for F28377S MCU for Univerity of Texas at Dallas Applied Research Center and Speech Privacy Systems
 * Written by Matt Kramer and Scott Gayados
 *
 *
 */

#include "F2837xS_Cla_defines.h"
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif


#define PI				3.141592653589 //pi will be heavily used
#define INV2PI			0.159154943

extern float inputToCLA; 	//Sample input
extern float outputFromCLA; //Sample output
extern float octave;		//decade in which needs to be modified
extern float decibel;		//how much to modify the the selected decade by in decibels

__interrupt void Cla1Task1();	//execute filter
__interrupt void Cla1Task2();	//modify filter
__interrupt void Cla1Task3();	//set up enviornment


#ifdef __cplusplus
}
#endif // extern "C"

#endif //end of CLA_SHARED_H_ definition
