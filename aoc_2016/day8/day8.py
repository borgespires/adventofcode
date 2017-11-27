import re

def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines

h = 6
w = 50
screen = [[0]*w for i in range(h)]

def litPixels():
    r = 0
    for row in screen:
        for pixel in row:
            r += pixel
    return r

def rotateList(l, n):
    print "rotateList", l, n
    return l[-n:] + l[:-n]

def draw(x, y):
    for i in range(0, x):
        for j in range(0, y):
            screen[j][i] = 1

def rotate(cmd, index, amt):
    print cmd, index, amt
    if cmd == 'row':
        screen[index] = rotateList(screen[index], amt)
    else:
        column = [row[index] for row in screen]
        ncolumn = rotateList(column, amt)
        print "new column", ncolumn
        for i in range(0, len(ncolumn)):
            screen[i][index] = ncolumn[i]


for line in io().split('\n'):
    matchObj = re.match(r'(?:.*)(row|column|rect) (?:(?:[xy]=(\d+) by (\d+))|(?:(\d+)x(\d+)))$', line, re.M|re.I)
    if matchObj:
        cmd = matchObj.group(1)
        if cmd == "rect":
            draw(int(matchObj.group(4)), int(matchObj.group(5)))
        else:
            rotate(cmd, int(matchObj.group(2)), int(matchObj.group(3)))
    else:
       print "No match!!"

print litPixels()

for row in screen:
    print row