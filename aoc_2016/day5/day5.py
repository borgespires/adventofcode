import md5

LOWER_BOUND = 3231928
UPPER_BOUND = 10000000000
INPUT = 'uqwqemis'

def solve(doorId):
    base = md5.new()
    base.update(doorId)

    r = ''
    count = 0

    for x in xrange(LOWER_BOUND, UPPER_BOUND):
        test = base.copy()
        test.update(str(x))
        
        digest = test.hexdigest()

        if sum(map((lambda x: 0 if x == '0' else 1), digest[:5])) == 0:
            r += digest[5]
            count += 1
            print  count, digest, r, x

            if count >= 8:
                return r

def solve_(doorId):
    base = md5.new()
    base.update(doorId)

    r = ['_', '_', '_', '_', '_', '_', '_', '_']
    count = 0

    for x in xrange(LOWER_BOUND, UPPER_BOUND):
        test = base.copy()
        test.update(str(x))
        
        digest = test.hexdigest()

        if sum(map((lambda x: 0 if x == '0' else 1), digest[:5])) == 0:
            index = digest[5]

            if ord('7') >= ord(index) >= ord('0') and r[int(index)] == '_':
                r[int(index)] = digest[6]
                count += 1
                print  count, digest, r

                if count >= 8:
                    return ('').join(r)

print solve(INPUT)
print solve_(INPUT)