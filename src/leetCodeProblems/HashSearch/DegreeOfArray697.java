package leetCodeProblems.HashSearch;

/**
 * LeetCode - https://leetcode.com/problems/degree-of-an-array/
 */

import java.util.HashMap;

public class DegreeOfArray697 {

    static class DegreeWithLength {

        int degree;
        int startLength;
        int endLength;

        DegreeWithLength(int degree, int startLength, int endLength) {
            this.degree = degree;
            this.startLength = startLength;
            this.endLength = endLength;
        }
    }


    public int findShortestSubArray(int[] nums) {

        HashMap<Integer, DegreeWithLength> degreeMap = new HashMap<>();

        int maxElement = 0;
        int maxDegree = 0;
        int minLength = Integer.MAX_VALUE;

        for (int i =0; i < nums.length; i++) {

            if (degreeMap.containsKey(nums[i])) {

                DegreeWithLength obj = degreeMap.get(nums[i]);

                obj.degree++;
                obj.endLength = i;

                degreeMap.put(nums[i], obj);
            }

            else {

                DegreeWithLength obj = new DegreeWithLength(1, i, i);

                degreeMap.put(nums[i], obj);
            }
        }
        
        for (Integer element: degreeMap.keySet()) {

            DegreeWithLength obj = degreeMap.get(element);

            if (obj.degree >= maxDegree) {

                if (obj.degree > maxDegree) {
                    minLength = Integer.MAX_VALUE;
                    maxDegree = obj.degree;
                }

                maxElement = element;
                System.out.println("----start");
                System.out.println("maxElement -> "+ maxElement);
                System.out.println("maxDegree -> "+ maxDegree);
                System.out.println("minLength before -> "+ minLength);
                minLength = Math.min(minLength, (obj.endLength - obj.startLength));
                System.out.println("minLength after -> "+ minLength);
                System.out.println("----end");
            }
        }

        //System.out.println(maxDegree);
        minLength++;

        return minLength;
    }

    public static void main(String[] args) {

        DegreeOfArray697 obj = new DegreeOfArray697();

        int[] arr = {1,2,2,3,1}; // expected output = 2

        //int[] arr = {1,2,2,3,1,4,2}; // expected output = 6

        //int[] arr = {1,3,2,2,3,1}; // expected output = 2

        System.out.println(obj.findShortestSubArray(arr));

    }
}
