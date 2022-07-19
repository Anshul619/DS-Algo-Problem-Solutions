package leetCodeProblems.ArrayMatrixTwoD;

/**
 * LeetCode - https://leetcode.com/problems/transpose-matrix/solution/
 *
 * TimeComplexity - O(R*C)
 * SpaceComplexity - O(R*C) where R denotes number of rows, C denotes number of columns
 */

import java.util.Arrays;

public class TransposeMatrix867 {

    private void printMatrix(int[][] matrix) {

        System.out.println("printMatrix---");

        for(int i=0; i < matrix.length; i++) {
            System.out.println(Arrays.toString(matrix[i]));
        }
    }

    public int[][] transpose(int[][] matrix) {

        int rowsCount = matrix.length;
        int columnsCount = matrix[0].length;

        int[][] output = new int[columnsCount][rowsCount];

        for(int i=0; i < matrix.length; i++) {

            for(int j=0; j < matrix[i].length; j++) {
                output[j][i] = matrix[i][j];
            }
        }

        return output;
    }

    public static void main(String[] args) {

        int[][] input = {{1,2,3}, {4,5,6}};
        //int[][] input = {{1,2,3}, {4,5,6}, {7,8,9}};

        TransposeMatrix867 obj = new TransposeMatrix867();
        int[][] output = obj.transpose(input);

        obj.printMatrix(output);
    }
}