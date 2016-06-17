import os, sys

# Try making the directory
try:
	os.mkdir("./rawjs", 0755)
except OSError:
	print "dir already created."

# Walk the file structure to get all the files
walkvar = os.walk("/Users/scottgaydos/Documents/ProjectOmaha/software/src/omaha/templates/components/")

for files in walkvar:
	# For every file in files[2], which is a list of the files found in the pwd
	for file in files[2]:
		print file
		filename, fileextenstion = os.path.splitext(file)
		if fileextenstion == ".html":
			# Use sed to strip out the script and then pipe that into sed again and use it to trim it again
			command_string = "sed -n '/<script>/,/<\/script>/p' " + files[0] + "/" + file + " | sed '1d; $d' > ./rawjs/" + file + ".js"
			os.system(command_string)


# Make documentation file
os.system("documentation build ~/Documents/ProjectOmaha/software/src/omaha/Documentation/rawjs/ -f html --github -o ./")
#os.system("esdoc -c esdoc.json")