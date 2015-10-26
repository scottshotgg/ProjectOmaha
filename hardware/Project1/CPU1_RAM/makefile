################################################################################
# Automatically-generated file. Do not edit!
################################################################################

SHELL = cmd.exe

CG_TOOL_ROOT := C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.2

ORDERED_OBJS += \
"./main.obj" \
"./Source/F2837xS_Adc.obj" \
"./Source/F2837xS_CodeStartBranch.obj" \
"./Source/F2837xS_CpuTimers.obj" \
"./Source/F2837xS_DBGIER.obj" \
"./Source/F2837xS_DefaultISR.obj" \
"./Source/F2837xS_Dma.obj" \
"./Source/F2837xS_EPwm.obj" \
"./Source/F2837xS_EQep.obj" \
"./Source/F2837xS_Emif.obj" \
"./Source/F2837xS_GlobalVariableDefs.obj" \
"./Source/F2837xS_Gpio.obj" \
"./Source/F2837xS_I2C.obj" \
"./Source/F2837xS_Mcbsp.obj" \
"./Source/F2837xS_PieCtrl.obj" \
"./Source/F2837xS_PieVect.obj" \
"./Source/F2837xS_Sci.obj" \
"./Source/F2837xS_Spi.obj" \
"./Source/F2837xS_SysCtrl.obj" \
"./Source/F2837xS_Upp.obj" \
"./Source/F2837xS_sci_io.obj" \
"./Source/F2837xS_sdfm_drivers.obj" \
"./Source/F2837xS_struct.obj" \
"./Source/F2837xS_usDelay.obj" \
"./Source/FIR_f32.obj" \
"./Source/examples_setup.obj" \
$(GEN_CMDS__FLAG) \
-l"rts2800_fpu32.lib" \
-l"2837xS_Generic_RAM_lnk.cmd" \
-l"F2837xS_Headers_nonBIOS.cmd" \
-l"libc.a" \

-include ../makefile.init

RM := DEL /F
RMDIR := RMDIR /S/Q

# All of the sources participating in the build are defined here
-include sources.mk
-include subdir_vars.mk
-include Source/subdir_vars.mk
-include subdir_rules.mk
-include Source/subdir_rules.mk
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
Project1.out \

EXE_OUTPUTS__QUOTED += \
"Project1.out" \

BIN_OUTPUTS += \
Project1.hex \

BIN_OUTPUTS__QUOTED += \
"Project1.hex" \


# All Target
all: Project1.out

# Tool invocations
Project1.out: $(OBJS) $(CMD_SRCS) $(GEN_CMDS)
	@echo 'Building target: $@'
	@echo 'Invoking: C2000 Linker'
	"C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.2/bin/cl2000" -v28 -ml -mt --float_support=fpu32 --tmu_support=tmu0 --cla_support=cla1 --vcu_support=vcu2 --fp_mode=relaxed -g --define=CPU1 --diag_warning=225 --display_error_number --diag_suppress=10063 -z -m"Project1.map" --stack_size=0x100 --warn_sections -i"C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.2/lib" -i"C:/ti/controlSUITE/device_support/F2837xS/v160/F2837xS_common/cmd" -i"C:/ti/controlSUITE/device_support/F2837xS/v160/F2837xS_headers/cmd" -i"C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.2/include" --reread_libs --display_error_number --xml_link_info="Project1_linkInfo.xml" --entry_point=code_start --rom_model -o "Project1.out" $(ORDERED_OBJS)
	@echo 'Finished building target: $@'
	@echo ' '

Project1.hex: $(EXE_OUTPUTS)
	@echo 'Invoking: C2000 Hex Utility'
	"C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.2/bin/hex2000"  -o "Project1.hex" $(EXE_OUTPUTS__QUOTED)
	@echo 'Finished building: $@'
	@echo ' '

# Other Targets
clean:
	-$(RM) $(EXE_OUTPUTS__QUOTED)$(BIN_OUTPUTS__QUOTED)
	-$(RM) "main.pp" "Source\F2837xS_Adc.pp" "Source\F2837xS_CpuTimers.pp" "Source\F2837xS_DefaultISR.pp" "Source\F2837xS_Dma.pp" "Source\F2837xS_EPwm.pp" "Source\F2837xS_EQep.pp" "Source\F2837xS_Emif.pp" "Source\F2837xS_GlobalVariableDefs.pp" "Source\F2837xS_Gpio.pp" "Source\F2837xS_I2C.pp" "Source\F2837xS_Mcbsp.pp" "Source\F2837xS_PieCtrl.pp" "Source\F2837xS_PieVect.pp" "Source\F2837xS_Sci.pp" "Source\F2837xS_Spi.pp" "Source\F2837xS_SysCtrl.pp" "Source\F2837xS_Upp.pp" "Source\F2837xS_sci_io.pp" "Source\F2837xS_sdfm_drivers.pp" "Source\F2837xS_struct.pp" "Source\examples_setup.pp" 
	-$(RM) "main.obj" "Source\F2837xS_Adc.obj" "Source\F2837xS_CodeStartBranch.obj" "Source\F2837xS_CpuTimers.obj" "Source\F2837xS_DBGIER.obj" "Source\F2837xS_DefaultISR.obj" "Source\F2837xS_Dma.obj" "Source\F2837xS_EPwm.obj" "Source\F2837xS_EQep.obj" "Source\F2837xS_Emif.obj" "Source\F2837xS_GlobalVariableDefs.obj" "Source\F2837xS_Gpio.obj" "Source\F2837xS_I2C.obj" "Source\F2837xS_Mcbsp.obj" "Source\F2837xS_PieCtrl.obj" "Source\F2837xS_PieVect.obj" "Source\F2837xS_Sci.obj" "Source\F2837xS_Spi.obj" "Source\F2837xS_SysCtrl.obj" "Source\F2837xS_Upp.obj" "Source\F2837xS_sci_io.obj" "Source\F2837xS_sdfm_drivers.obj" "Source\F2837xS_struct.obj" "Source\F2837xS_usDelay.obj" "Source\FIR_f32.obj" "Source\examples_setup.obj" 
	-$(RM) "Source\F2837xS_CodeStartBranch.pp" "Source\F2837xS_DBGIER.pp" "Source\F2837xS_usDelay.pp" "Source\FIR_f32.pp" 
	-@echo 'Finished clean'
	-@echo ' '

.PHONY: all clean dependents
.SECONDARY:

-include ../makefile.targets

