package leetCodeProblems.SlidingWindow;

/**
 * LeetCode - https://leetcode.com/problems/longest-substring-without-repeating-characters/
 */

import java.util.HashMap;

public class LongestSubstringWithoutRepeat3 {

    HashMap<Character, Integer> uniqueCharactersMap = new HashMap<>();

    public void removeLowerIndexes(int maxIndex) {

        for (Character characterVal: uniqueCharactersMap.keySet()) {

            if (uniqueCharactersMap.get(characterVal) < maxIndex) {

                //System.out.println("Remove characterVal -> " + characterVal);
                //System.out.println("Before -> uniqueCharactersMap =>" + uniqueCharactersMap);
                uniqueCharactersMap.put(characterVal, -1);
                //System.out.println("After -> uniqueCharactersMap =>" + uniqueCharactersMap);
            }
        }

        return;
    }

    public int lengthOfLongestSubstring(String s) {

        int localLeftPointer = 0;
        int localRightPointer = 0;

        int ans = 0;

        if (s.length() > 0) {
            ans = 1;
        }

        uniqueCharactersMap = new HashMap<>();


        while (localRightPointer < (s.length() - 1)) {
            uniqueCharactersMap.put(s.charAt(localRightPointer), localRightPointer);

            localRightPointer++;

            char currentCharacter = s.charAt(localRightPointer);
            System.out.println("localRightPointer ->" + localRightPointer);

            // Reset counters
            if (uniqueCharactersMap.containsKey(currentCharacter) && uniqueCharactersMap.get(currentCharacter) != -1) {
                localLeftPointer = uniqueCharactersMap.get(s.charAt(localRightPointer)) + 1;
                removeLowerIndexes(localLeftPointer);
                System.out.println("CHANGE - localLeftPointer ->" + localLeftPointer);
            }

            ans = Math.max(ans, (localRightPointer-localLeftPointer+1));

            System.out.println("ans ->" + ans);
        }

        return ans;
    }

    public static void main(String[] args) {

        LongestSubstringWithoutRepeat3 obj = new LongestSubstringWithoutRepeat3();

        //String inputString = "bccbababd";

        String inputString = "abcabcbb"; // Ouput = 3

        //String inputString = "bbbbb"; // Ouput = 1

        //String inputString = "aa"; // Output = 1

        //String inputString = "a"; // Output = 1

        //String inputString = "ab"; // Output = 2

        //String inputString = "pwwkew"; // Output = 3

        //String inputString = "abba"; // Output = 2

        int longestLength = obj.lengthOfLongestSubstring(inputString);

        System.out.println("longestLength ->" + longestLength);
    }
}
