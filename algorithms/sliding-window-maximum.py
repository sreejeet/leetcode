class Solution:

    def maxSlidingWindow(self, nums: [int], k: int) -> [int]:

        l, r = 0, k
        max_r = len(nums)
        ans = [max(nums[l:r])]

        while r < max_r:
            
            r+=1
            l+=1

            # If the new number is larger than the last largest number,
            # we have the new largest number.
            if nums[r-1] > ans[-1]:
                ans.append(nums[r-1])

            # If we have lost the largest number,
            # find new largest number.
            elif nums[l-1] == ans[-1]:
                ans.append(max(nums[l:r]))

            # If neither cases match, repeat the last max.
            else:
                ans.append(ans[-1])

        return ans


if __name__=='__main__':

    sol = Solution()
    print(sol.maxSlidingWindow([1,-1], 1))

"""
[1,3,-1,-3,5,3,6,7]
3
[1,3,-1,-3,5,3,6,7]
1
[1,3,-1,-3,5,3,6,7]
8
[1,3,-1,-3,5,3,6,7]
7
[1,-1]
1
"""