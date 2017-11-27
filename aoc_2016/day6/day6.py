
def io():
    fo = open("input.txt", "r+")
    lines = fo.read();
    fo.close()

    return lines


lines = io().split('\n')

length = len(lines[0])
counter =  [{} for k in range(length)]

for line in lines:
    i = 0
    for ch in line:
        di = counter[i]
        di[ch] = di.get(ch, 0) + 1
        i += 1

r = ""
for di in counter:
    # 1 star
    # r += sorted(di.items(), key=lambda x: x[1], reverse = True)[0][0]

    # 2 star
    r += sorted(di.items(), key=lambda x: x[1], reverse = False)[0][0]

print r