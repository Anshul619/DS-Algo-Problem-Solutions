package leetCodeProblems.StacksAndQueues;

public class StackCheckIfBalancedParenthesis {

	public int solve(String A) {

		ParenthesesStack stack = new ParenthesesStack(A.length());

		for (int i = 0; i < A.length(); i++) {

			if (A.charAt(i) == '(') {
				stack.push("(");
			} else {
				if (!stack.isEmpty()) {
					stack.pop();
				} else {
					return 0;
				}
			}
		}

		if (stack.isEmpty()) {
			return 1;
		} else {
			return 0;
		}
	}

	static class ParenthesesStack {
		private int maxSize;
		private String[] stackArray;
		private int top;

		public ParenthesesStack(int s) {
			maxSize = s;
			stackArray = new String[maxSize];
			top = -1;
		}

		public void push(String j) {
			top += 1;
			stackArray[top] = j;
		}

		public String pop() {
			String temp = stackArray[top];
			top -= 1;
			return temp;
		}

		public String peek() {
			return stackArray[top];
		}

		public boolean isEmpty() {
			return (top == -1);
		}

		public boolean isFull() {
			return (top == maxSize - 1);
		}
	}

}
