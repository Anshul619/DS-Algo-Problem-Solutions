package CodingInterviewQuestions;

/**
 * Vonage - Demo test on 4th June, 2022
 * Link - https://app.codility.com/c/feedback/demo98QKYC-WV7/
 */

import java.util.Arrays;

public class SmallestPositiveInteger {

    public int solution(int[] A) {
        // write your code in Java SE 8

        Arrays.sort(A);

        int smallestPositiveNumber = 1;

        for(int i =0; i<A.length; i++) {

            if (A[i] > 0) {

                if (A[i] == smallestPositiveNumber) {
                    smallestPositiveNumber += 1;
                }
            }
        }

        return smallestPositiveNumber;
    }

    public static void main(String[] args) {
        //int[] input = {1, 3, 6, 4, 1, 2};

        //int[] input = {1, 2, 3};

        int[] input = {-1, -3};

        SmallestPositiveInteger obj = new SmallestPositiveInteger();
        System.out.println(obj.solution(input));
    }
}
