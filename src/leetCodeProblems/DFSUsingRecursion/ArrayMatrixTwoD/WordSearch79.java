package leetCodeProblems.BackTracking;

/**
 * LeetCode Link - https://leetcode.com/problems/word-search/
 * 
 * @author anshul.agrawal
 *
 */
public class WordSearch79 {

	static int rows;
    static int cols;

    static int[] dx = {0, 1, 0, -1};
    static int[] dy = {1, 0, -1, 0};
    
    public boolean backTrace(char[][] board, String word, int i, int j, int wIndex) {

        // This means, all words are traversed successfully
        if (wIndex >= word.length()) {
            return true;
        }

        if (i < 0 || j < 0 || i >= rows || j >= cols || word.charAt(wIndex) != board[i][j]) {
            return false;
        }
        else {
        	//Nothing
        }
        
        char matched = board[i][j];
        board[i][j] = '$'; // Since this is matched, hence replace with "#" to remove unrequired traversals

        for (int z=0; z<4; z++) {
            if (backTrace(board, word, i+dx[z], j+dy[z], wIndex+1)) {
                return true;
            }
        }

        board[i][j] = matched;

        return false;
    }
    
    public boolean patternSearch(char[][] board, String word) {

        rows = board.length;
        cols = board[0].length;

        for(int i=0; i < rows; i++) {
            for(int j=0; j < cols; j++) {
                if (backTrace(board, word, i, j, 0)){
                    return true;
                }
            }
        }

        return false;

    }
    
    /**
     * Driver method.
     * 
     * @param board
     * @param word
     * @return
     */
    public boolean exist(char[][] board, String word) {
    	
    	WordSearch79 obj = new WordSearch79();
    	
        return obj.patternSearch(board, word);
    }

}
