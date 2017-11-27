
def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines

unsafeTiles = [ '^^.', '.^^', '^..', '..^']

def isSafe(tiles):
    return 

def findTraps(row, start, end):
    nrow = ''
    l = len(row)

    for x in range(l):
        s = x + start
        e = x + end + 1

        tiles = ''

        if s < 0:
            s = 0
            tiles += '.'

        tiles += row[s:e]

        if e > l: tiles += '.'

        if tiles not in unsafeTiles: nrow += '.'
        else: nrow += '^'
    return nrow

prevRow = ''
nextRow = io()

safeCount = 0

for x in xrange(400000):
    prevRow = nextRow

    # print prevRow
    nextRow = findTraps(prevRow, -1, 1)
    safeCount += prevRow.count('.')


print safeCount

# print findTraps('..^^.', -1, 2)