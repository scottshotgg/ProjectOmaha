#This class needs to be renamed. 

import cv2
import numpy as np
import math
import json
import argparse

def parseCircle(circle):
    x = int(circle.pt[0])
    y = int(circle.pt[1])
    cv2.circle(orig, (x, y), int(circle.size / 2), (0, 255, 0), 2)
    cv2.rectangle(orig, (x - 5, y - 5), (x + 5, y + 5), (0, 128, 255), -1)
    print(x + ", " + y)
    return {"x": x, "y": y}

# merge blobs that are intersecting
def supress(x):
    for f in fs:
        distx = f.pt[0] - x.pt[0]
        disty = f.pt[1] - x.pt[1]
        dist = math.sqrt(distx * distx + disty * disty)
        if (f.size > x.size) and (dist < f.size / 2):
            return True
                        
# check if executing this script
if __name__ == "__main__":
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
    lower_blue = np.array([98, 100, 100])
    upper_blue = np.array([130, 255, 255])
    mask = cv2.inRange(hsv, lower_blue, upper_blue)
    # get rid of the specks in the image
    mask = cv2.medianBlur(mask, 7)
    cv2.imwrite('mask.png', mask)
    
    # Here we need to check the version because the MSER object has changed between versions
    # You MAY need to install libgl1-mesa-dri
    
    version = cv2.__version__
    
    print("Using Open CV version: " + version)
    
    detector = 'null'
    if int(version[0]) < 3:
        detector = cv2.MSER()
    else:
        detector = cv2.MSER_create()

    fs = detector.detect(mask)
    print fs
    fs.sort(key = lambda x: -x.size)
    print fs
    
    sfs = [x for x in fs if not supress(x)]
    
    circleMap = {"circleLocations" : [parseCircle(x) for x in sfs]}
    
    outputFile.write(json.dumps(circleMap, indent=4))