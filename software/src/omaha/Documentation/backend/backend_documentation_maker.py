import os

file_array = ["database", "handlers", "main", "system", "util"]

for file in file_array:
	command_string = "godoc -html ../../" + file + " > " + file + ".html"
	os.system(command_string)