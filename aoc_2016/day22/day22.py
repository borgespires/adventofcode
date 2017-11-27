import re
import heapq

def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()
    return lines

class Node(object):
    def __init__(self, size, used):
        self.size = size
        self.used = used
        self.valid = True

    def __str__(self):
        return "size: %s | used: %s" % (self.size, self.used)


def getViablePairs(cluster):
    c = sum(cluster, [])
    viable = []

    for i in range(0, len(c)):
        for j in range(i+1, len(c)):
            n1 = c[i]
            n2 = c[j]

            if n1.used != 0 and n1.used <= n2.size - n2.used: viable.append((n1, n2))
            if n2.used != 0 and n2.used <= n1.size - n1.used: viable.append((n2, n1))

    return viable

cluster = []

for line in io().split('\n'):
    m = re.search(r'x(\d{1,2})-y(\d{1,2}).+?(\d{1,3})T.+?(\d{1,3})T.+?(\d{1,3})T.+?(\d{1,3})%', line)

    x = int(m.group(1))
    y = int(m.group(2))
    size = int(m.group(3))
    used = int(m.group(4))
    avail = int(m.group(5))
    usepc = int(m.group(6))

    if len(cluster) <= y: cluster.append([])

    cluster[y].append(Node(size, used))

# 1 star
print len(getViablePairs(cluster))

# 2 star
# Making a few assumptions:
#       there's one empty node
#       there are no mergeable nodes
#       there aren't enough walls in the top few rows to cause a degenerate case

def getEmptyNode(cluster):
    for y in range(len(cluster)):
        for x in range(len(cluster[y])):
            if cluster[y][x].used == 0:
                return (x, y)


def getNeighbors(x, y, maze):
    ns = []

    if x > 0: ns.append((x-1, y))
    if y > 0: ns.append((x, y-1))
    if x < len(maze[0])-1: ns.append((x+1, y))
    if y < len(maze)-1: ns.append((x, y+1))

    used = maze[y][x].used

    return [n for n in ns if maze[int(n[1])][int(n[0])].size >= used and maze[int(n[1])][int(n[0])].valid]

def astar(src, dest, maze):
    q = []
    visited = []

    h = abs(src[0] - dest[0]) + abs(src[1] - dest[1])
    f = h
    heapq.heappush(q, (f, src, []))

    while len(q) > 0:
        f, pos, g = heapq.heappop(q)

        if pos in visited: continue
        
        visited.append(pos)

        if pos == dest: return g

        for n in getNeighbors(pos[0], pos[1], maze):
            if n not in visited:
                h = abs(n[0] - dest[0]) + abs(n[1] - dest[1])
                f = h + len(g) + 1
                heapq.heappush(q, (f, n, g + [n]))


emptyNode = getEmptyNode(cluster)
goal = (len(cluster[0])-1, 0)
pathTopRight = astar(goal, (0,0), cluster)

s = 0
i = 0

while goal != (0,0):
    s += len(astar(emptyNode, pathTopRight[i], cluster))
    emptyNode = goal
    cluster[emptyNode[1]][emptyNode[0]].valid = True
    goal = pathTopRight[i]
    cluster[goal[1]][goal[0]].valid = False
    i += 1

print s + len(pathTopRight)


