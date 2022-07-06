package leetCodeProblems.StacksAndQueues;

/**
 * LeetCode - https://leetcode.com/problems/min-stack/
 *
 * TimeComplexity - O(1)
 */

import java.util.Stack;

public class MinStack155 {

    Stack<Integer> mainStack;
    Stack<Integer> monotonicStack;

    public MinStack155() {
        mainStack = new Stack<>();
        monotonicStack = new Stack<>();
    }

    public void push(int val) {
        if (!monotonicStack.isEmpty()) {
            if (monotonicStack.peek() >= val) {
                monotonicStack.push(val);
            }
        }
        else {
            monotonicStack.push(val);
        }

        //System.out.println(monotonicStack);
        mainStack.push(val);
    }

    public void pop() {
        int element = mainStack.pop();

        if (!monotonicStack.isEmpty()) {
            if (monotonicStack.peek() == element) {
                monotonicStack.pop();
            }
        }

        return;
    }

    public int top() {
        return mainStack.peek();
    }

    public int getMin() {
        if (monotonicStack.isEmpty()) {
            return -1;
        }

        return monotonicStack.peek();
    }
}
