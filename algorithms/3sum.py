class Solution:
    def threeSum(self, nums: [int]) -> [[int]]:

        nums.sort()
        maxl = len(nums)
        ans = []

        for z in range(maxl-2):

            if z > 0 and nums[z] == nums[z-1]:
                continue

            l, r = z+1, maxl-1
            while l < r:
                tsum = nums[l] + nums[r] + nums[z]
                if tsum < 0:
                    l += 1
                elif tsum > 0:
                    r -= 1
                else:
                    ans.append([nums[l], nums[r], nums[z]])
                    while l < r and nums[l] == nums[l+1]:
                        l += 1
                    while l < r and nums[r] == nums[r-1]:
                        r -= 1
                    l += 1
                    r -= 1

        return ans
