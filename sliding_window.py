class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        n = len(s)
        maxLength = 0
        charIndex = [-1] * 128
        left = 0
        
        for right in range(n):
            left = self.new_method(s, charIndex, left, right)
            maxLength = max(maxLength, right - left + 1)
        
        return maxLength
<<<<<<< HEAD
        
        
# Example usage:
=======

    def new_method(self, s, charIndex, left, right):
        if charIndex[ord(s[right])] >= left:
            left = charIndex[ord(s[right])] + 1
        charIndex[ord(s[right])] = right
        return left
# Example  usage:
>>>>>>> b674656 (Add example usage commen)
solution = Solution()
print(solution.lengthOfLongestSubstring("abcabcbb"))  # Output: 3