def rot(p: int, k: int, n: int) -> int:
    return (k + p) % n


def solve(n: int, a: int, b: int):

    if (a == 0) != (b == 0) or a + b > n:
        print("NO")
        return
    print("YES")

    for i in range(n):
        print(i + 1, end=" ")
    print()

    fixed = n - a - b
    for i in range(fixed):
        print(1 + i, end=" ")
    for i in range(n - fixed):
        print(1 + fixed + rot(a, i, a + b), end=" ")
    print()


t = int(input())

for _ in range(t):
    n, a, b = map(int, input().split())

    solve(n, a, b)
