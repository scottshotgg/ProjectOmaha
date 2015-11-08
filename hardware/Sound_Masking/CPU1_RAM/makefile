################################################################################
# Automatically-generated file. Do not edit!
################################################################################

SHELL = cmd.exe

CG_TOOL_ROOT := C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.2

ORDERED_OBJS += \
"./eq.obj" \
"./main.obj" \
"./source/F2837xS_Adc.obj" \
"./source/F2837xS_CodeStartBranch.obj" \
"./source/F2837xS_CpuTimers.obj" \
"./source/F2837xS_DBGIER.obj" \
"./source/F2837xS_DefaultISR.obj" \
"./source/F2837xS_Dma.obj" \
"./source/F2837xS_ECap.obj" \
"./source/F2837xS_EPwm.obj" \
"./source/F2837xS_EQep.obj" \
"./source/F2837xS_Emif.obj" \
"./source/F2837xS_GlobalVariableDefs.obj" \
"./source/F2837xS_Gpio.obj" \
"./source/F2837xS_I2C.obj" \
"./source/F2837xS_Mcbsp.obj" \
"./source/F2837xS_PieCtrl.obj" \
"./source/F2837xS_PieVect.obj" \
"./source/F2837xS_Sci.obj" \
"./source/F2837xS_Spi.obj" \
"./source/F2837xS_SysCtrl.obj" \
"./source/F2837xS_TempSensorConv.obj" \
"./source/F2837xS_Upp.obj" \
"./source/F2837xS_sci_io.obj" \
"./source/F2837xS_sdfm_drivers.obj" \
"./source/F2837xS_struct.obj" \
"./source/F2837xS_usDelay.obj" \
"./source/FIR_f32.obj" \
"./include/examples_setup.obj" \
$(GEN_CMDS__FLAG) \
-l"rts2800_fpu32.lib" \
-l"F2837xS_RAM_CLA_lnk_cpu1.cmd" \
-l"F2837xS_Headers_nonBIOS.cmd" \
-l"libc.a" \

-include ../makefile.init

RM := DEL /F
RMDIR := RMDIR /S/Q

# All of the sources participating in the build are defined here
-include sources.mk
-include subdir_vars.mk
-include source/subdir_vars.mk
-include include/subdir_vars.mk
-include subdir_rules.mk
-include source/subdir_rules.mk
-include include/subdir_rules.mk
-include objects.mk

ifneq ($(MAKECMDGOALS),clean)
ifneq ($(strip $(S_DEPS)),)
-include $(S_DEPS)
endif
ifneq ($(strip $(S_UPPER_DEPS)),)
-include $(S_UPPER_DEPS)
endif
ifneq ($(strip $(S62_DEPS)),)
-include $(S62_DEPS)
endif
ifneq ($(strip $(C64_DEPS)),)
-include $(C64_DEPS)
endif
ifneq ($(strip $(ASM_DEPS)),)
-include $(ASM_DEPS)
endif
ifneq ($(strip $(CC_DEPS)),)
-include $(CC_DEPS)
endif
ifneq ($(strip $(S55_DEPS)),)
-include $(S55_DEPS)
endif
ifneq ($(strip $(C67_DEPS)),)
-include $(C67_DEPS)
endif
ifneq ($(strip $(CLA_DEPS)),)
-include $(CLA_DEPS)
endif
ifneq ($(strip $(C??_DEPS)),)
-include $(C??_DEPS)
endif
ifneq ($(strip $(CPP_DEPS)),)
-include $(CPP_DEPS)
endif
ifneq ($(strip $(S??_DEPS)),)
-include $(S??_DEPS)
endif
ifneq ($(strip $(C_DEPS)),)
-include $(C_DEPS)
endif
ifneq ($(strip $(C62_DEPS)),)
-include $(C62_DEPS)
endif
ifneq ($(strip $(CXX_DEPS)),)
-include $(CXX_DEPS)
endif
ifneq ($(strip $(C++_DEPS)),)
-include $(C++_DEPS)
endif
ifneq ($(strip $(ASM_UPPER_DEPS)),)
-include $(ASM_UPPER_DEPS)
endif
ifneq ($(strip $(K_DEPS)),)
-include $(K_DEPS)
endif
ifneq ($(strip $(C43_DEPS)),)
-include $(C43_DEPS)
endif
ifneq ($(strip $(INO_DEPS)),)
-include $(INO_DEPS)
endif
ifneq ($(strip $(S67_DEPS)),)
-include $(S67_DEPS)
endif
ifneq ($(strip $(SA_DEPS)),)
-include $(SA_DEPS)
endif
ifneq ($(strip $(S43_DEPS)),)
-include $(S43_DEPS)
endif
ifneq ($(strip $(OPT_DEPS)),)
-include $(OPT_DEPS)
endif
ifneq ($(strip $(PDE_DEPS)),)
-include $(PDE_DEPS)
endif
ifneq ($(strip $(S64_DEPS)),)
-include $(S64_DEPS)
endif
ifneq ($(strip $(C_UPPER_DEPS)),)
-include $(C_UPPER_DEPS)
endif
ifneq ($(strip $(C55_DEPS)),)
-include $(C55_DEPS)
endif
endif

