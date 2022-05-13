package leetCodeProblems.TwoDArrayMatrix;

/**
 * Leet Code Link - https://leetcode.com/problems/set-matrix-zeroes/
 * 
 * Interview Bit Link - https://www.interviewbit.com/problems/set-matrix-zeros/
 * 
 * @author anshul.agrawal
 *
 */

public class SetMatrixZeros73 {
	static int[][] setZerosRow(int[][] matrix, int rowNumber, int startColNumber) {
        
        int noOfColumns = matrix[rowNumber].length;
        
        for(int j=startColNumber; j < noOfColumns;j++) {
            matrix[rowNumber][j] = 0;
        }
        
        return matrix;
    }
    
    static int[][] setZerosColumn(int[][] matrix, int columnNumber, int startRowNumber) {
        
        int noOfRows = matrix.length;
        
        for(int i=startRowNumber; i < noOfRows; i++) {
            matrix[i][columnNumber] = 0;
        }
        
        return matrix;
    }
    
    public void setZeroes(int[][] matrix) {
        
        boolean isFirstRowZero = false;
        boolean isFirstColumnZero = false;
        
        int noOfRows = matrix.length;
        int noOfColumns = matrix[0].length;
        
        // Traverse the first row & set isFirstRowZero if its supposed to be zero
        for(int j=0; j < noOfColumns; j++) {
            if (matrix[0][j] == 0) {
                isFirstRowZero = true;
            }
        }
        
         // Traverse the first column & set isFirstColumnZero if its supposed to be zero
        for(int i=0; i < noOfRows; i++) {
            if (matrix[i][0] == 0) {
                isFirstColumnZero = true;
            }
        }
        
        // Traverse all rows & columns ( except first and last ) & mark first cell/row ZERO based on the condition
        for(int i=1; i< noOfRows; i++) {
            for(int j=1; j < noOfColumns;j++) {
                
                if (matrix[i][j] == 0) {
                    matrix[i][0] = 0;
                    matrix[0][j] = 0;
                }
            }
        }
        
        // Traverse the first row and set its columns ( not first column ) to ZERO based on condition
        for(int j=1; j < noOfColumns; j++) {
            if (matrix[0][j] == 0) {

                setZerosColumn(matrix, j, 1);
                
            }
        }
        
        // Traverse the first column and set its rows ( not first ROW ) to ZERO based on condition
        for(int i=1; i < noOfRows; i++) {
            
            if (matrix[i][0] == 0) {
                
                setZerosRow(matrix, i, 1);
                
            }
        }
        
        if (isFirstRowZero) {
            setZerosRow(matrix, 0, 0);
        }
        
        if (isFirstColumnZero) {
            setZerosColumn(matrix, 0, 0);
        }
    }

}
