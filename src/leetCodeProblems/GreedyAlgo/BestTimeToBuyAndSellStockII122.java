package leetCodeProblems.GreedyAlgo;

/**
 * LeetCode - https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
 * InterviewBit - https://www.interviewbit.com/problems/best-time-to-buy-and-sell-stocks-ii/
 * 
 * @author anshul.agrawal
 *
 */
public class BestTimeToBuyAndSellStockII122 {
	
	
	public int maxProfit(int[] prices) {
	     
        int currentProfit = 0;
        int currentBuy = 0;
        int totalProfit = 0;
        
        if (prices.length > 0) {
            currentBuy = prices[0];
        }
        
        for (int i=0; i<prices.length; i++) {
            
            // Check for minimum buy prices
            if (currentBuy > prices[i]) {
                currentBuy = prices[i];
                totalProfit += currentProfit;
                currentProfit = 0;
            }
            else if (prices[i] - currentBuy > currentProfit) { // Else check for max profit
                currentProfit = prices[i] - currentBuy;
            }
            else {
            	
                int predictedProfit = prices[i] - currentBuy;
                
                // Sell the stocks since predictedProfit is less than currentProfit.
                if (predictedProfit <= currentProfit) {
                    totalProfit += currentProfit;
                    currentBuy = prices[i];
                    currentProfit = 0;
                }
                
            }
            
            //print(currentProfit, currentBuy, totalProfit);
        }
        
        totalProfit += currentProfit;
        
        return totalProfit;
    }
    
	private void print(int currentProfit, int currentBuy, int totalProfit) {
		System.out.println("Local Profit -> " + currentProfit);
        System.out.println("Buy -> " + currentBuy);
        System.out.println("totalProfit -> " + totalProfit);
        System.out.println("----");
	}
	
	
    public static void main(String[] args) {
        int[] prices = {7, 1, 5, 3, 6, 4}; // Ans = 7
        
        //int[] prices = {1, 2, 3, 4, 5}; // Ans = 4
        
        //int[] prices = {7, 6, 4, 3, 1}; // Ans = 0
        
        //int[] prices = {2, 1, 2, 0, 1}; // Ans = 2
        
        //int[] prices = {2, 4, 1}; // Ans = 2
        
    	BestTimeToBuyAndSellStockII122 obj = new BestTimeToBuyAndSellStockII122();
        
        System.out.println(obj.maxProfit(prices));
    }
    

}
