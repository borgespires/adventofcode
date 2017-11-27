# INPUT = 10
INPUT = 1352

def isWall(x, y):
    return bin((x*x + 3*x + 2*x*y + y + y*y) + INPUT).count('1') % 2 != 0

def neighbours(x, y):
    ns = [(x+1, y),(x, y+1)]

    if x > 0: ns.append((x-1, y))
    if y > 0: ns.append((x, y-1))

    return ns

def findpath(src, dest, maxDepth):
    q = [(0, src)]
    visited = [src]

    while len(q) > 0:
        dist, u = q.pop(0)

        # print dist, u

        # star 1
        if u == dest:
            print dist, u
            break

        # star 2
        if dist >= maxDepth > 0:
            print "reached %s and visited %s nodes" % (maxDepth, len(visited))
            break

        for n in neighbours(*u):
            if n not in visited and not isWall(*n):
                q.append((dist + 1, n))
                visited.append(n)

# findpath((1,1), (7,4))
findpath((1,1), (31,39), 0)
findpath((1,1), (31,39), 50)