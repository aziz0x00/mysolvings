class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        import queue

        pq = queue.PriorityQueue()

        for i, p in enumerate(points):
            dis = p[0] ** 2 + p[1] ** 2
            pq.put((dis, i))

        ans = []
        while k:
            ans.append(points[pq.get()[1]])
            k -= 1

        return ans
