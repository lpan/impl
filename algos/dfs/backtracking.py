"""
backtracking!
"""

def subset(*nums):
    """
    https://leetcode.com/problems/subsets/description/
    pre: nums has no duplicates
    """

    result = [[]]
    tmp = []

    def backtrack(start = 0):
        if start == len(nums):
            return

        for i in range(start, len(nums)):
            tmp.append(nums[i])
            result.append(tmp[:])

            backtrack(i + 1)

            tmp.pop()

    backtrack()
    return result

def permutation_sum(nums, target):
    """
    modified version of https://leetcode.com/problems/combination-sum/description/
    pre: nums has no duplicates and is sorted
    """

    result = []
    tmp = []

    def backtrack():
        if sum(tmp) == target:
            result.append(tmp[:])

        if sum(tmp) >= target:
            return

        for n in nums:
            tmp.append(n)
            backtrack()
            tmp.pop()

    backtrack()
    return result

def combination_sum(nums, target):
    """
    https://leetcode.com/problems/combination-sum/description/
    pre: nums has no duplicates and is sorted
    """

    result = []
    tmp = []

    def backtrack(start = 0):
        if sum(tmp) == target:
            result.append(tmp[:])

        if sum(tmp) >= target:
            return

        for i in range(start, len(nums)):
            tmp.append(nums[i])
            backtrack(i) # since nums is sorted, we only get one permutation of the subset
            tmp.pop()

    backtrack()
    return result

def palindrome_partition(input_str):
    """
    https://leetcode.com/problems/palindrome-partitioning/description/
    """

    def is_palindrome(s):
        low, high = 0, len(s) - 1
        while (low < high):
            if s[low] != s[high]:
                return False

            low += 1
            high -= 1

        return True

    result = []
    tmp = []

    def backtrack(start = 0):
        if start == len(input_str):
            result.append(tmp[:])

        # try all possible sub strings from start
        for i in range(start, len(input_str)):
            end = i + 1
            s = input_str[start:end]

            if is_palindrome(s):
                tmp.append(s)
                backtrack(end)
                tmp.pop()

    backtrack()
    return result


if __name__ == '__main__':
    print(subset(0, 1, 2, 3))
    print(permutation_sum([2, 3, 6, 7], 7))
    print(combination_sum([2, 3, 6, 7], 7))
    print(palindrome_partition("aab"))
