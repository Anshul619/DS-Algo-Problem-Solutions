package leetCodeProblems.Sorting;

/**
 * LeetCode - https://leetcode.com/problems/maximum-units-on-a-truck/solution/
 *
 * TimeComplexity - O(nlogn)
 * SpaceComplexity - O(1)
 */

import java.util.Arrays;
import java.util.Comparator;

public class MaxUnitsOnTruck1710 {

    static class UnitsComparator implements Comparator<int[]> {

        public int compare(int[] a, int[] b) {
            return b[1]-a[1];
        }
    }
    public int maximumUnits(int[][] boxTypes, int truckSize) {

        Arrays.sort(boxTypes, new UnitsComparator());

        //System.out.println(Arrays.toString(boxTypes));

        int remainingSizeCounter = truckSize;
        int answer = 0;

        for (int i=0; i < boxTypes.length; i++) {

            //System.out.println(Arrays.toString(boxTypes[i]));

            if (boxTypes[i][0] <= remainingSizeCounter) {
                answer += boxTypes[i][0]*boxTypes[i][1];
                remainingSizeCounter -= boxTypes[i][0];
            }
            else {
                answer += remainingSizeCounter*boxTypes[i][1];
                remainingSizeCounter = 0;
            }
        }

        return answer;
    }

    public static void main(String[] args) {

        int[][] boxes = {{1, 3}, {2, 2}, {3,1}};
        int truckSize = 4;

        MaxUnitsOnTruck1710 obj = new MaxUnitsOnTruck1710();
        System.out.println(obj.maximumUnits(boxes, truckSize));
    }
}
