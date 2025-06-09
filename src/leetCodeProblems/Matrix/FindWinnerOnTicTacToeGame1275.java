package leetCodeProblems.ArrayMatrixTwoD;

/**
 * LeetCode - https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game/solution/
 *
 * TimeComplexity - O(m*n)
 * SpaceComplexity - O(n2)
 */
public class FindWinnerOnTicTacToeGame1275 {

    private String getWinnerName(int winner) {

        switch (winner) {
            case 0:
                return "NoWinner";
            case 1:
                return "A";
            case -1:
                return "B";
        }
        return "NoWinner";
    }

    private int checkDiagonal(int[][] board) {

        int currentPlayer = board[0][0];
        int currentPlayerWinnerCount = 1;

        for(int i=1; i < 3; i++) {

            if (board[i][i] == currentPlayer) {
                currentPlayerWinnerCount++;
            }
        }

        if (currentPlayerWinnerCount == 3) {
            return currentPlayer;
        }

        return 0;
    }

    private int checkAntiDiagonal(int[][] board) {

        int currentPlayer = board[0][2];

        if (board[1][1] == currentPlayer && board[2][0] == currentPlayer) {
            return currentPlayer;
        }

        return 0;
    }

    private int checkRows(int[][] board) {

        for(int i=0; i < 3; i++) {
            int currentPlayer = board[i][0];

            if (board[i][1] == currentPlayer && board[i][2] == currentPlayer) {
                return currentPlayer;
            }
        }

        return 0;
    }

    private int checkColumns(int[][] board) {

        for(int i=0; i < 3; i++) {
            int currentPlayer = board[0][i];

            if (board[1][i] == currentPlayer && board[2][i] == currentPlayer) {
                return currentPlayer;
            }
        }

        return 0;
    }

    public String tictactoe(int[][] moves) {

        int[][] board = new int[3][3];

        for(int i=0; i < moves.length; i++) {
            if (i%2 == 0) {
                board[moves[i][0]][moves[i][1]] = 1; //A player
            }
            else {
                board[moves[i][0]][moves[i][1]] = -1; //B player
            }
        }

        String diagonalWinner = getWinnerName(checkDiagonal(board));

        if (!diagonalWinner.equals("NoWinner")) {
            return diagonalWinner;
        }

        String antiDiagonalWinner = getWinnerName(checkAntiDiagonal(board));

        if (!antiDiagonalWinner.equals("NoWinner")) {
            return antiDiagonalWinner;
        }

        String rowWinner = getWinnerName(checkRows(board));

        if (!rowWinner.equals("NoWinner")) {
            return rowWinner;
        }

        String columnWinner = getWinnerName(checkColumns(board));

        if (!columnWinner.equals("NoWinner")) {
            return columnWinner;
        }

        if (moves.length < 9) {
            return "Pending";
        }

        return "Draw";
    }

    public static void main(String[] args) {
        //int[][] moves = {{0,0}, {1,1}, {2,0}, {1,0}, {1,2}, {2,1}, {0,1}, {0,2}, {2,2}};

        //int[][] moves = {{0,0}, {2,0}, {1,1}, {2,1}, {2,2}};

        //int[][] moves = {{0,0}, {1,1}, {0,1}, {0,2}, {1,0}, {2,0}};

        int[][] moves = {{1,2},{2,1},{1,0},{0,0},{0,1},{2,0},{1,1}};

        FindWinnerOnTicTacToeGame1275 obj = new FindWinnerOnTicTacToeGame1275();

        System.out.println(obj.tictactoe(moves));
    }
}
