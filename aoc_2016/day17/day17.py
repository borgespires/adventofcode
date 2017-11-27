import md5

OPEN_CODES =['b', 'c', 'd', 'e', 'f']
# PASS_CODE = 'hijkl'
# PASS_CODE = 'ihgpwlah'
PASS_CODE = 'njfxhljp'

base = md5.new(PASS_CODE)

def openedDoors(path):
    m = base.copy()
    m.update(path)
    return m.hexdigest()[:4]

q = []
q.append(((0, 0), ''))

longest = 0

while len(q) > 0:
    (x, y), path = q.pop(0)
    u, d, l, r = openedDoors(path)

    if x == 3 and y == 3:
        if len(path) > longest: longest = len(path)
        # print 'Final path:', path
        # break
        continue

    if y > 0 and u in OPEN_CODES: q.append(((x, y-1), path + 'U'))
    if y < 3 and d in OPEN_CODES: q.append(((x, y+1), path + 'D'))
    if x > 0 and l in OPEN_CODES: q.append(((x-1, y), path + 'L'))
    if x < 3 and r in OPEN_CODES: q.append(((x+1, y), path + 'R'))

print longest