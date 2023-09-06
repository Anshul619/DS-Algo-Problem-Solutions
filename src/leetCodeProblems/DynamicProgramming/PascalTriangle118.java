package leetCodeProblems.MathCalculations;

/**
 * LeetCode - https://leetcode.com/problems/pascals-triangle/
 */

import java.util.ArrayList;
import java.util.List;

public class PascalTriangle118 {

    public List<List<Integer>> generate(int numRows) {

        List<List<Integer>> output = new ArrayList<List<Integer>>();

        for(int i=1; i <= numRows; i++) {

            List<Integer> temp = new ArrayList<>();
            int c = 1;

            for(int j=0; j < i; j++) {

                if (j == 0) {
                    c = 1;
                }
                else {
                    c = c*(i-j)/j;
                }

                temp.add(c);
            }

            output.add(temp);
        }

        return output;
    }

    public static void main(String[] args) {

        PascalTriangle118 obj = new PascalTriangle118();

        System.out.println(obj.generate(5));
    }
}
