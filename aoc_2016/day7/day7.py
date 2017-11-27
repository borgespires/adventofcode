import re

def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines

def isABBA(s):
    start = s[:2]
    end = s[2:][::-1]

    return start == end and start[0] != start[1]

def hasABBA(s):
    for i in xrange(0, len(s) - 3):
        if isABBA(s[i:i+4]): return True
    return False

def isABA(s): return s[0] == s[2] and s[0] != s[1]

def getABAs(s):
    listABAs = []
    for i in xrange(0, len(s) - 2):
        subs = s[i:i+3]
        if isABA(subs): listABAs.append(subs)
    return listABAs

def invert(aba):
    return aba[1] + aba[0] + aba[1]

c = 0
c2 = 0

for line in io().split('\n'):
    parsed = re.split('[\]\[]', line)

    # 1 star
    isValidABBA = False
    for i in range(0, len(parsed)):
        if i % 2 == 0:
            if isValidABBA: continue
            if hasABBA(parsed[i]): isValidABBA = True
        elif hasABBA(parsed[i]):
            isValidABBA = False
            break

    if isValidABBA:
        c += 1
        print "ABBA", line
    
    # 2 star
    abaList = []

    for i in range(0, len(parsed), 2): abaList += getABAs(parsed[i])
    for i in range(1, len(parsed), 2):
        if any(invert(aba) in parsed[i] for aba in abaList):
            c2 += 1
            print "ABA", line

print c
print c2