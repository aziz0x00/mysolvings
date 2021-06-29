import sys

# this one cost me time <3
# ive seen the binsearch solution by one of the winners, and it was sooo lit *.*
#                           it was small efficient clean code ^^, unlike mine xd
# UPDATE: yeyyyyy i wrote another coooooooooooool code!!! <33333333

a = open(sys.argv[1])
TEST_CASES = int(a.readline())

"""
class Serv:
 
    def __init__(self, p, s):
        self.p = int(p)
        self.s = int(s)

for t in range(1, TEST_CASES+1):
    N, K, M = [int(i) for i in a.readline().split()]

    S = [Serv(*a.readline().split()) for _ in range(N)]

    hi = min(S[i].s*M + S[i].p for i in range(N)) 
    lo = min(S[i].s + S[i].p for i in range(N))

    while hi != lo:
        c = (hi + lo) // 2

        if sum(sorted((c - S[i].p)//S[i].s for i in range(N))[-K:]) >= M:
            hi = c
        elif lo == c:
            c = lo = hi
        else:
            lo = c

    print('Case #%d: %d' % (t, c))
"""

# OLD CODE: (im too busy to fix it :/, i wish i could do it.. im way way too busy, d3ouli moulane tewfi9, i wish it helps.) <3
class Server:

    def __init__(self, i, p, s):
        self.id = i
        self.time = p + s # 1-task time
        self.s = s
        self._m = 1
        self.p = p

    @property
    def m(self):
        return self._m

    @m.setter
    def m(self, v):
        self.time = self.p + self.s * v
        self._m = v

    def __lt__(self, other):
        return self.time < other.time

    def __repr__(self):
        return "S#%d<p=%d,s=%s,m=%d>" % (self.id, self.p, self.s, self.m)

for t in range(1, TEST_CASES+1):
    N, K, M = [int(i) for i in a.readline().split()]

    S = [Server(i, *map(int, a.readline().split())) for i in range(N)]
    S.sort()
    l = 0
    _m = []
    # stage 1/2
    while True:

        m = [(S[l+1].time - S[i].p) // S[i].s for i in range(l+1)]

        # if/else to update S by new m
        if _m != [] and _m == m[:-1]:
            S[l].m = m[l]
        else:
            for Si, mi in zip(S, m):
                Si.m = mi

        # gtfo if we got >=M
        if (_s:=sum(m, S[l+1].m if l != K-1 else 0)) >= M: break

        # prepare vars for next loop
        l += 1
        if l == K:
            for i in range(N):
                S[i].m += 1      # i think here lies the problem!
            S.sort()
            l = 0
            _m = []
        else:
            _m = m

    # stage 2/2
    s = _s - (S[l+1].m if l != K-1 else 0)
    if s < M:
        print('yey')
        print(S[l+1])
        c = S[l+1].time
    elif l != 0:
        S = S[:l+1]
        i = S.index(sorted(S, key=lambda s: s.s)[0])
        hi = S[i].m+1
        lo = m[i]
        while hi != lo:
            m = (hi + lo) // 2
            if sum(((S[i].s*m+S[i].p-S[j].p)//S[j].s\
                    for j in range(l+1) if j != i), m) >= M:
                hi = m
            elif lo == m:
                m = lo = hi
            else:
                lo = m
        c = S[i].s*m+S[i].p
    else:
        c = S[0].p + S[0].s * M

    print('Case #%d: %d' % (t, c))

