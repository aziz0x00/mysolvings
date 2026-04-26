import sys

a = open(sys.argv[1]).readlines()
a.pop(0)
iii = 1

while a:
    N = [int(x) for x in a.pop(0).split()]
    minN = min(N)
    M = int(a.pop(0))
    F = (a.pop(0), a.pop(0), a.pop(0), a.pop(0))
    OUT = []
    for i in range(minN):
        s = F[N.index(minN)][i:i+M]
        for Fi in F:
            try:
                OUT.append(str(Fi.index(s)))
            except:
                OUT = []
                break
        else:
            break


    print('Case #%d: %s' % (iii, ' '.join(OUT)))
    iii += 1