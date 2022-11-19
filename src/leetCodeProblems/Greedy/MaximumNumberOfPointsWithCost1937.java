package leetCodeProblems.GreedyAlgo;

/**
 * LeetCode - https://leetcode.com/problems/maximum-number-of-points-with-cost/
 */
public class MaximumNumberOfPointsWithCost1937 {

    public long maxPoints(int[][] points) {

        int answer = 0;

        int previousMaxColumnIndex = -1;

        int localCurrentGain = 0;

        int currentRowMaxGain = 0;
        int currentRowMaxColumnIndex = 0;

        for (int i=0; i < points.length; i++) {

            for (int j=0; j < points[i].length; j++) {

                if (i != 0) {
                    localCurrentGain = points[i][j] - Math.abs(j-previousMaxColumnIndex);
                }
                else {
                    localCurrentGain = points[i][j];
                }

                System.out.println("localCurrentGain ->" + localCurrentGain);
                if (currentRowMaxGain < localCurrentGain) {
                    currentRowMaxGain = localCurrentGain;
                    currentRowMaxColumnIndex = j;
                }
            }


            if (points[i].length > 0) {

                System.out.println("-------");
                System.out.println("Row Index --->" + i);
                System.out.println("rowMaxGain ->" + currentRowMaxGain);
                System.out.println("rowMaxColumnIndex ->" + currentRowMaxColumnIndex);

                answer += currentRowMaxGain;
                previousMaxColumnIndex = currentRowMaxColumnIndex;

                currentRowMaxGain = 0;
                System.out.println("answer ->" + answer);
                System.out.println("-------");
            }

        }

        return answer;
    }

    public static void main(String[] args) {

        //int[][] points = {{1,2,3}, {1,5,1}, {3,1,1}}; //Expected o/p = 9

        //int[][] points = {{1,5}, {2,3}, {4,2}}; // Expected o/p = 11

        //int[][] points = {{0,3,0,4,2}, {5,4,2,4,1}, {5,0,0,5,1}, {2,0,1,0,3}}; // Expected o/p = 15

        // [[2,4,0,5,5],[0,5,4,2,5],[2,0,2,3,1],[3,0,5,5,2]]
        int[][] points = {{2,4,0,5,5}, {0,5,4,2,5}, {2,0,2,3,1}, {3,0,5,5,2}}; //Expected o/p = 9

        MaximumNumberOfPointsWithCost1937 obj = new MaximumNumberOfPointsWithCost1937();
        System.out.println(obj.maxPoints(points));
    }
}
