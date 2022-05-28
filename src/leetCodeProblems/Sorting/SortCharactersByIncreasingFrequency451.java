package leetCodeProblems.Sorting;

/**
 * LeetCode - https://leetcode.com/problems/sort-characters-by-frequency/
 * TimeComplexity - O(nlogn)
 * Space Complexity - O(n)
 */

import java.util.*;

public class SortCharactersByIncreasingFrequency451 {

    static HashMap<Character, Integer> map;
    static class ArrayFrequencySort implements Comparator<Character>{

        public int compare(Character a, Character b) {

            int aFrequency = map.get(a);
            int bFrequency = map.get(b);

            if (aFrequency == bFrequency) {
                return b -a;
            }
            return bFrequency - aFrequency;
        }

    }
    public String frequencySort(String s) {

        ArrayList<Character> inputList = new ArrayList<>();
        map = new HashMap<>();

        String outputString = "";

        for (int i=0; i < s.length(); i++) {

            if (map.containsKey(s.charAt(i))) {
                map.put(s.charAt(i), map.get(s.charAt(i)) + 1);
            }
            else {
                map.put(s.charAt(i), 1);
            }

            inputList.add(s.charAt(i));

        }

        System.out.println(map);

        Collections.sort(inputList, new ArrayFrequencySort());

        //System.out.println(inputList);

        for (int i=0; i < s.length(); i++) {
            outputString += inputList.get(i);
        }

        return outputString;
    }

    public static void main(String[] args) {

        //String input = "tree";

        String input = "loveleetcode";

        SortCharactersByIncreasingFrequency451 obj = new SortCharactersByIncreasingFrequency451();

        System.out.println(obj.frequencySort(input));

    }
}
