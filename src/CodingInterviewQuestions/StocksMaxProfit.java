package CodingInterviewQuestions;

/**
 * Asked in Oracle Coding Round@17May2022
 */

import java.util.Comparator;
import java.util.Arrays;

public class StocksMaxProfit {

    static class StockValComparison implements Comparator<int[]> {

        public int compare(int[] a, int[] b) {
            return a[0] - b[0];
        }
    }

    /**
     * This would calculate the maximum profit, based on the current savings.
     *
     * Total time complexity = O(nLogn)
     * Total space complexity = 2n
     *
     * @param savings
     * @param currentValue
     * @param futureValue
     * @return
     */
    public int maxProfit(int savings, int[] currentValue, int[] futureValue) {

        int[][] stockValuations = new int[currentValue.length][2];
        int maxProfit = 0;

        // O(n) time complexity
        for (int i=0; i < currentValue.length; i++) {
            stockValuations[i][0] = currentValue[i];
            stockValuations[i][1] = futureValue[i] - currentValue[i];
        }

        // O(nLogn) time complexity
        Arrays.sort(stockValuations, new StockValComparison());

        // O(n) time complexity
        for(int i=stockValuations.length-1; i >= 0; i--) {

            if (savings >= stockValuations[i][0] && stockValuations[i][1] > 0) {
                maxProfit += stockValuations[i][1];
                savings -= stockValuations[i][0];
            }
        }

        return maxProfit;
    }

    public static void main(String[] args){

        // TestCase1
        //int[] currentValue = {1, 2, 3, 4}; // Current Value of the stocks
        //int[] futureValue = {5, 6, 1, 0}; // Future value of the stocks, after 1 year
        //int totalSavings = 10; // Total current savings

        // TestCase2
        int[] currentValue = {10, 2, 3, 4}; // Current Value of the stocks
        int[] futureValue = {15, 6, 10, 1}; // Future value of the stocks, after 1 year
        int totalSavings = 13; // Total current savings. Expected Output = 12

        StocksMaxProfit obj = new StocksMaxProfit();

        int profit = obj.maxProfit(totalSavings, currentValue, futureValue);

        System.out.println(profit);

        //1606636800
        System.out.println(System.currentTimeMillis());
    }

}
