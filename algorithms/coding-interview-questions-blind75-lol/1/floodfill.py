class Solution:
    def floodFill(self, image: List[List[int]], sr: int, sc: int, color: int) -> List[List[int]]:
        
        n = len(image)
        m = len(image[0])
        init_color = image[sr][sc]


        def dfs(i, j):
            image[i][j] = color
            if i + 1 < n and image[i+1][j] == init_color:
                dfs(i+1, j)
            if j + 1 < m and image[i][j+1] == init_color:
                dfs(i, j+1)
            if i > 0 and image[i-1][j] == init_color:
                dfs(i-1, j)
            if j > 0 and image[i][j-1] == init_color:
                dfs(i, j-1)
        
        if image[sr][sc] != color:
            dfs(sr, sc)

        return image
