
def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines

def findLowest(blacklist):
    lowest = 0
    for (start, end) in blacklist:
        if lowest < start: return lowest
        elif lowest < end: lowest = end + 1

def countAllowed(blacklist):
    upperBound = 0
    alowed = 0
    for (start, end) in blacklist:
        if upperBound < start: 
            alowed += start - upperBound
        if upperBound < end: upperBound = end + 1

    return alowed

# blacklist = [(5, 8), (0, 2), (4, 7)]
blacklist = []

for line in io().split('\n'):
    s, e = line.split('-')
    blacklist.append((int(s), int(e)))

blacklist = sorted(blacklist)

# print findLowest(blacklist)
print countAllowed(blacklist)