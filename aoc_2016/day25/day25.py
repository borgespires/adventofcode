
def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines

def searchPattern(s):
    PATTERN = ['inc', 'dec', 'jnz', 'dec', 'jnz']
    expect = 0
    found = []
    for i in range(0, len(s)):
        if expect == 5: #found
            d = s[i - 5].split()[1]
            m1 = s[i - 4].split()[1]
            m2 = s[i - 2].split()[1]
            found.append((i - 5, m1, m2, d))
            expect = 0

        cmd = s[i]
        if PATTERN[expect] in cmd: expect += 1

    return found

def optimize(s):
    optimizations = {}
    found = searchPattern(s)

    for i, m1, m2, d in found:
        optimizations[i] = (m1, m2, d)

    return optimizations

flag = 0
starti = 1
counter = 0

# if check fails restart machine
def check(x):
    global starti
    global counter
    global flag
    global pc
    global memory

    if x != flag:
        print "fail for ", starti, counter, (x,flag)
        starti += 1
        flag = 0
        memory = { 'a': starti, 'b': 0, 'c': 1, 'd': 0 }
        pc = -1
        counter = 0
    else:
        flag = (flag + 1) % 2
        counter += 1

    if counter > 1000:
        print "result:", starti
        pc = len(stack)

def evaluate(cmd):
    global pc
    parts = cmd.split()
    op = parts[0]

    if op == 'cpy':
        x = parts[1]
        y = parts[2]

        if y.lstrip('-').isdigit(): return
        
        if x.lstrip('-').isdigit(): memory[y] = int(x)
        else: memory[y] = memory[x]

    if op == 'inc':
        x = parts[1]
        memory[x] += 1

    if op == 'dec':
        x = parts[1]
        memory[x] -= 1

    if op == 'jnz':
        x = parts[1]
        y = parts[2]

        if x.isdigit(): w = int(x)
        else: w = memory[x]

        if y.lstrip('-').isdigit(): z = int(y)
        else: z = memory[y]
        
        if w != 0 and pc != 0: pc += int(z) - 1
        if pc < 0: pc += len(stack)

    if op == 'mult':
        x = parts[1]
        y = parts[2]
        d = parts[3]

        if d.lstrip('-').isdigit(): return

        if x.isdigit(): w = int(x)
        else: w = memory[x]

        if y.lstrip('-').isdigit(): z = int(y)
        else: z = memory[y]

        memory[d] += w * z

    if op == 'out':
        x = parts[1]
        if x.isdigit(): z = int(x)
        else: z = memory[x]
        check(z)

pc = 0
stack = io().split('\n')
optimizations = optimize(stack)

memory = { 'a': starti, 'b': 0, 'c': 1, 'd': 0 }

while pc < len(stack):
    opt = optimizations.get(pc)
    if opt:
        cmd = 'mult %s %s %s' % (opt[0], opt[1], opt[2])
        step = 5
    else:
        cmd = stack[pc]
        step = 1
    
    evaluate(cmd)
    # print cmd, memory
    pc += step

print memory