# binary exponantiation of the Fibonacci matrix
class Solution:
    def climbStairs(self, n: int) -> int:
        F = [[1, 1], [1, 0]] # F^n = [ [F_(n+1), F_n], [F_n, F_(n-1)] ]
        return self.mat_pow(F, n)[0][0]
    
    def mat_mult(self, A, B):
        C = [[0, 0], [0, 0]]
        for i in range(len(A)):
            for j in range(len(B[0])):
                for k in range(len(A[0])):
                    C[i][j] += A[i][k] * B[k][j]
        print(C)
        return C
    
    def mat_pow(self, A, n):
        B = [[1, 0], [0, 1]]
        k = n.bit_length() - 1
        while k >= 0:
            B = self.mat_mult(B, B)
            if n & (1 << k):
                B = self.mat_mult(B, A)
            k -= 1
        return B

"""
A^n, n = 2^k + ... + 2^m = 10001, 101011, ..
"""
