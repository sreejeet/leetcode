'''
Brute-force solution
'''


class Solution:
    def largestTimeFromDigits(self, arr: [int]) -> str:

        if min(arr) > 2:
            return ''

        arr.sort()
        for H in range(23, -1, -1):
            h = str(H).zfill(2)
            if not all([int(x) in arr for x in h]):
                continue
            for M in range(59, -1, -1):
                m = str(M).zfill(2)
                if arr == sorted([int(i) for i in h+m]):
                    return h+':'+m

        return ''
