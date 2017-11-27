import re

def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines

def checksum(s):
    di = {}

    for ch in s:
        di[ch] = di.get(ch, 0) + 1

    sortd = sorted(di.items(), key=lambda x: (-x[1], x[0]))
    return reduce((lambda acc, v: acc + v[0]), sortd, "")[:5]

def isRoom(name, chcksum):
    return checksum(name) == chcksum

def decrypt(name, id):
    r = ""

    for l in name:
        if (l == ' '): r += l
        else: r += chr((ord(l) - ord('a') + id) % 26 + ord('a'))

    return r

def solve(roomList, lookupRoom):
    realAcc = 0
    lookupSectorId = 0

    for room in roomList:
        parsed = re.split('[-\]\[]', room)

        name = ' '.join(parsed[:len(parsed)-3])
        chcksum = parsed[-2]
        sectorId = int(parsed[-3])

        if isRoom(name.replace(" ", ""), chcksum):
            # star 1
            realAcc += sectorId

            # star 2
            realName = decrypt(name, sectorId)
            if lookupRoom in realName: 
                lookupSectorId = sectorId

    return realAcc, lookupSectorId

sectorIdSum, npStorageId = solve(io().split('\n'), "northpole")

print sectorIdSum
print npStorageId

