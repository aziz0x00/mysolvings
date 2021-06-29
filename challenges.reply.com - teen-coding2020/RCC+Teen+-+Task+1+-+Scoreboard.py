import sys

a = open(sys.argv[1])
i = int(a.readline().strip())
I = i


class Team:

    def __init__(self, k):
        self.id = k
        self.score = 0
        self.ts = 0
        self._solves = []

    def submit(self, ts, solve):
        if solve in self._solves: # :)
            return
        self._solves.append(solve)
        self.score += solve[1]
        self.ts += ts

    def __gt__(self, other):
        if self.score > other.score:
            return True
        elif self.score < other.score:
            return False
        else:
            if self.ts < other.ts:
                return True
            elif self.ts > other.ts:
                return False
            else:
                return self.id < other.id

    def __str__(self):
        return str(self.id)

    def __repr__(self):
        return '<Team#%d(s=%d, t=%d)>' % (self.id, self.score, self.ts)

while i:

    N, l = map(int, a.readline().strip().split())

    OUT = [Team(_) for _ in range(1,N+1)]

    for _ in range(l):
        L =  [int(_) for _ in a.readline().strip().split()]
        if not L[-1]:
            continue

        OUT[L[1]-1].submit(L[0], [L[2], L[3]])

    # print(OUT)
    OUT.sort(reverse=True)
    # print(OUT)

    print('Case #%d: %s' % (I-i+1, ' '.join(map(str,OUT))))
    i -= 1
