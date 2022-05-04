package leetCodeProblems.StacksAndQueues;

/**
 * Reference - https://www.educba.com/stack-vs-queue/
 * @author anshul.agrawal
 *
 */

public class QueueRawImpl {
	private int maxSize;
	private String[] stackArray;
	private int first = 0;
	private int last = -1;

	public QueueRawImpl(int s) {
		maxSize = s;
		stackArray = new String[maxSize];
	}

	public void push(String j) {
		last += 1;
		stackArray[last] = j;
	}

	public String pop() {
		String temp = stackArray[first];
		first += 1;
		return temp;
	}

	public String peek() {
		return stackArray[first];
	}

	public boolean isEmpty() {
		return (last == -1);
	}

	public boolean isFull() {
		return (last == maxSize - 1);
	}

}
