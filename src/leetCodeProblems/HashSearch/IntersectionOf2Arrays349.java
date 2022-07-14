package leetCodeProblems.HashSearch;

/**
 * LeetCode - https://leetcode.com/problems/intersection-of-two-arrays/solution/
 *
 * TimeComplexity - O(m+n)
 * SpaceComplexity - O(m+n)
 */

import java.util.Arrays;
import java.util.HashSet;

public class IntersectionOf2Arrays349 {

    HashSet<Integer> uniqueInputSet;

    HashSet<Integer> outputSet;

    private void addElementsInUniqueSet(int[] array) {

        for(int num: array) {
            if (!uniqueInputSet.contains(num)) {
                uniqueInputSet.add(num);
            }
        }
    }

    private void addElementsInOutputSet(int[] array) {

        for(int num: array) {
            if (uniqueInputSet.contains(num) && !outputSet.contains(num)) {
                outputSet.add(num);
            }
        }
    }

    public int[] intersection(int[] nums1, int[] nums2) {

        uniqueInputSet = new HashSet<>();
        outputSet = new HashSet<>();

        int nums1Length = nums1.length;
        int nums2Length = nums2.length;

        if (nums1Length < nums2Length) {
            addElementsInUniqueSet(nums2);
            addElementsInOutputSet(nums1);
        }
        else {
            addElementsInUniqueSet(nums1);
            addElementsInOutputSet(nums2);
        }

        int[] outputArray = new int[outputSet.size()];
        int index = 0;

        for(int num: outputSet) {
            outputArray[index] = num;
            index++;
        }

        return outputArray;
    }

    public static void main(String[] args) {
        //int[] nums1 = {1,2,2,1};
        //int[] nums2 = {2,2};

        int[] nums1 = {4,9,5};
        int[] nums2 = {9,4,9,8,4};

        IntersectionOf2Arrays349 obj = new IntersectionOf2Arrays349();

        int[] output = obj.intersection(nums1, nums2);
        System.out.println(Arrays.toString(output));

    }
}
