
def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines

def isPossible(triangle):
    s = sum(triangle)
    m = max(triangle)
    return s - m > m

# star 1
def solve(tringleList):
    acc = 0
    for triangle in tringleList:
        if (isPossible(map(int, triangle.split()))):
            acc += 1
    return acc

# star 2
def solve_(tringleList):
    acc = 0
    triangles = [[],[],[]]
    lenlist = len(tringleList)

    for i in range(0, lenlist):
        if (i != 0 and i % 3 == 0):
            for t in triangles:
                if (isPossible(t)): acc += 1
            triangles = [[],[],[]]
        
        triangle = map(int, tringleList[i].split())

        triangles[0].append(triangle[0])
        triangles[1].append(triangle[1])
        triangles[2].append(triangle[2])

    for t in triangles:
        if (isPossible(t)): acc += 1

    return acc

lines = io().split('\n')

print solve(lines)
print solve_(lines)