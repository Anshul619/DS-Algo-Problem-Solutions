package leetCodeProblems.GreedyAlgo;

/**
 * LeetCode - https://leetcode.com/problems/find-pivot-index/
 * Similar Code - https://leetcode.com/problems/find-the-middle-index-in-array/
 * Time Complexity - O(n)
 * Space Complexity - O(1)
 * 
 */
public class FindPivotIndex724 {

    private int calSum(int[] nums, int start, int end) {

        int sum = 0;

        for (int i=start; i <= end; i++) {
            sum += nums[i];
        }

        return sum;
    }
    public int pivotIndex(int[] nums) {

        int length = nums.length;

        if (length == 0){
            return -1;
        }

        if (length == 1){
            return 0;
        }

        int leftSum = 0;
        int rightSum = 0;

        int totalSum = calSum(nums, 0, length-1);

        for (int i=0; i < nums.length; i++) {

            if (i > 0) {
                leftSum += nums[i-1];
            }

            rightSum = totalSum - leftSum - nums[i];

            if (leftSum == rightSum) {
                return i;
            }

        }

        return -1;

    }

    public static void main(String[] args) {

        //int[] inputArray = {1, 2, 2, 1};
        //int targetDifference = 1; // Output = 4

        int[] inputArray = {1,7,3,6,5,6};

        //int[] inputArray = {2,1,-1};

        //int[] inputArray = {1,2,3};

        //int[] inputArray = {1, 3, 6};
        //int targetDifference = 2; // Output = 1

        FindPivotIndex724 obj = new FindPivotIndex724();

        System.out.println(obj.pivotIndex(inputArray));
    }
}
