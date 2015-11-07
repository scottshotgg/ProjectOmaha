################################################################################
# Automatically-generated file. Do not edit!
################################################################################

# Each subdirectory must supply rules for building sources it contributes
eq.obj: ../eq.cla $(GEN_OPTS) $(GEN_HDRS)
	@echo 'Building file: $<'
	@echo 'Invoking: C2000 Compiler'
	"C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.2/bin/cl2000" -v28 -ml -mt --cla_support=cla1 --float_support=fpu32 --tmu_support=tmu0 --vcu_support=vcu2 --opt_for_speed=4 --fp_reassoc=on --fp_mode=relaxed --include_path="C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.2/include" --include_path="C:/SD/ProjectOmaha/hardware/Sound_Masking/include" -g --define=CPU1 --diag_warning=225 --display_error_number --diag_suppress=10068 --preproc_with_compile --preproc_dependency="eq.pp" --cla_support=cla1 $(GEN_OPTS__FLAG) "$<"
	@echo 'Finished building: $<'
	@echo ' '

main.obj: ../main.c $(GEN_OPTS) $(GEN_HDRS)
	@echo 'Building file: $<'
	@echo 'Invoking: C2000 Compiler'
	"C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.2/bin/cl2000" -v28 -ml -mt --cla_support=cla1 --float_support=fpu32 --tmu_support=tmu0 --vcu_support=vcu2 --opt_for_speed=4 --fp_reassoc=on --fp_mode=relaxed --include_path="C:/ti/ccsv6/tools/compiler/ti-cgt-c2000_6.4.2/include" --include_path="C:/SD/ProjectOmaha/hardware/Sound_Masking/include" -g --define=CPU1 --diag_warning=225 --display_error_number --diag_suppress=10068 --preproc_with_compile --preproc_dependency="main.pp" --cla_support=cla1 $(GEN_OPTS__FLAG) "$<"
	@echo 'Finished building: $<'
	@echo ' '


