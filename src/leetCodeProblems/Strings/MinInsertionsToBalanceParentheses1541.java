package leetCodeProblems.Strings;

public class MinInsertionsToBalanceParentheses1541 {
	public int minInsertions(String s) {

		int closingBracketNeeded = 0;
		int ans = 0;

		for (int i = 0; i < s.length(); i++) {

			if (s.charAt(i) == '(') {
				closingBracketNeeded += 2;

				// This is tricky and IMPORTANT code.
				// There was a closing bracket needed, hence we save it separately ( by
				// incrementing ans ) & decrementing c_need.
				// Needed for “())” - Minimum Insertions to balance a parentheses string “())”
				if (closingBracketNeeded%2 == 1) { 
					closingBracketNeeded --; 
					ans++;
				}
				
			} else {
				closingBracketNeeded -= 1;

				// This is TRICKY and IMPORTANT CODE.
				// To reset the variables. This is the only closing bracket.
				if (closingBracketNeeded == -1) {
					closingBracketNeeded = 1;
					ans++; // an open bracket is needed here. Hence incrementing "ans".
				}
			}
		}

		return closingBracketNeeded + ans;
	}

}
