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

#include "F2806x_Cla_defines.h"
#include <stdint.h>
#include "constants.h"
#ifdef __cplusplus
extern "C" {
#endif


#define PI				3.141592653589 //pi will be heavily used
#define INV2PI			0.159154943

extern float inputToCLA; 	//Sample input
extern float outputFromCLA; //Sample output
extern int	band;			//decade in which needs to be modified
extern float decibel;		//how much to modify the the selected decade by in decibels
extern float bandfreq [NUM_BAND];


__interrupt void 	Cla1Task1();	//modify filter
__interrupt void 	Cla1Task2();	//execute filter
__interrupt void 	Cla1Task3();	//set up enviornment
			void 	Cla_Init(void);
__interrupt void 	cla1_task1_isr(void);
__interrupt void 	cla1_task2_isr(void);
__interrupt void 	cla1_task3_isr(void);

#ifdef __cplusplus
}
#endif // extern "C"

#endif //end of CLA_SHARED_H_ definition
