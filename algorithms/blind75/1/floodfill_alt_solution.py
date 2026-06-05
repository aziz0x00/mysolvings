# not mine, sadly lost ref.
class Solution:
    def floodFill(
        self, image: List[List[int]], sr: int, sc: int, color: int
    ) -> List[List[int]]:
        m = len(image)
        n = len(image[0])
        visited = set()
        directions = [[0, 1], [1, 0], [0, -1], [-1, 0]]
        cell_value = image[sr][sc]

        def bfs(r, c):
            q = collections.deque()
            q.append((r, c))
            visited.add((r, c))
            image[r][c] = color

            while q:
                r, c = q.popleft()
                for dr, dc in directions:
                    r1, c1 = r + dr, c + dc
                    if (
                        0 <= r1 < m
                        and 0 <= c1 < n
                        and image[r1][c1] == cell_value
                        and (r1, c1) not in visited
                    ):
                        q.append((r1, c1))
                        visited.add((r1, c1))
                        image[r1][c1] = color

        bfs(sr, sc)
        print(image)
        return image