-include ../makefile.defs

# Add inputs and outputs from these tool invocations to the build variables 
EXE_OUTPUTS += \
Sound_Masking.out \

EXE_OUTPUTS__QUOTED += \
"Sound_Masking.out" \


# All Target
all: Sound_Masking.out

# Tool invocations
Sound_Masking.out: $(OBJS) $(CMD_SRCS) $(GEN_CMDS)
	@echo 'Building target: $@'
	@echo 'Invoking: C2000 Linker'
	"C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.2/bin/cl2000" -v28 -ml -mt --cla_support=cla1 --float_support=fpu32 --tmu_support=tmu0 --vcu_support=vcu2 --opt_for_speed=4 --fp_reassoc=on --fp_mode=relaxed -g --define=CPU1 --diag_warning=225 --display_error_number --diag_suppress=10068 -z -m"Sound_Masking.map" --stack_size=0x100 --warn_sections -i"C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.2/lib" -i"C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.2/include" -i"C:/SD/ProjectOmaha/hardware/Sound_Masking/CPU1_RAM/" -i"C:/Users/Matt/F2837xS_headers/cmd" --reread_libs --define=CLA_C=1 --display_error_number --xml_link_info="Sound_Masking_linkInfo.xml" --rom_model -o "Sound_Masking.out" $(ORDERED_OBJS)
	@echo 'Finished building target: $@'
	@echo ' '

# Other Targets
clean:
	-$(RM) $(EXE_OUTPUTS__QUOTED)
	-$(RM) "eq.pp" 
	-$(RM) "main.pp" "source\F2837xS_Adc.pp" "source\F2837xS_CpuTimers.pp" "source\F2837xS_DefaultISR.pp" "source\F2837xS_Dma.pp" "source\F2837xS_ECap.pp" "source\F2837xS_EPwm.pp" "source\F2837xS_EQep.pp" "source\F2837xS_Emif.pp" "source\F2837xS_GlobalVariableDefs.pp" "source\F2837xS_Gpio.pp" "source\F2837xS_I2C.pp" "source\F2837xS_Mcbsp.pp" "source\F2837xS_PieCtrl.pp" "source\F2837xS_PieVect.pp" "source\F2837xS_Sci.pp" "source\F2837xS_Spi.pp" "source\F2837xS_SysCtrl.pp" "source\F2837xS_TempSensorConv.pp" "source\F2837xS_Upp.pp" "source\F2837xS_sci_io.pp" "source\F2837xS_sdfm_drivers.pp" "source\F2837xS_struct.pp" "include\examples_setup.pp" 
	-$(RM) "eq.obj" "main.obj" "source\F2837xS_Adc.obj" "source\F2837xS_CodeStartBranch.obj" "source\F2837xS_CpuTimers.obj" "source\F2837xS_DBGIER.obj" "source\F2837xS_DefaultISR.obj" "source\F2837xS_Dma.obj" "source\F2837xS_ECap.obj" "source\F2837xS_EPwm.obj" "source\F2837xS_EQep.obj" "source\F2837xS_Emif.obj" "source\F2837xS_GlobalVariableDefs.obj" "source\F2837xS_Gpio.obj" "source\F2837xS_I2C.obj" "source\F2837xS_Mcbsp.obj" "source\F2837xS_PieCtrl.obj" "source\F2837xS_PieVect.obj" "source\F2837xS_Sci.obj" "source\F2837xS_Spi.obj" "source\F2837xS_SysCtrl.obj" "source\F2837xS_TempSensorConv.obj" "source\F2837xS_Upp.obj" "source\F2837xS_sci_io.obj" "source\F2837xS_sdfm_drivers.obj" "source\F2837xS_struct.obj" "source\F2837xS_usDelay.obj" "source\FIR_f32.obj" "include\examples_setup.obj" 
	-$(RM) "source\F2837xS_CodeStartBranch.pp" "source\F2837xS_DBGIER.pp" "source\F2837xS_usDelay.pp" "source\FIR_f32.pp" 
	-@echo 'Finished clean'
	-@echo ' '

.PHONY: all clean dependents
.SECONDARY:

-include ../makefile.targets

