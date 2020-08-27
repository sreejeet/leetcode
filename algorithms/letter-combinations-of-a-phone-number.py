class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        keypad = {
            2: "abc",
            3: "def",
            4: "ghi",
            5: "jkl",
            6: "mno",
            7: "pqrs",
            8: "tuv",
            9: "wxyz",
        }

        ans_list = []
        max_l = len(digits)

        alphas = [keypad[int(x)] for x in digits]

        def __recur__(ans="", chars_i=0):
            chars = alphas[chars_i] if chars_i < max_l else []
            nex = alphas[chars_i+1] if chars_i+1 < max_l else []
            if nex == []:
                for c in chars:
                    ans_list.append(ans+c)
            else:
                for c in chars:
                    r(ans+c, chars_i+1)

        __recur__()
        return ans_list