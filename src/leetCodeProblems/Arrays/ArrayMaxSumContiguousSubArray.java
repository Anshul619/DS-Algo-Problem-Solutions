package leetCodeProblems.Arrays;

public class ArrayMaxSumContiguousSubArray {

	// DO NOT MODIFY THE ARGUMENTS WITH "final" PREFIX. IT IS READ ONLY
    public int maxSubArray(final int[] A) {

        int max_so_far = Integer.MIN_VALUE;
        int max_ending_here = 0;

        for (int element : A) {

            max_ending_here = max_ending_here + element;

            if (max_so_far < max_ending_here) {
                max_so_far = max_ending_here;
            }

            if (max_ending_here < 0) {
                max_ending_here = 0;
            }
        }

        return max_so_far;
    }

}
