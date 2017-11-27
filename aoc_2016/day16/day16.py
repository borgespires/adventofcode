
def fill(state, size):
    i = 0
    while len(state) < size:
        a = state
        b = ''.join([str(1 - int(v)) for v in reversed(state)])
        state = a + '0' + b
        i += 1

    return state[:size]

def checksum(state):
    c = ''
    for i in range(0, len(state), 2):
        if state[i] == state[i + 1]: c += '1'
        else: c += '0'

    if len(c) % 2 == 0: return checksum(c)

    return c


# optimezed version !!
def fill_(state, size):
    a = map(int, list(state))
    b = [1 - v for v in reversed(a)]
    a = a + [0] + b

    while len(a) < size: 
        a = a + [0] + a

        ## change the separator of the second half of the string
        ##                      iteration
        ## 11(0)00              0
        ## 11(0)00|0|11(1)00    1
        l = len(a)
        splitIndex = l/2 + l/4 + 1
        a[splitIndex] = 1 - a[splitIndex]

    return a[:size]

# print fill('11', 10)
# print checksum('110010110100')
# print checksum(fill('10000', 20))

# print checksum(fill('11011110011011101', 272))
# print checksum(fill_('11011110011011101', 35651584))