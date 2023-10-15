package leetCodeProblems.BruteForce;

/**
 * LeetCode - https://leetcode.com/problems/sign-of-the-product-of-an-array/submissions/
 *
 * TimeComplexity - O(n)
 * SpaceComplexity - O(1)
 */
public class SignOfProductOfArray1822 {

    public int (int[] nums) {

        double product = 1;

        //int positiveNumbersCount = 0;

        for(int i=0; i < nums.length; i++) {
            //System.out.println("number ->" + nums[i]);
            //System.out.println("before product ->" + product);
            product *= nums[i];
            //System.out.println("after product ->" + product);
        }

        //System.out.println(product);

        if (product > 0) {
            return 1;
        }
        else if (product < 0) {
            return -1;
        }
        else {
            return 0;
        }
    }

    public static void main(String[] args) {

        //int[] nums = {-1,-2,-3,-4,3,2,1};

        //int[] nums = {1,5,0,2,-3};

        //int[] nums = {-1,1,-1,1,-1};

        //int[] nums = {41,65,14,80,20,10,55,58,24,56,28,86,96,10,3,84,4,41,13,32,42,43,83,78,82,70,15,-41};

        int[] nums = {9,72,34,29,-49,-22,-77,-17,-66,-75,-44,-30,-24};

        SignOfProductOfArray1822 obj = new SignOfProductOfArray1822();
        System.out.println(obj.(nums));
    }
}
