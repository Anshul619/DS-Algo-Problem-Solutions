package leetCodeProblems.HashSearch;

/**
 * LeetCode - https://leetcode.com/problems/unique-number-of-occurrences/
 */

import java.util.HashMap;
import java.util.HashSet;

public class UniqueNumberOfOccurrences1207 {

    public boolean uniqueOccurrences(int[] arr) {

        HashMap<Integer, Integer> occurrencesMap = new HashMap<>();
        HashSet<Integer> occurrences = new HashSet<>();

        for (int i=0; i < arr.length; i++) {

            if (occurrencesMap.containsKey(arr[i])) {
                int currentCount = occurrencesMap.get(arr[i]);
                currentCount++;
                occurrencesMap.put(arr[i], currentCount);
            }
            else {
                occurrencesMap.put(arr[i], 1);
            }

        }

        for (int frequency: occurrencesMap.values()) {

            if (occurrences.contains(frequency)) {
                return false;
            }

            occurrences.add(frequency);

        }

        return true;

    }

    public static void main(String[] args) {

        //int[] arr = {1,2,2,1,1,3};

        int[] arr = {1,2};

        UniqueNumberOfOccurrences1207 obj = new UniqueNumberOfOccurrences1207();

        System.out.println(obj.uniqueOccurrences(arr));
    }
}
