import re

def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines

class Arr():
    def __init__(self, p):        
        self.passwd = list(p)

    def swap(self, x, y):
        self.passwd[x], self.passwd[y] = self.passwd[y], self.passwd[x]

    def swapl(self, a, b):
        x = self.passwd.index(a)
        y = self.passwd.index(b)
        self.swap(x, y)

    def rotate(self, n, d):
        self.passwd = self.passwd[d*n:] + self.passwd[:d*n]

    def rotateb(self, a, d):
        i = self.passwd.index(a)
        if d < 0:
            n = 1 + i + (1 if i >= 4 else 0)
        else:
            n = 1 + i/2 + (4 if i % 2 == 0 and i != 0 else 0)

        self.rotate(n % len(self.passwd), d)

    def reverse(self, x, y):
        self.passwd[x:y+1] = reversed(self.passwd[x:y+1])

    def move(self, x, y):
        l = self.passwd[x]
        del self.passwd[x]
        self.passwd = self.passwd[:y] + [l] + self.passwd[y:]

    def __str__(self):
        return ''.join(self.passwd)

def scramble(cmd, undo):
    global arr

    if 'swap position' in cmd:
        m = re.search(r'(\d).+(\d)', cmd)
        x = int(m.group(1))
        y = int(m.group(2))
        arr.swap(x, y)
    if 'swap letter' in cmd:
        m = re.search(r'letter (\w) with letter (\w)', cmd)
        x = m.group(1)
        y = m.group(2)
        arr.swapl(x, y)
    if 'rotate' in cmd:
        if 'based' in cmd:
            m = re.search(r'letter (\w)', line)
            x = m.group(1)
            
            if undo: arr.rotateb(x, 1)
            else: arr.rotateb(x, -1)
        else:
            m = re.search(r'rotate (.+) (\d)', line)
            d = 1 if m.group(1) == 'left' else -1
            n = int(m.group(2))

            if undo: d = d * -1

            arr.rotate(n, d)
    if 'reverse' in cmd:
        m = re.search(r'(\d).+(\d)', cmd)
        x = int(m.group(1))
        y = int(m.group(2))
        arr.reverse(x, y)
    if 'move' in cmd:
        m = re.search(r'(\d).+(\d)', cmd)
        x = int(m.group(1))
        y = int(m.group(2))
        
        if undo: arr.move(y, x)
        else: arr.move(x, y)

# arr = list('abcde')
arr = Arr('abcdefgh')

for line in io().split('\n'):
    scramble(line, False)
    # print arr

print arr

arr = Arr('fbgdceah')

for line in reversed(io().split('\n')):
    scramble(line, True)
    # print arr

print arr