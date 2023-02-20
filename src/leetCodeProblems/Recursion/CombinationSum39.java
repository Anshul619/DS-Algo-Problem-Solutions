package leetCodeProblems.BackTracking;

/**
 * LeetCode - https://leetcode.com/problems/combination-sum/
 * InterviewBit - https://www.interviewbit.com/problems/combination-sum/
 */

import java.util.*;

public class CombinationSum39 {
	
	List<List<Integer>> result = new ArrayList<List<Integer>>();
    ArrayList<Integer> currentCombination = new ArrayList<Integer>();
    
    public int statsAllLoopCounts = 0;
    
    /**
     * Construct the new array list using a base array list
     * 
     * @param baseArrayList
     * @return
     */
    private ArrayList<Integer> getNewArrayList(ArrayList<Integer> baseArrayList) {
        ArrayList<Integer> newArrayList = new ArrayList<Integer>();
        
        for(int i=0; i< baseArrayList.size(); i++) {
            newArrayList.add(baseArrayList.get(i));
        }
        
        return newArrayList;
    }
    
    /**
     * Find the combination sum recursively using a target sum.
     * 
     * @param candidates
     * @param targetSum
     * @param index
     */
    private void findCombinationRecursively(int[] candidates, int targetSum, int index){
        
        if (targetSum == 0) {
            //System.out.println("targetSum is ZERO. Hence insert.");
            //System.out.println("Current Combination -> " + currentCombination);
         
            result.add(getNewArrayList(currentCombination));
            //System.out.println("Result ->" + result);
            return;
        }
        
        for(int i = index; i < candidates.length; i++) {
            
            int subSum = targetSum - candidates[i];
            
            statsAllLoopCounts++;
            
            // This element can be used for the sum
            if (subSum >= 0) {
                
                currentCombination.add(candidates[i]);
                
                /*System.out.println("Before Recursion - Start");
                System.out.println("Before Recursion - Current Combination ->" + currentCombination);
                System.out.println("Before Recursion - Current Element ->" + candidates[i]);
                System.out.println("Before Recursion - Current Index ->" + i);
                System.out.println("Before Recursion - End");*/
                
                findCombinationRecursively(candidates, subSum, i);
                
                /*System.out.println("After Recursion - Start");
                System.out.println("After Recursion - Current Index ->" + i);
                
                System.out.println("After Recursion - Current Combination ->" + currentCombination);
                System.out.println("After Recursion - Current Element ->" + candidates[i]);
                System.out.println("After Recursion - End");*/
                
                if (currentCombination.contains(candidates[i])) {                    
                    currentCombination.remove(new Integer(candidates[i]));
                }
            }
        }
        
        return;
        
    }
    
    /**
     * Driver method
     * 
     * @param candidates
     * @param target
     * @return
     */
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        
        result = new ArrayList<List<Integer>>();
        currentCombination = new ArrayList<Integer>();
        
        findCombinationRecursively(candidates, target, 0);

        return result;   
    }
    
    /**
     * Main method
     * 
     * @param args
     */
    public static void main(String[] args) {
    	
    	//int[] inputArray = {2, 3, 6, 7};
    	//int targetSum = 7;
    	
    	//int[] inputArray = {2, 3, 5};
    	//int targetSum = 8;

        int[] inputArray = {1, 2};
        int targetSum = 3;
    			
    	CombinationSum39 obj = new CombinationSum39();
    	
    	System.out.println(obj.combinationSum(inputArray, targetSum));
    	
    	System.out.println("Post iteration - statsAllLoopCounts ->"+ obj.statsAllLoopCounts);
    }

}
