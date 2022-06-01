package leetCodeProblems.MathCalculations;

/**
 * LeetCode - https://leetcode.com/problems/add-strings/
 * Time Complexity - O(n)
 * Space Complexity - O(1)
 */
public class AddStrings415 {

    public String addStrings(String num1, String num2) {

        int num1Pointer = num1.length()-1;
        int num2Pointer = num2.length()-1;
        int carry = 0;

        int localOutputNumb = 0;
        int localNum1Val = 0;
        int localNum2Val = 0;

        String output = "";

        while (num1Pointer >= 0 || num2Pointer >= 0) {

            localNum1Val = 0;
            localNum2Val = 0;

            if (num1Pointer >= 0) {
                localNum1Val = Integer.parseInt(String.valueOf(num1.charAt(num1Pointer)));
            }

            if (num2Pointer >= 0) {
                localNum2Val = Integer.parseInt(String.valueOf(num2.charAt(num2Pointer)));
            }

            localOutputNumb = localNum1Val + localNum2Val + carry;

            output = localOutputNumb%10 + output;
            carry = localOutputNumb/10;

            num1Pointer--;
            num2Pointer--;
        }

        if (carry != 0) {
            output = carry + output;
        }

        return output;
    }

    public static void main(String[] args) {

        String num1 = "11";
        String num2 = "123";
        AddStrings415 obj = new AddStrings415();
        //System.out.println(obj.addStrings(num1, num2)); // expected o/p = 134

        num1 = "456";
        num2 = "77";

        System.out.println(obj.addStrings(num1, num2)); // expected o/p = 533

        //num1 = "0";
        //num2 = "0";

        //System.out.println(obj.addStrings(num1, num2)); // expected o/p = 0

    }
}
