package leetCodeProblems.TwoDArrayMatrix;

/**
 * LeetCode - https://leetcode.com/problems/spiral-matrix/
 * 
 */
import java.util.*;

public class SprialOrderMatrixI54 {
	public List<Integer> spiralOrder(int[][] matrix) {
        
        List<Integer> output = new ArrayList<Integer>();
        
        if (matrix.length == 0) {
            return output;
        }
        
        // Clockwise direction traversal
        int[] directionX = {0, 1, 0,-1};
        int[] directionY = {1, 0, -1, 0};
        
        int rows = matrix.length;
        int columns = matrix[0].length;
        
        boolean[][] visited = new boolean[rows][columns]; 
        
        int direction = 0;
        
        int currentRowIndex = 0;
        int currentColumnIndex = 0;
        
        for(int i=0; i < rows*columns; i++) {
            
            output.add(matrix[currentRowIndex][currentColumnIndex]);
            visited[currentRowIndex][currentColumnIndex] = true;
            
            int nextRowIndex = currentRowIndex + directionX[direction];
            int nextColumnIndex = currentColumnIndex + directionY[direction];
            
            // This is positive case.
            if ( 0 <= nextRowIndex &&
               nextRowIndex < rows &&
                 0 <= nextColumnIndex &&
               nextColumnIndex < columns &&
               !visited[nextRowIndex][nextColumnIndex]) {
                
                currentRowIndex = nextRowIndex;
                currentColumnIndex = nextColumnIndex;
                
            }
            else { // Otherwise change direction
                
                direction = (direction + 1 ) % 4; // This is IMPORTANT code.
                
                currentRowIndex = currentRowIndex + directionX[direction];
                currentColumnIndex = currentColumnIndex + directionY[direction];
                
            }  
        }
        
        return output;
    }
}
