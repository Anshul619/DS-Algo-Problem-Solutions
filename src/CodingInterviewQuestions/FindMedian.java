package CodingInterviewQuestions;

import java.util.Arrays;

/**
 * Asked in Google, 30thJune2022 - PreQualifier Round
 */

public class FindMedian {

    private int calculateMedian(int[] input) {

        int length = input.length;

        if (length ==0) {
            return 0;
        }


        if (length%2 == 0) { // Array length is even

            int sum = 0;

            for(int i=0; i < length; i++) {
                sum+= input[i];
            }

            return sum/2;

        }
        else { // array length is odd
            Arrays.sort(input);

            return input[length/2];
        }
    }

    public static void main(String[] args) {

        FindMedian obj = new FindMedian();

        //int[] input = {1, 2, 3, 5, 10}; //output = 3, middle of the array
        int[] input = {1, 2, 3, 5, 10, 15}; //output = 18, mean of the array

        System.out.println(obj.calculateMedian(input));
    }
}
