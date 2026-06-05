"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
"""

from typing import Optional
class Solution:
    def cloneGraph(self, node: Optional['Node']) -> Optional['Node']:
        visited = {}

        if node == None: return None

        def clone(n):
            if not visited.get(n):
                visited[n] = Node(val=n.val)
                visited[n].neighbors = [clone(m) for m in n.neighbors if n]
            return visited[n]

        return clone(node)
