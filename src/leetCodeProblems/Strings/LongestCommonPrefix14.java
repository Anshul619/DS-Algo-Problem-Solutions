package leetCodeProblems.Strings;

/**
 * LeetCode - https://leetcode.com/problems/longest-common-prefix/
 */

public class LongestCommonPrefix14 {

    public String longestCommonPrefixUtil(String str1, String str2) {

        String prefix = "";

        for (int i=0, j=0; i < str1.length() && j < str2.length(); i++, j++) {

            if (str1.charAt(i) == str2.charAt(j)) {
                prefix += str1.charAt(i);
            }
            else {
                break;
            }
        }

        return prefix;

    }

    public String longestCommonPrefix(String[] strs) {

        String commonPrefix = "";

        if (strs.length > 1) {
            commonPrefix = longestCommonPrefixUtil(strs[0], strs[1]);
        }

        System.out.println(commonPrefix);

        for (int i=2; i < strs.length; i++) {

            commonPrefix = longestCommonPrefixUtil(commonPrefix, strs[i]);

            if (commonPrefix.isEmpty()) {
                return "";
            }
        }

        return commonPrefix;
    }

    public static void main(String[] args) {

        //String[] strs = {"flower","flow","flight"};

        String[] strs = {"cir","car"};

        LongestCommonPrefix14 obj = new LongestCommonPrefix14();
        System.out.println(obj.longestCommonPrefix(strs));

    }
}
