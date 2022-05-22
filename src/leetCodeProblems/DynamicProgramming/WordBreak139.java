package leetCodeProblems.DynamicProgramming;

import java.util.ArrayList;
import java.util.List;

/**
 * LeetCode - https://leetcode.com/problems/word-break/
 */

public class WordBreak139 {
    public boolean wordBreak(String s, List<String> wordDict) {

        // DP
        boolean[] alreadySavedResults = new boolean[s.length()+1];

        alreadySavedResults[0] = true;

        for (int lengthToCheck = 1; lengthToCheck <= s.length(); lengthToCheck ++) {

            for (int i = 0; i < lengthToCheck; i++) {

                // alreadySavedResults[i] - Check whether rest of the string is breaking or not
                if (alreadySavedResults[i] && wordDict.contains(s.substring(i, lengthToCheck))) {

                    alreadySavedResults[lengthToCheck] = true;
                    break;
                }

            }
        }

        return alreadySavedResults[s.length()];
    }

    public static void main(String[] args) {

        WordBreak139 obj = new WordBreak139();

        List<String> wordDict = new ArrayList<>();
        wordDict.add("i");
        wordDict.add("like");
        wordDict.add("sam");
        wordDict.add("sun");
        wordDict.add("samsung");

        System.out.println(obj.wordBreak("ilikesamsung", wordDict));

    }
}
