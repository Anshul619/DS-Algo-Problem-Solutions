package leetCodeProblems.HashSearch;

import java.util.Collections;
import java.util.HashMap;

public class LongestSubstringWithKUniqueCharacters340 {

    public int lengthOfLongestSubstringKDistinct(String s, int k) {

        int ansLongestLength = 0;

        int localLeftPointer = 0;
        int localRightPointer = 0;

        HashMap<Character, Integer> uniqueCharactersMap = new HashMap<>();

        if (s.length() == 1 && k > 0) {
            ansLongestLength = 1;
        }

        while(localRightPointer < s.length()) {

            uniqueCharactersMap.put(s.charAt(localRightPointer), localRightPointer);
            localRightPointer++;

            //System.out.println(uniqueCharactersMap);

            if (uniqueCharactersMap.size() > k) {
                int minUniqueCharacterIndex = Collections.min(uniqueCharactersMap.values());
                uniqueCharactersMap.remove(s.charAt(minUniqueCharacterIndex));

                localLeftPointer = minUniqueCharacterIndex+1;
            }

            if (uniqueCharactersMap.size() <= k) {
                ansLongestLength = Math.max(ansLongestLength, localRightPointer-localLeftPointer);
            }
        }

        return ansLongestLength;

    }

    public static void main(String[] args) {

        LongestSubstringWithKUniqueCharacters340 obj = new LongestSubstringWithKUniqueCharacters340();

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
