package leetCodeProblems.SlidingWindow;

/**
 * LeetCode link - https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/
 */

import java.util.HashSet;

public class LongestSubstringWithKUniqueCharacters340WithoutSlidingWindow {

    public int lengthOfLongestSubstringKDistinct(String s, int k) {

        int ansLongestLength = 0;
        int ansStartLength = 0;

        HashSet<Character> charactersHashSet = new HashSet<>();
        int localCountUniqueCharacters = 0;

        if (s.length() == 1 && k > 0) {
            ansLongestLength = 1;
        }

        for(int i=0; i<s.length(); i++) {

            localCountUniqueCharacters = 1;
            charactersHashSet.clear();
            charactersHashSet.add(s.charAt(i));

            int j;

            for(j=i+1; j<s.length();j++) {

                if (charactersHashSet.contains(s.charAt(j))) {
                    // Don't increase localCountUniqueCharacters
                }
                else {

                    if (localCountUniqueCharacters+1 > k) {
                        break;
                    }

                    localCountUniqueCharacters++;
                    charactersHashSet.add(s.charAt(j));
                }
            }

            if (localCountUniqueCharacters <= k) {
                //System.out.println("Increase count");
                ansLongestLength = Math.max(ansLongestLength, j-i);
            }
        }

        return ansLongestLength;

    }

    public static void main(String[] args) {

        LongestSubstringWithKUniqueCharacters340WithoutSlidingWindow obj = new LongestSubstringWithKUniqueCharacters340WithoutSlidingWindow();

        String inputString = "bccbababd";
        int k = 2; // answer = 5, Substring = babab

        //String inputString = "aa";
        //int k = 1; // answer = 2

        //String inputString = "a";
        //int k = 1; // answer = 1

        //String inputString = "aa";
        //int k = 2; // answer = 1

        //String inputString = "ab";
        //int k = 1; // answer = 1

        int longestLength = obj.lengthOfLongestSubstringKDistinct(inputString, k);

        System.out.println("longestLength ->" + longestLength);
    }

}
