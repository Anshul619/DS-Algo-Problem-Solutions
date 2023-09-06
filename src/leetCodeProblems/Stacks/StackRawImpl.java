package leetCodeProblems.StacksAndQueues;

/**
 * Reference - https://www.educba.com/stack-vs-queue/
 * @author anshul.agrawal
 *
 */
public class StackRawImpl {
	private int maxSize;
	private String[] stackArray;
	private int top;

	public StackRawImpl(int s) {
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
