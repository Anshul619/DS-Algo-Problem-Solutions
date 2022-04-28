package leetCodeProblems.String;

public class StringMinimumParentheses {
	public int minInsertions(String s) {

		int c_need = 0;
		int ans = 0;

		for (int i = 0; i < s.length(); i++) {

			if (s.charAt(i) == '(') {
				c_need += 1;

				// This is again tricky and IMPORTANT code.
				// There was a closing bracket needed, hence save it separately ( by
				// incrementing ans ) & decrementing c_need.
				// Needed for “())” - Minimum Insertions to balance a parentheses string “())”
				/*
				 * if (c_need%2 == 1) { c_need --; ans++; }
				 */
			} else {
				c_need -= 1;

				// This is TRICKY and IMPORTANT CODE.
				// To reset the variables. This is the only closing bracket.
				if (c_need == -1) {
					c_need = 0;
					ans++; // an open bracket is needed here. Hence incrementing "ans".
				}
			}
		}

		return c_need + ans;
	}

}
