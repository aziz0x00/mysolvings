import sys

a = open(sys.argv[1]).readlines()
a.pop(0)
i = 1

while a:
    a.pop(0)
    print('Case #%d: %d' % (i, sum(int(i) for i in a.pop(0).split() if i[0] != '-' )))
    i += 1

