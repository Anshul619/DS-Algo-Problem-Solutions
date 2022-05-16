package leetCodeProblems.DynamicProgramming;

/**
 * 
 * LeetCode - https://leetcode.com/problems/longest-palindromic-substring/
 * 
 * @author anshul.agrawal
 *
 */
public class LongestPalindromicSubString5 {
	public String longestPalindrome(String s) {
        
        int length = s.length();
        int ansSubStringLength = 0;
        int ansSubStringStartIndex = 0;
        
        boolean[][] table = new boolean[length][length];
        
        // All substrings of length 1, are PALINDROME
        ansSubStringLength = 1;
        for(int i=0; i < length; i++) {
            table[i][i] = true;
        }
        
        // Check for substrings with length 2.
        // This can't be handled in (>3) loop due to logic condition mismatch
        for(int i=0; i< length-1; i++) {
            if (s.charAt(i) == s.charAt(i+1)) {
                ansSubStringStartIndex = i;
                ansSubStringLength = 2;
                table[i][i+1] = true;
            }
        }
        
        // Check for substrings with length greater than 3
        for(int subStringLengthToCheck = 3; subStringLengthToCheck <= length; subStringLengthToCheck++) {
            
            for(int startIndex=0; startIndex < length-subStringLengthToCheck+1; startIndex++) {
                
                int endIndex = startIndex+subStringLengthToCheck-1;
                
                if (table[startIndex+1][endIndex-1]) { // sub-sub string is already PALINDROME
                    if (s.charAt(startIndex) == s.charAt(endIndex)) {
                        table[startIndex][endIndex] = true;
                        
                        // Set the answer variables
                        if (subStringLengthToCheck > ansSubStringLength) {
                            ansSubStringStartIndex = startIndex;
                            ansSubStringLength = subStringLengthToCheck;
                        }
                    }
                }
                
            }
        }
        
        String ansSubString = s.substring(ansSubStringStartIndex, ansSubStringStartIndex + ansSubStringLength);
        
        //System.out.println(ansSubString);
        
        return ansSubString;
        
    }

}
