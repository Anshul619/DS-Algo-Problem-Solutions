package leetCodeProblems.StacksAndQueues;

/**
 * This is HARD problem.
 * 
 * LeetCode - https://leetcode.com/problems/longest-valid-parentheses/
 * 
 */
import java.util.*;

public class LongestValidParenthesis32 {
	public int longestValidParentheses(String s) {

		Stack<Integer> stack = new Stack<>();
		stack.push(-1);

		int result = 0;

		for (int i = 0; i < s.length(); i++) {

			if (s.charAt(i) == '(') {
				stack.push(i);
			} else {
				if (!stack.empty()) {
					stack.pop();
				}

				if (!stack.empty()) {
					result = Math.max(result, i - stack.peek());
				} else {
					stack.push(i);
				}
			}
		}

		return result;
	}
}
