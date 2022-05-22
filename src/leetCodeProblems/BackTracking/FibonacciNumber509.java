package leetCodeProblems.BackTracking;

/**
 * LeetCode - https://leetcode.com/problems/fibonacci-number/
 */

public class FibonacciNumber509 {

    public int fib(int n) {

        if (n == 0) {
            return 0;
        }

        if (n == 1) {
            return 1;
        }

        return fib(n-1) + fib(n-2);
    }

    public static void main(String[] args) {

        FibonacciNumber509 obj = new FibonacciNumber509();

        System.out.println(obj.fib(4));
    }
}
