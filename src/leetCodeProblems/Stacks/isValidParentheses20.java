package leetCodeProblems.StacksAndQueues;

/**
 * LeetCode Link - https://leetcode.com/problems/valid-parentheses/
 * InterviewBit Link - https://www.interviewbit.com/problems/balanced-parantheses/
 * 
 * @author anshul.agrawal
 *
 */
public class isValidParentheses20 {
	
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
	
	public boolean isValid(String s) {
        
        ParenthesesStack stack = new ParenthesesStack(s.length());

        for (int i = 0; i < s.length(); i++) {
            
        	if (s.charAt(i) == '(' ||
        		s.charAt(i) == '[' ||
        		s.charAt(i) == '{') {
                stack.push(String.valueOf(s.charAt(i)));
            }
            else {
                if (!stack.isEmpty()) {
                	if (s.charAt(i) == ')' && stack.peek().equals("(")) {
                		stack.pop();
                	}
                	else if (s.charAt(i) == ']' && stack.peek().equals("[")) {
                		stack.pop();
                	}
                	else if (s.charAt(i) == '}' && stack.peek().equals("{")) {
                		stack.pop();
                	}
                	else {
                		return false;
                	}
                }
                else {
                    return false;
                }
            }
        }

        if (stack.isEmpty()) {
            return true;
        }
        else {
            return false;
        }
        
    }
	
	public static void main(String[] args) {
		
		String inputString = "()[]{}";
		
		//String inputString = "()";
		
		//String inputString = "(]";
		
		//String inputString ="(])";
		
		isValidParentheses20 obj = new isValidParentheses20();
		
		System.out.println(obj.isValid(inputString));
		
		
	}

	

}
