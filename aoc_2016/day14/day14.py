import md5
import re

def normal(base, idx):
    m = base.copy()
    m.update(str(idx))
    return m.hexdigest()

def stretched(base, idx):
    h = normal(base, idx)
    for x in range(2016):
        h = md5.new(h).hexdigest()
    return h

def solve(salt, encode):
    threeOfKind = []
    keys = []
    x = 0

    base = md5.new(salt)

    while len(keys) < 64 and x < 1000000:
        digest = encode(base, x)
        
        foks = re.findall(r'(\w)\1{4}', digest)

        updated = []

        for tupl in threeOfKind:
            if tupl[1] in foks: keys.append(tupl)
            elif x - tupl[0] < 1000: updated.append(tupl)

        threeOfKind = updated

        tok = re.search(r'(\w)\1{2}', digest)
        if tok != None:
            threeOfKind.append((x, tok.group(1)))
        x += 1

    print len(sorted(keys))
    return sorted(keys)[:64][-1]

# print solve('abc')
print solve('cuanljph', normal)
print solve('cuanljph', stretched)