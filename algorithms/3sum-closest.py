class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:

        # Closest is a tuple of (sum, difference_to_target)
        closest = (None, None)
        nums.sort()
        maxl = len(nums)

        for z in range(maxl-2):

            if z > 0 and nums[z] == nums[z-1]:
                continue

            l, r = z+1, maxl-1

            while l < r:

                tsum = nums[l] + nums[r] + nums[z]

                if closest[0] == None or abs(target-tsum) < closest[1]:
                    closest = (tsum, abs(target-tsum))

                if tsum < target:
                    l += 1
                elif tsum > target:
                    r -= 1
                else:
                    return tsum

        return closest[0]
