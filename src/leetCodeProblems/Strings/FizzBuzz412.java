package leetCodeProblems.Strings;

/**
 * https://leetcode.com/problems/fizz-buzz/submissions/
 * TimeComplexity - O(n)
 * Space Complexity - O(1)
 */

import java.util.ArrayList;
import java.util.List;

public class FizzBuzz412 {

    public List<String> fizzBuzz(int n) {

        List<String> output = new ArrayList<>();

        for(int i =1; i<=n; i++) {

            if (i%3 == 0 && i%5 ==0) {
                output.add("FizzBuzz");
            }
            else if(i%3 == 0) {
                output.add("Fizz");
            }
            else if(i%5 == 0) {
                output.add("Buzz");
            }
            else {
                output.add(String.valueOf(i));
            }
        }

        return output;
    }

    public static void main(String[] args) {

        FizzBuzz412 obj = new FizzBuzz412();
        System.out.println(obj.fizzBuzz(3));
        System.out.println(obj.fizzBuzz(5));
        System.out.println(obj.fizzBuzz(15));
    }
}
