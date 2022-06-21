package leetCodeProblems.HashSearch;

/**
 * LeetCode - https://leetcode.com/problems/first-unique-character-in-a-string/submissions/
 *
 * TimeComplexity - O(n)
 * SpaceComplexity - O(n)
 */

import java.util.HashMap;

public class FirstUniqueCharacterInAString387 {

    public int firstUniqChar(String s) {

        HashMap<Character, Integer> map = new HashMap<>();

        for(int i=0; i < s.length(); i++) {

            if (map.containsKey(s.charAt(i))) {
                map.put(s.charAt(i), map.get(s.charAt(i))+1);
            }
            else {
                map.put(s.charAt(i), 1);
            }
        }

        for(int i=0; i < s.length(); i++) {

            if (map.get(s.charAt(i)) == 1) {
                return i;
            }
        }

        return -1;
    }

    public static void main(String[] args) {

        //String s = "loveleetcode";

        String s = "leetcode";

        FirstUniqueCharacterInAString387 obj = new FirstUniqueCharacterInAString387();

        System.out.println(obj.firstUniqChar(s));
    }
}
