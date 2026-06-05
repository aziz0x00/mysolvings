# ugly
class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        if not intervals:
            return [newInterval]

        added = False
        ret = []
        for s, e in intervals:
            if e < newInterval[0] or (s > newInterval[1] and added):
                ret.append([s, e])
            else:
                added = True
                if len(ret) and ret[-1][1] >= newInterval[0]:
                    ret[-1][1] = max(newInterval[1], e)
                else:
                    if s > newInterval[1]:
                        ret.append(newInterval)
                        ret.append([s, e])
                    else:
                        new_i, new_j = min(newInterval[0], s), max(newInterval[1], e)
                        ret.append([new_i, new_j])

        if not added:
            ret.append(newInterval)

        return ret
