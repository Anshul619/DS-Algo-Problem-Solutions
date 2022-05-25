package leetCodeProblems.StacksAndQueues;

/**
 * LeetCode - https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/solution/
 */

import java.util.HashSet;
import java.util.Stack;

public class MinRemovalForValidParentheses1249 {

    static class ParenthesisEntity {

        char parenChar;
        int index;

        ParenthesisEntity(char parenChar, int index) {
            this.parenChar = parenChar;
            this.index = index;
        }
    }
    public String minRemoveToMakeValid(String s) {

        Stack<ParenthesisEntity> stack = new Stack<>();

        HashSet<Integer> toBeRemovedList = new HashSet<>();
        String outputString = "";

        for(int i=0; i< s.length(); i++) {

            if (s.charAt(i) == '(') {
                ParenthesisEntity pe = new ParenthesisEntity('(', i);
                stack.add(pe);
            }
            else if(s.charAt(i) == ')') {

                if (!stack.isEmpty() && stack.peek().parenChar == '(') {
                    stack.pop();
                }
                else {
                    toBeRemovedList.add(i);
                }
            }
        }

        while(!stack.isEmpty()) {
            toBeRemovedList.add(stack.pop().index);
        }

        //System.out.println(toBeRemovedList);

        for (int i=0; i < s.length(); i++) {

            if (!toBeRemovedList.contains(i)) {
                outputString += s.charAt(i);
            }
        }

        return outputString;
    }

    public static void main(String[] args) {
        String s = "lee(t(c)o)de)";

        //String s = "a)b(c)d";

        MinRemovalForValidParentheses1249 obj = new MinRemovalForValidParentheses1249();

        System.out.println(obj.minRemoveToMakeValid(s));
    }

}
