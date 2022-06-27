package leetCodeProblems.BitsManupulation;

/**
 * LeetCode - https://leetcode.com/problems/single-number/
 *
 * Time Complexity - O(n)
 * Space Complexity - O(1)
 */
public class SingleNumber136 {

    public int singleNumber(int[] nums) {

        int ansXOR = 0;

        for (int i=0; i < nums.length; i++) {
            //System.out.println("Before ->" + ansXOR);
            ansXOR = ansXOR ^ nums[i];
            //System.out.println("After ->" + ansXOR);
        }

        return ansXOR;
    }

    public static void main(String[] args) {

        SingleNumber136 obj = new SingleNumber136();

        int[] nums = {4,1,2,1,2};
        System.out.println(obj.singleNumber(nums));
    }
}
