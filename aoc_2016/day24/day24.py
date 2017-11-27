import heapq

def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines

def loadMaze():
    maze = []
    beacons = []
    
    lines = io().split('\n')

    for y in range(0, len(lines)):
        row = lines[y]
        maze.append(list(row))

        for x in range(0, len(row)):
            tile = row[x]
            if tile.isdigit(): beacons.append((tile, (x, y)))
    return maze, sorted(beacons)

def getNeighbors(x, y, maze):
    ns = [(x+1, y),(x, y+1),(x-1, y),(x, y-1)]

    return [n for n in ns if maze[int(n[1])][int(n[0])] != '#']

def astar(src, dest, maze):
    q = []
    visited = []

    h = abs(src[0] - dest[0]) + abs(src[1] - dest[1])
    f = h
    heapq.heappush(q, (f, src, 0))

    while len(q) > 0:
        f, pos, g = heapq.heappop(q)

        if pos in visited: continue
        
        visited.append(pos)

        if pos == dest: return g

        for n in getNeighbors(pos[0], pos[1], maze):
            if n not in visited:
                h = abs(n[0] - dest[0]) + abs(n[1] - dest[1])
                f = h + g + 1
                heapq.heappush(q, (f, n, g+1))

def createGraph(maze, beacons):
    graph = []

    for i in range(0, len(beacons)):
        b0, p0 = beacons[i]

        for j in range(i+1, len(beacons)):
            b1, p1 = beacons[j]

            steps = astar(p0, p1, maze)

            graph.append((steps, (b0, b1)))
    return sorted(graph)

def childnodes(beacon, graph):
    children = []
    
    for cost, (n1, n2) in graph:
        if n1 == beacon: children.append((n2, cost))
        elif n2 == beacon: children.append((n1, cost))

    return children

def goHomeCost(beacon, graph):
    for cost, (n1, n2) in graph:
        if n1 == '0' and n2 == beacon: return cost

def bfs(graph, nBeacons):
    q = []

    heapq.heappush(q, (0, '0', []))

    while len(q) > 0:
        cost, beacon, visited = heapq.heappop(q)

        v = list(visited)

        if beacon not in v: v = v + [beacon]
        elif len(v) < nBeacons: continue

        # 1 star
        # if len(v) < nBeacons:
        #     for b, c in childnodes(beacon, graph):
        #         heapq.heappush(q, (cost + c, b, v))
        # else:
        #     return cost

        # 2 star
        if len(v) == nBeacons:
            c = goHomeCost(beacon, graph)
            heapq.heappush(q, (cost + c, '0', v + ['0']))
        elif len(v) < nBeacons:
            for b, c in childnodes(beacon, graph):
                heapq.heappush(q, (cost + c, b, v))
        else:
            return cost


maze, beacons = loadMaze()
# print beacons
graph = createGraph(maze, beacons)
# print graph
print bfs(graph, len(beacons))

#_____________________________

s = []
connected = [['0']]

def find(n):
    for i in range(0,len(connected)):
        if n in connected[i]: return i

def add(n1, n2, size):
    global connected
    global s
    subtree1 = find(n1)
    subtree2 = find(n2)

    if subtree1 == None and subtree2 == None:
        connected.append([n1, n2])
    elif subtree1 != None and subtree2 != None:
        if subtree1 != subtree2:
            print connected
            merged = connected[subtree1] + connected[subtree2]
            del connected[max(subtree1, subtree2)]
            del connected[min(subtree1, subtree2)]
            connected.append(merged)
        else:
            return 0
    elif subtree1 != None and subtree2 == None:
        connected[subtree1].append(n2)
    elif subtree1 == None and subtree2 != None:
        connected[subtree2].append(n1)

    s.append(((n1, n2), size))
    return size

def findMST(graph):
    total = 0
    for size, (n1, n2) in graph:
        total += add(n1, n2, size)

    return s, total

# print findMST(graph)
