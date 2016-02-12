#ifndef _TIMERS_H_
#define _TIMERS_H_

#include 	"constants.h"
#include "F2806x_Device.h"
#include "F2806x_Prototypes.h"

			void	initializeTimers();
__interrupt void 	cpu_timer0_isr(void);
__interrupt	void 	cpu_timer1_isr(void);
__interrupt	void 	cpu_timer2_isr(void);
__interrupt void 	adc_isr(void);

#endif
