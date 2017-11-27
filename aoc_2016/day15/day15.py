import re

def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines

def droppedCapsule(stime):
    time = stime + 1
    
    for p, tp in discs:
        # print time, (p, tp), (time + p) % tp
        if (time + p) % tp != 0:
            return False
        time += 1
    return True


# discs = [(4, 5), (1, 2)]

# star 1
discs = []
stime = 0

for line in io().split('\n'):
    m = re.search(r'(\d+) positions.*position (\d+)', line)
    discs.append((int(m.group(2)), int(m.group(1))))

# star 2
discs.append((0, 11))

while stime < 10000000:
    if droppedCapsule(stime):
        print stime
        break
    stime += 1