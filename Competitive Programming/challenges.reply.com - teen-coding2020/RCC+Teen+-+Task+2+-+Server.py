import sys

a = open(sys.argv[1])

T = int(a.readline())
t = T


while t:
    N, P = map(int, a.readline().split())
    CP = [int(_) for _ in a.readline().split()]

    closest_P = P*N

    j = jj = 0
    while j < N:
        for k in range(jj+1, N+1):
            v = sum(CP[j:k])
            if v >= P:
                if closest_P > v:
                    closest_P = v
                    l, r = j, k-1
                break
            else:
                jj = k
        else:
            break
        j += 1

    print('Case #%d: %d %d' % (T-t+1, l, r))
    t -= 1
