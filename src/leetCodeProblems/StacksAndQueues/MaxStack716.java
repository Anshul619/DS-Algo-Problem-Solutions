package leetCodeProblems.StacksAndQueues;

/**
 * LeetCode - https://leetcode.com/problems/max-stack/
 *
 * Time-Complexity of all operations ( except popMax ) - O(1) time
 */

import java.util.Stack;

public class MaxStack716 {

    Stack<Integer> mainStack;
    Stack<Integer> monotonicStack;

    public MaxStack716() {
        mainStack = new Stack<>();
        monotonicStack = new Stack<>();
    }

    public void push(int x) {

        if (!monotonicStack.isEmpty()) {
            if (monotonicStack.peek() <= x) {
                monotonicStack.push(x);
            }
        }
        else {
            monotonicStack.push(x);
        }

        //System.out.println(monotonicStack);
        mainStack.push(x);
    }

    public int pop() {
        int element = mainStack.pop();

        if (!monotonicStack.isEmpty()) {
            if (monotonicStack.peek() == element) {
                monotonicStack.pop();
            }
        }

        return element;
    }

    public int top() {
        return mainStack.peek();
    }

    public int peekMax() {
        if (monotonicStack.isEmpty()) {
            return -1;
        }

        return monotonicStack.peek();

    }

    public int popMax() {

        int max = peekMax();

        Stack<Integer> buffer = new Stack();

        while (top() != max) {
            buffer.push(pop());
        }

        pop();

        while (!buffer.isEmpty()) {
            push(buffer.pop());
        }

        return max;

    }

    public static void main(String[] args) {
        MaxStack716 stack = new MaxStack716();

        stack.push(1);
        stack.push(4);
        stack.push(5);
        stack.push(10);
        stack.push(8);

        System.out.println(stack.pop());
        System.out.println(stack.peekMax());

    }
}
