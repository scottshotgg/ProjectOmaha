import cv2
import numpy as np
import math
import json
import argparse

# parse input parameters
parser = argparse.ArgumentParser(description="Get speaker coordinates from map")
parser.add_argument('imagePath', metavar="I")
parser.add_argument('outputPath', metavar="O", type=argparse.FileType('w'))
parsedInput = vars(parser.parse_args())
outputFile = parsedInput["outputPath"]
inputFileName = parsedInput["imagePath"]

orig = cv2.imread(inputFileName)
origCopy = orig.copy()
hsv = cv2.cvtColor(origCopy, cv2.COLOR_BGR2HSV)
# define range of blue color in HSV
lower_blue = np.array([98,100,100])
upper_blue = np.array([130,255,255])
# lower_blue = np.array([0,100,100])
# upper_blue = np.array([30,255,255])
mask = cv2.inRange(hsv, lower_blue, upper_blue)
mask = cv2.medianBlur(mask,7)

# Here we need to check the version because the MSER object has changed has changed
# You MAY need to install libgl1-mesa-dri

version = cv2.__version__

print("Using Open CV version: " + version)

if int(version[0]) < 3:
    detector = cv2.MSER()
else:
    detector = cv2.MSER_create()

fs = detector.detect(mask)
fs.sort(key = lambda x: -x.size)

# merge blobs that are intersecting
def supress(x):
        for f in fs:
                distx = f.pt[0] - x.pt[0]
                disty = f.pt[1] - x.pt[1]
                dist = math.sqrt(distx*distx + disty*disty)
                if (f.size > x.size) and (dist<f.size/2):
                        return True

sfs = [x for x in fs if not supress(x)]


def parseCircle(circle):
    x = int(circle.pt[0])
    y = int(circle.pt[1])
    cv2.circle(orig, (x, y), int(circle.size/2), (0, 255, 0), 2)
    cv2.rectangle(orig, (x - 5, y - 5), (x + 5, y + 5), (0, 128, 255),-1)
    return {"x": x, "y": y}
circleMap = {"circleLocations" : [parseCircle(x) for x in sfs]}

outputFile.write(json.dumps(circleMap, indent=4))
