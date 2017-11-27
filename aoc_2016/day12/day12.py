import re

def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines

def evaluate(cmd):
    global pc
    parts = cmd.split()
    op = parts[0]

    if op == 'cpy':
        x = parts[1]
        y = parts[2]
        
        if x.isdigit(): memory[y] = int(x)
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

        if x.isdigit(): z = int(x)
        else: z = memory[x]
        
        if z != 0 and pc != 0: pc += int(y) - 1
        if pc < 0: pc += len(stack)

pc = 0
stack = io().split('\n')
memory = { 'a': 0, 'b': 0, 'c': 1, 'd': 0 }

while pc < len(stack):
    cmd = stack[pc]
    evaluate(cmd)
    pc += 1

print memory