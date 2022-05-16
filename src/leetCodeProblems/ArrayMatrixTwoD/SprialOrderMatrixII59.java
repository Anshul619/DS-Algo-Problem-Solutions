package leetCodeProblems.ArrayMatrixTwoD;

public class SprialOrderMatrixII59 {
	
	public int[][] generateMatrix(int n) {
        
        if (n == 0) {
            return new int[0][0];
        }
        
        int[][] matrix = new int[n][n];
        
        // Clockwise direction traversal
        int[] directionX = {0, 1, 0,-1};
        int[] directionY = {1, 0, -1, 0};
        
        int rows = n;
        int columns = n;
        
        boolean[][] visited = new boolean[rows][columns]; 
        
        int direction = 0;
        
        int currentRowIndex = 0;
        int currentColumnIndex = 0;
        
        for(int i=0; i < rows*columns; i++) {
            
            matrix[currentRowIndex][currentColumnIndex] = i+1;
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
        
        return matrix;
    }

}
