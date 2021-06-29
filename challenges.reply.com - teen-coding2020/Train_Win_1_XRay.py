import sys

a = open(sys.argv[1]).readlines()
a.pop(0)
iiii = 1

def u(l):
    k = 0
    o = []

    for i in l:
        if k == i:
            continue

        o.append(i)
        k = i

    if not len(o):
        return [0]

    if o[-1] == 0:
        o = o[:-1]

    return o

def x(i):

    if 0 in i:
        if len(i) == 1:
            return 0

        return sum(x(u(int(j) for j in j.strip(',').split(','))) for j in str(i)[1:-1].split(' 0,'))

    return x(u(j - min(i) for j in i)) + min(i)

while a:
    a.pop(0)
    print('Case #%d: %d' % (iiii, x(u(int(j) for j in a.pop(0).split()))))
    iiii += 1
