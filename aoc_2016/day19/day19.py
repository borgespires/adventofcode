import math

# 1 star
def solve(x):
    return 2 * (x - pow(2, int(math.log(x, 2)))) + 1

# 2 star
def solve_(x):
    n = pow(3, int(math.log(x, 3)))
    a = 2 * n - x
    return (2 * x - 3 * n) + (a + abs(a)) / 2

def s(arr, x):
    i = 0
    while len(arr) > 1:
        d = len(arr)/2
        del arr[d]
        arr.append(arr.pop(0))
    print x, arr


print int(solve(3001330))
print int(solve_(3001330))

# for x in xrange(2, 100):
#     s(range(1,x), x-1)
