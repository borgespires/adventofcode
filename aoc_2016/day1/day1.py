
def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines

class Compass():
    def __init__(self):
        self.location = (0,0)
        self.direction = (0,1)

    def move(self, direction, distance):
        if (direction == 'R'): self._rotateR()
        if (direction == 'L'): self._rotateL()
        self._move(self.direction[0] * distance, self.direction[1] * distance)

    def manhattan(self):
        return abs(self.location[0]) + abs(self.location[1])
        
    def _rotateR(self): 
        x, y = self.direction
        self.direction = (y, -x)

    def _rotateL(self): 
        x, y = self.direction
        self.direction = (-y, x)

    def _move(self, x, y):
        lx, ly = self.location
        self.location = (lx + x, ly + y)

# star 1
def solve(moves):
    compass = Compass()
    for m in moves:
        compass.move(m[0], int(m[1:]))
    return compass.manhattan()

# star 2
def solve_(moves):
    compass = Compass()
    visited = []

    for m in moves:
        compass.move(m[0], 0)
        
        for i in range(0, int(m[1:])): 
            compass.move(None, 1)
        
            location = compass.location

            if location in visited: 
                return compass.manhattan()
            else: 
                visited.append(location)

moves = io().split(', ')
print solve(moves)
print solve_(moves)