
def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines

class Pad():
    def __init__(self, layout, starting):
        self.location = starting
        self.layout = layout

    def move(self, cmd):
        if (cmd == 'U'): p = (self.location[0], self.location[1] - 1)
        elif (cmd == 'D'): p = (self.location[0], self.location[1] + 1)
        elif (cmd == 'L'): p = (self.location[0] - 1, self.location[1])
        elif (cmd == 'R'): p = (self.location[0] + 1, self.location[1])

        if self._isValid(p):
            self.location = p

    def _isValid(self, p):
        if (p[0] < 0 or p[1] < 0 or p[0] >= len(self.layout) or p[1] >= len(self.layout)):
            return False

        key = self.layout[p[1]][p[0]]
        return key != '_'

    def current(self):
        return self.layout[self.location[1]][self.location[0]]

def solve(instructions, layout, starting):
    pad = Pad(layout, starting)
    answer = ''
    for moves in instructions:
        for m in moves:
            pad.move(m)
        answer += pad.current()
    return answer

lines = io().split('\n')

pad1 = [
    ['1','2','3'],
    ['4','5','6'],
    ['7','8','9']
]

pad2 = [
    ['_', '_', '1', '_', '_'],
    ['_', '2', '3', '4', '_'],
    ['5', '6', '7', '8', '9'],
    ['_', 'A', 'B', 'C', '_'],
    ['_', '_', 'D', '_', '_']
]

# star 1
print solve(lines, pad1, (1,2))

# star 2
print solve(lines, pad2, (0,2))