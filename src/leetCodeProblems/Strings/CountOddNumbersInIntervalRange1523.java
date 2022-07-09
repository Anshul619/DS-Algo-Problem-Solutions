package leetCodeProblems.Strings;

/**
 * LeetCode - https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/submissions/
 *
 * TimeComplexity - O(1)
 */
public class CountOddNumbersInIntervalRange1523 {

    public int countOdds(int low, int high) {

        int ans = (high-low)/2;

        if (low%2 == 1 || high%2 == 1) {
            ans++;
        }

        return ans;
    }

    public static void main(String[] args) {

        CountOddNumbersInIntervalRange1523 obj = new CountOddNumbersInIntervalRange1523();

        System.out.println(obj.countOdds(3, 7));

        System.out.println(obj.countOdds(278382788, 569302584));
    }
}
