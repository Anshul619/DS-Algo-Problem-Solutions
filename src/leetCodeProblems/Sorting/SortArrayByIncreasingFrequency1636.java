package leetCodeProblems.Sorting;

/**
 * LeetCode - https://leetcode.com/problems/sort-array-by-increasing-frequency/submissions/
 * TimeComplexity - O(n)
 * Space Complexity - O(n)
 */

import java.util.*;

public class SortArrayByIncreasingFrequency1636 {

    static HashMap<Integer, Integer> map;
    static class ArrayFrequencySort implements Comparator<Integer>{

        public int compare(Integer a, Integer b) {

            int aFrequency = map.get(a);
            int bFrequency = map.get(b);

            if (aFrequency == bFrequency) {
                return b - a;
            }
            else {
                return aFrequency - bFrequency;
            }
        }

    }
    public int[] frequencySort(int[] nums) {

        ArrayList<Integer> inputList = new ArrayList<>();
        map = new HashMap<>();

        for (int i=0; i < nums.length; i++) {

            if (map.containsKey(nums[i])) {
                map.put(nums[i], map.get(nums[i]) + 1);
            }
            else {
                map.put(nums[i], 1);
            }

            inputList.add(nums[i]);

        }

        Collections.sort(inputList, new ArrayFrequencySort());

        //System.out.println(inputList);

        for (int i=0; i < nums.length; i++) {
            nums[i] = inputList.get(i);
        }

        return nums;
    }

    public static void main(String[] args) {

        int[] input = {1,1,2,2,2,3};

        SortArrayByIncreasingFrequency1636 obj = new SortArrayByIncreasingFrequency1636();

        System.out.println(Arrays.toString(obj.frequencySort(input)));

    }
}
