import os, sys

try:
	os.mkdir("./rawjs", 0755)
except OSError:
	print "dir already created."

walkvar = os.walk("/Users/scottgaydos/Documents/ProjectOmaha/software/src/omaha/templates/components/")

for files in walkvar:
		#print "\n\n\n", files[0], "\n\n\n"
		#print files[2], "\n\n\n"
		for file in files[2]:
			print file
			filename, fileextenstion = os.path.splitext(file)
			if fileextenstion == ".html":
				command_string = "sed -n '/<script>/,/<\/script>/p' " + files[0] + "/" + file + " | sed '1d; $d' > ./rawjs/" + file + ".js"
				os.system(command_string)