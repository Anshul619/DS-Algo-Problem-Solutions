package leetCodeProblems.GreedyAlgo;

/**
 * LeetCode - https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
 * InterviewBit - https://www.interviewbit.com/problems/best-time-to-buy-and-sell-stocks-i/
 * 
 * @author anshul.agrawal
 *
 */
public class BestTimeToBuyAndSellStockI121 {
	
    public int maxProfit(int[] prices) {
        
    	int maxProfit = 0;
        int buy = 0;
        
        if (prices.length > 0) {
            buy = prices[0];
        }
        
        for (int i=0; i < prices.length; i++) {
            
            // Check for minimum buy prices
            if (buy > prices[i]) {
                buy = prices[i];
            }
            else if (prices[i] - buy > maxProfit) { // Else check for max profit
                maxProfit = prices[i] - buy;
            }
        }
        
        return maxProfit;
    }
    
    public static void main(String[] args) {
        
    	int[] prices = {7, 1, 5, 3, 6, 4};
        
        //int[] prices = {7, 6, 4, 3, 1};
        
        BestTimeToBuyAndSellStockI121 obj = new BestTimeToBuyAndSellStockI121();
        System.out.println(obj.maxProfit(prices));
    }
    

}
