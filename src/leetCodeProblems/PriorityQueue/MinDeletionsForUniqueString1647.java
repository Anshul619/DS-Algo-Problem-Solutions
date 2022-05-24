package leetCodeProblems.PriorityQueue;

/**
 * LeetCode - https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/submissions/
 */

import java.util.*;


public class MinDeletionsForUniqueString1647 {
	    
    /**
     * minDeletions
     * 
     * @param s
     * @return
     */
    public int minDeletions(String s) {
        
    	HashMap<String, Integer> charCountMap = new HashMap<String, Integer>();
        
        PriorityQueue<Integer> countsPriorityQueue = new PriorityQueue<Integer>(Collections.reverseOrder());
        
        int minDeletions = 0;
        
        for (int i=0; i < s.length(); i++) {
            
            String characterStr = String.valueOf(s.charAt(i));
            
            if (charCountMap.containsKey(characterStr)) {
				
				int currentCount = charCountMap.get(characterStr);
				currentCount = currentCount+1;
				
				charCountMap.put(characterStr, currentCount);
			}
			else {
				charCountMap.put(characterStr, 1);
			}
        }
        
        for (Integer count: charCountMap.values()) {
        	countsPriorityQueue.add(count);
        }
        
        while (!countsPriorityQueue.isEmpty()) {
        	
        	int lastCount = countsPriorityQueue.remove(); 
        	
        	//System.out.println("last count ->" + lastCount);
        	
        	if (!countsPriorityQueue.isEmpty() && lastCount == countsPriorityQueue.peek()) {
        		//System.out.println("peek ->" + countsPriorityQueue.peek());
        		
        		if (lastCount-1 > 0) {
        			countsPriorityQueue.add(lastCount-1);
        		}
        		
        		minDeletions++;
        	}
        }
        //System.out.println("After -> " + charCountMap);
        
        return minDeletions;
    }
    
    public static void main(String[] args) {
        
    	MinDeletionsForUniqueString1647 obj = new MinDeletionsForUniqueString1647();
        
        //System.out.println(obj.minDeletions("hello")); // expected output = 2
        //System.out.println(obj.minDeletions("aab")); // // expected output = 0
        //System.out.println(obj.minDeletions("aaabbbccdt")); // // expected output = 2
    	//System.out.println(obj.minDeletions("aaabbbcc")); // // expected output = 2
    	System.out.println(obj.minDeletions("bbcebab")); // // expected output = 2
        //System.out.println(obj.minDeletions("ceabaacb")); // // expected output = 2
        
        return;
        
    }

}
