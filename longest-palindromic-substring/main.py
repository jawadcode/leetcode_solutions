class Solution:
  # Take a given centre of a string using the start ("l") and the end ("r") and then expand outwards until the outermost characters are not equal and then return the length (difference between start and end)
  def expand(self, s, l, r):
    while l >= 0 and r < len(s) and s[l] == s[r]:
      l -= 1
      r += 1
    
    return r - l - 1


  def longestPalindrome(self, s: str) -> str:
    start = 0
    end = 0
    rev = s[::-1] # String "s" in reverse
    if s == None or len(s) < 1:
      return ""
    # Check if "s" is equal to its reverse to avoid unnecessary looping 
    elif s == rev:
      return rev
    # Loop through characters in s
    for i in range(len(s)):
      # Get the maximum length expanding from centre "i"
      len1 = self.expand(s, i, i)
      # Get the maximum length expanding from centre "i + 1/2" (between "i" and "i + 1")
      len2 = self.expand(s, i, i + 1)
      # Compare the two to get the max length between centres "i" and "i + 1/2"
      length = max(len1, len2)
      # If the length is greater than the previous max length (difference between old end and old start) then assign the start: the current position in string minus half of the length, and the end: the current position add half of the length
      if length > end - start:
        start = i - (length - 1) // 2
        end = i + length // 2

    return s[start:end + 1]

# Demo code
sol = Solution()
while True:
  test = input("Please input a string to check for the longest palindromic substring: ")
  print(test)
  print("Result:", sol.longestPalindrome(test))