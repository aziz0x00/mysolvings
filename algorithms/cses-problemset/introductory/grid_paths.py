g = [min(i % 9, i // 9) == 0 or max(i % 9, i // 9) == 8 for i in range(9 * 9)]


def search(p, n):
    if p == 9 * 7 + 1:
        return int(n == 48)

    r = 0
    g[p] = True
    for m in flow[n]:
        q = p + m

        if g[q]:
            continue

        a = (g[q - 1] | g[q + 1]) < (g[q + 9] & g[q - 9])
        b = (g[q - 1] & g[q + 1]) > (g[q + 9] | g[q - 9])

        if a ^ b:
            continue

        r += search(q, n + 1)

    g[p] = False

    return r

dz = (-9, 9, -1, 1)
flow = [dz if i == "?" else dz["UDLR".index(i)] for i in input()]

print(search(1 + 1 * 9, 0))
