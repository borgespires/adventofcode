import re

def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines

def findMarker(data, start):
    r = ''
    i = start + 1
    while data[i] != ')' and i < len(data):
        r += data[i]
        i += 1
    return r

def decrypt_v1(data):
    i=0
    dcrypt = 0

    while i < len(data):
        ch = data[i]
        if ch == '(':
            marker = findMarker(data, i)

            markerSize = len(marker) + 2
            contentSize, multiplier = map(int, re.split('[x]', marker)) # 2x3 split
            
            dcrypt += contentSize * multiplier
            i += markerSize + contentSize
        else:
            dcrypt += 1
            i += 1

    return dcrypt

def decrypt_v2(data):
    i=0
    dcrypt = 0

    while i < len(data):
        ch = data[i]
        if ch == '(':
            marker = findMarker(data, i)

            markerSize = len(marker) + 2
            contentSize, multiplier = map(int, re.split('[x]', marker))

            contentStart = i + markerSize
            contentEnd = contentStart + contentSize
            content = data[contentStart:contentEnd]
            
            dcrypt += decrypt_v2(content) * multiplier
            i += markerSize + contentSize
        else:
            dcrypt += 1
            i += 1

    return dcrypt
 
total1 = 0
total2 = 0
for s in io().split('\n'):
    total1 += decrypt_v1(s.replace(" ", ""))
    total2 += decrypt_v2(s.replace(" ", ""))

print total1
print total2