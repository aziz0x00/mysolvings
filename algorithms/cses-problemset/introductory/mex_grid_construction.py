n = int(input())

mex_mat = [[0 for _ in range(n)] for _ in range(n)]


def mex(lst):
    return min(set(range(len(lst) + 1)) - set(lst))


# at worst case O(n^3 log n)
for i in range(n):
    for j in range(n):
        selected = [mex_mat[i][k] for k in range(j)] + [mex_mat[k][j] for k in range(i)]
        mex_mat[i][j] = mex(selected)

for i in range(n):
    for j in range(n):
        print(mex_mat[i][j], end=" ")
    print()
