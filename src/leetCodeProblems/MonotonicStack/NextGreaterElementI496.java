package leetCodeProblems.MonotonicStack;

/**
 * LeetCode - https://leetcode.com/problems/next-greater-element-i/solution/
 */

import java.util.Arrays;
import java.util.HashMap;
import java.util.Stack;

public class NextGreaterElementI496 {

    public int[] nextGreaterElement(int[] nums1, int[] nums2) {

        Stack<Integer> nums2Stack = new Stack<>();

        HashMap<Integer, Integer> nums2NextGreaterMap = new HashMap<>();

        int[] output = new int[nums1.length];

        for (int i=0; i < nums2.length; i++) {

            while (!nums2Stack.isEmpty() && nums2[i] > nums2Stack.peek()) {
                nums2NextGreaterMap.put(nums2Stack.pop(), nums2[i]);
            }

            nums2Stack.push(nums2[i]);
        }

        System.out.println(nums2NextGreaterMap);
        while(!nums2Stack.isEmpty()) {
            nums2NextGreaterMap.put(nums2Stack.pop(), -1);
        }

        for (int i=0; i < nums1.length; i++) {
            output[i] = nums2NextGreaterMap.get(nums1[i]);
        }

        return output;
    }

    public static void main(String[] args) {

        //int[] nums1 = {4,1,2};
        //int[] nums2 = {1,3,4,2};

        //int[] nums1 = {2, 4};
        //int[] nums2 = {1,2,3,4};

        int[] nums1 = {1,3,5,2,4};
        int[] nums2 = {6,5,4,3,2,1,7};

        NextGreaterElementI496 obj = new NextGreaterElementI496();

        int[] output = obj.nextGreaterElement(nums1, nums2);

        System.out.println(Arrays.toString(output));
    }
}
