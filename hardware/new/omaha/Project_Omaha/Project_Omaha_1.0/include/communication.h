#ifndef _COMMUNICATION_H_
#define _COMMUNICATION_H_

#include "constants.h"
#include "F2806x_Device.h"
#include "F2806x_Prototypes.h"

			void	initializeSCI();
__interrupt void	SciaRxFifoIsr();
__interrupt void	ScibRxFifoIsr();
			void	addressCheck();
			void	readCommand();


#endif
