package leetCodeProblems.DynamicProgramming;

/**
 * LeetCode - https://leetcode.com/problems/climbing-stairs/submissions/
 */

public class ClimbingStairs70 {

    int[] dp;

    public int fib(int n) {

        if (dp[n] != -1) {
            return dp[n];
        }

        if (n == 0) {
            dp[0] = 0;
            return 0;
        }

        if (n == 1) {
            dp[1] = 1;
            return 1;
        }

        int fibN = fib(n-1) + fib(n-2);

        dp[n] = fibN;

        return fibN;
    }

    public int climbStairs(int n) {

        dp = new int[n+2];

        for(int i=0; i<dp.length; i++) {
            dp[i] = -1;
        }

        return fib(n+1);
    }

    public static void main(String[] args) {

        ClimbingStairs70 obj = new ClimbingStairs70();

        System.out.println(obj.climbStairs(44));
    }
}
