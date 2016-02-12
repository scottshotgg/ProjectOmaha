################################################################################
# Automatically-generated file. Do not edit!
################################################################################

SHELL = cmd.exe

CG_TOOL_ROOT := C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.6

GEN_OPTS__FLAG := 
GEN_CMDS__FLAG := 

ORDERED_OBJS += \
"./source_TI/DSP28xxx_SectionCopy_nonBIOS.obj" \
"./source_TI/F2806x_Adc.obj" \
"./source_TI/F2806x_CodeStartBranch.obj" \
"./source_TI/F2806x_CpuTimers.obj" \
"./source_TI/F2806x_DBGIER.obj" \
"./source_TI/F2806x_DefaultIsr.obj" \
"./source_TI/F2806x_EPwm.obj" \
"./source_TI/F2806x_GlobalVariableDefs.obj" \
"./source_TI/F2806x_Gpio.obj" \
"./source_TI/F2806x_PieCtrl.obj" \
"./source_TI/F2806x_PieVect.obj" \
"./source_TI/F2806x_Sci.obj" \
"./source_TI/F2806x_SysCtrl.obj" \
"./source_TI/F2806x_usDelay.obj" \
"./source_TI/FIR_f32.obj" \
"./source/Cla_Funcs.obj" \
"./source/communication.obj" \
"./source/eq.obj" \
"./source/main.obj" \
"./source/timers.obj" \
"./source/whiteNoise.obj" \
"./cla_source_TI/CLASinCosTable.obj" \
"./cla_source_TI/CLAcos.obj" \
"./cla_source_TI/CLAexp10.obj" \
"./cla_source_TI/CLAexpTable.obj" \
"./cla_source_TI/CLAsin.obj" \
"./cla_source_TI/CLAsqrt.obj" \
"../F2806x_Headers_nonBIOS.cmd" \
"../Final.cmd" \
$(GEN_CMDS__FLAG) \
-l"libc.a" \

-include ../makefile.init

RM := DEL /F
RMDIR := RMDIR /S/Q

# All of the sources participating in the build are defined here
-include sources.mk
-include subdir_vars.mk
-include source_TI/subdir_vars.mk
-include source/subdir_vars.mk
-include cla_source_TI/subdir_vars.mk
-include subdir_rules.mk
-include source_TI/subdir_rules.mk
-include source/subdir_rules.mk
-include cla_source_TI/subdir_rules.mk
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
Project_Omaha_1.0.out \

EXE_OUTPUTS__QUOTED += \
"Project_Omaha_1.0.out" \

BIN_OUTPUTS += \
Project_Omaha_1.0.hex \

BIN_OUTPUTS__QUOTED += \
"Project_Omaha_1.0.hex" \


# All Target
all: Project_Omaha_1.0.out

# Tool invocations
Project_Omaha_1.0.out: $(OBJS) $(CMD_SRCS) $(GEN_CMDS)
	@echo 'Building target: $@'
	@echo 'Invoking: C2000 Linker'
	"C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.6/bin/cl2000" -v28 -ml -mt --float_support=fpu32 --cla_support=cla0 --vcu_support=vcu0 -O4 --opt_for_speed=5 --fp_mode=relaxed --advice:performance=all -g --diag_warning=225 --display_error_number --diag_wrap=off -z -m"Project_Omaha_1.0.map" --stack_size=0x300 --warn_sections -i"C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.6/lib" -i"C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.6/include" --reread_libs --display_error_number --diag_wrap=off --xml_link_info="Project_Omaha_1.0_linkInfo.xml" --rom_model -o "Project_Omaha_1.0.out" $(ORDERED_OBJS)
	@echo 'Finished building target: $@'
	@echo ' '

Project_Omaha_1.0.hex: $(EXE_OUTPUTS)
	@echo 'Invoking: C2000 Hex Utility'
	"C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.6/bin/hex2000"  -o "Project_Omaha_1.0.hex" $(EXE_OUTPUTS__QUOTED)
	@echo 'Finished building: $@'
	@echo ' '

# Other Targets
clean:
	-$(RM) $(EXE_OUTPUTS__QUOTED)$(BIN_OUTPUTS__QUOTED)
	-$(RM) "source\eq.pp" 
	-$(RM) "source_TI\F2806x_Adc.pp" "source_TI\F2806x_CpuTimers.pp" "source_TI\F2806x_DefaultIsr.pp" "source_TI\F2806x_EPwm.pp" "source_TI\F2806x_GlobalVariableDefs.pp" "source_TI\F2806x_Gpio.pp" "source_TI\F2806x_PieCtrl.pp" "source_TI\F2806x_PieVect.pp" "source_TI\F2806x_Sci.pp" "source_TI\F2806x_SysCtrl.pp" "source\Cla_Funcs.pp" "source\communication.pp" "source\main.pp" "source\timers.pp" "source\whiteNoise.pp" 
	-$(RM) "source_TI\DSP28xxx_SectionCopy_nonBIOS.obj" "source_TI\F2806x_Adc.obj" "source_TI\F2806x_CodeStartBranch.obj" "source_TI\F2806x_CpuTimers.obj" "source_TI\F2806x_DBGIER.obj" "source_TI\F2806x_DefaultIsr.obj" "source_TI\F2806x_EPwm.obj" "source_TI\F2806x_GlobalVariableDefs.obj" "source_TI\F2806x_Gpio.obj" "source_TI\F2806x_PieCtrl.obj" "source_TI\F2806x_PieVect.obj" "source_TI\F2806x_Sci.obj" "source_TI\F2806x_SysCtrl.obj" "source_TI\F2806x_usDelay.obj" "source_TI\FIR_f32.obj" "source\Cla_Funcs.obj" "source\communication.obj" "source\eq.obj" "source\main.obj" "source\timers.obj" "source\whiteNoise.obj" "cla_source_TI\CLASinCosTable.obj" "cla_source_TI\CLAcos.obj" "cla_source_TI\CLAexp10.obj" "cla_source_TI\CLAexpTable.obj" "cla_source_TI\CLAsin.obj" "cla_source_TI\CLAsqrt.obj" 
	-$(RM) "source_TI\DSP28xxx_SectionCopy_nonBIOS.pp" "source_TI\F2806x_CodeStartBranch.pp" "source_TI\F2806x_DBGIER.pp" "source_TI\F2806x_usDelay.pp" "source_TI\FIR_f32.pp" "cla_source_TI\CLASinCosTable.pp" "cla_source_TI\CLAcos.pp" "cla_source_TI\CLAexp10.pp" "cla_source_TI\CLAexpTable.pp" "cla_source_TI\CLAsin.pp" "cla_source_TI\CLAsqrt.pp" 
	-@echo 'Finished clean'
	-@echo ' '

.PHONY: all clean dependents
.SECONDARY:

-include ../makefile.targets

