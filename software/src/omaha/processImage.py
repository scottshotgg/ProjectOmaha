import cv2
import numpy as np
import math

orig = cv2.imread("lots.png")
origCopy = orig.copy()
hsv = cv2.cvtColor(origCopy, cv2.COLOR_BGR2HSV)
# define range of blue color in HSV
lower_blue = np.array([98,100,100])
upper_blue = np.array([120,255,255])
mask = cv2.inRange(hsv, lower_blue, upper_blue)
mask = cv2.medianBlur(mask,7)
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

for f in sfs:
    x = int(f.pt[0])
    y = int(f.pt[1])
    cv2.circle(orig, (x, y), int(f.size/2), (0, 255, 0), 2)
    cv2.rectangle(orig, (x - 5, y - 5), (x + 5, y + 5), (0, 128, 255),-1)


cv2.imwrite("result.png", orig)
