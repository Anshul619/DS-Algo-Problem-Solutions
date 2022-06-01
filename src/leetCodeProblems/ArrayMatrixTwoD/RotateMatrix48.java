package leetCodeProblems.ArrayMatrixTwoD;

/**
 * LeetCode - https://leetcode.com/problems/rotate-image/
 * 
 * @author anshul.agrawal
 *
 */
public class RotateMatrix48 {
	
	// Function for print matrix
	static void printMatrix(int arr[][])
	{
		 // This is SQUARE matrix.
        int N = arr.length;
        
	    for (int i = 0; i < N; i++)
	    {
	        for (int j = 0; j < N; j++)
	        System.out.print( arr[i][j] + " ");
	        System.out.println();
	    }
	}
	
	static void rotate90ClockwiseInPlace(int[][] matrix) {
        
        // This is SQUARE matrix.
        int N = matrix.length;
        
        // Because once one half is rotated, other half would be rotated automatically.
        int numberOfCyclesNeeded = N/2;
        
        for (int i=0; i < numberOfCyclesNeeded; i++) {
            
            // Its start & end indexes have to be based on i. This is needed to remove unrequired rotations.
            for(int j=i; j < N-i-1; j++) {
                
                int temp = matrix[i][j];
                matrix[i][j] = matrix[N-j-1][i];
                matrix[N-j-1][i] = matrix[N-i-1][N-j-1];
                matrix[N-i-1][N-j-1] = matrix[j][N-i-1];
                matrix[j][N-i-1] = temp;
                
            }
        }
        
        return;
    }
	
	// Driver code
	 
    public static void main (String[] args)
    {
        int arr[][] = { { 1, 2, 3, 4 },
                  { 5, 6, 7, 8 },
                  { 9, 10, 11, 12 },
                  { 13, 14, 15, 16 } };
        rotate90ClockwiseInPlace(arr);
        printMatrix(arr);
    }

}
