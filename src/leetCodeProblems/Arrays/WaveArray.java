package leetCodeProblems.Arrays;

/**
 * Interview Bit - https://www.interviewbit.com/problems/wave-array/
 */

import java.util.*;

public class WaveArray {
	public int[] wave(int[] A) {

        Arrays.sort(A);

        for(int i = 0; i < (A.length-1); i+=2) {

            int temp = A[i];
            A[i] = A[i+1];
            A[i+1] = temp;
        }

        return A;
    }

}
