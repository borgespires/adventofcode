import re
from itertools import combinations
from copy import deepcopy

import time
start_time = time.time()

def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines

def getInitialState():
    # state = [[], [], [], []]
    # elements = 0
    state = [['EM', 'EG', 'DM', 'DG'], [], [], []]
    elements = 4

    i = 0
    for line in io().split('\n'):
        m = re.findall(r'(:?\w+)(?:\s|-)(generator|\w+ microchip)', line, re.M|re.I)

        for ele in m:
            id = ele[0][0]+ele[1].split()[-1][0]
            state[i].append(id.upper())
            elements += 1

        print i, state[i]
        i += 1

    return state, elements

def getHash(e, state):
    n = []
    for floor in state:
        joined = ''.join(floor)
        n.append(str(joined.count('G')) + 'G' + str(joined.count('M')) + 'M')
    return str(e) + '|' + '|'.join(n)

def isValid(state):
    for floor in state:
        chips = filter(lambda x: x[1] == 'M', floor)
        rtgs = filter(lambda x: x[1] == 'G', floor)

        if len(chips) > 0 and len(rtgs) > 0:
                for chip in chips: 
                    if chip[0] + 'G' not in rtgs: return False
    return True

def addChildNodes(parent, src, dst, itemList, depth):
    moved = False

    for items in itemList:
        state = deepcopy(parent)
        state[src] = [i for i in state[src] if i not in items]
        state[dst] = state[dst] + items

        sHash = getHash(dst, state)

        if isValid(state) and sHash not in visited:
            q.append((dst, state, depth + 1))
            visited.append(sHash)
        moved = True

    return moved

def nextStates(e, state, depth):
    global q
    
    floor = state[e] # current floor
    uniq = [[x] for x in floor]
    comb = map(list, combinations(floor, 2))

    # improve
    # if i can move one downstairs don't bring two
    if e > 0 and not addChildNodes(state, e, e-1, uniq, depth):
        addChildNodes(state, e, e-1, comb, depth)
    # if i can move two items up don't bring only one
    if e < 3 and not addChildNodes(state, e, e+1, comb, depth): 
        addChildNodes(state, e, e+1, uniq, depth)

    # alll = uniq + comb
    # for items in alll:
    #     
    #     if e > 0: addState(deepcopy(state), e, e-1, items, depth + 1)
    #     if e < 3: addState(deepcopy(state), e, e+1, items, depth + 1)


# main()
q = []
visited = []
initialState, N_ELEMENTS = getInitialState()

# start processing with bfs
q.append((0, initialState, 0))
visited.append(getHash(0, initialState))

while len(q) > 0:
    e, state, d = q.pop(0)
    # print "(%s) - E:%s - %s - S3(%s) - queue(%s)" % (d, e, state, len(state[3]), len(q))

    if len(state[3]) == N_ELEMENTS:
        print d
        break;
    nextStates(e, state, d)

print("--- %s seconds ---" % (time.time() - start_time))
