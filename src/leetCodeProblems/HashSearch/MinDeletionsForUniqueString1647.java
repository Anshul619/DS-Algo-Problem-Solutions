package leetCodeProblems.HashSearch;

/**
 * LeetCode - https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/submissions/
 */

import java.util.*;

public class MinDeletionsForUniqueString1647 {
	
	
	/**
	 * getMaxLowerUnseenCount
	 * 
	 * @param currentCount
	 * @param uniqueCounts
	 * @return
	 */
    public int getMaxLowerUnseenCount(int currentCount,  HashSet<Integer> uniqueCounts) {
        
    	
        while(currentCount > 0) {
            currentCount = currentCount - 1;
            
            if (!uniqueCounts.contains(currentCount)) {
                return currentCount;
            }

        }
        
        return 0;
        
    }
    
    /**
     * minDeletions
     * 
     * @param s
     * @return
     */
    public int minDeletions(String s) {
        
    	HashMap<String, Integer> charCountMap = new HashMap<String, Integer>();
        HashSet<Integer> uniqueCounts = new HashSet<Integer>();
        HashSet<Integer> seenCounts = new HashSet<Integer>();
        
        int minDeletions = 0;
        
        for (int i=0; i < s.length(); i++) {
            
            String characterStr = String.valueOf(s.charAt(i));
            
            if (charCountMap.containsKey(characterStr)) {
				
				//System.out.println("Key is contained - " + votes[i]);
				int currentCount = charCountMap.get(characterStr);
				currentCount = currentCount+1;
				
				charCountMap.put(characterStr, currentCount);
			}
			else {
				charCountMap.put(characterStr, 1);
			}
        }
        
        for (Integer count: charCountMap.values()) {
            uniqueCounts.add(count);
        }
        
        //System.out.println("Before -> " + charCountMap);
        //System.out.println("Unique Counts ->" + uniqueCounts);
        
        for (String currentChar: charCountMap.keySet()) {
            
            int currentCharCount = charCountMap.get(currentChar);
            int nextCharUniqueCount = 0;
            
            // This is a duplicate count, hence we need to decrease it.
            if (seenCounts.contains(currentCharCount)) {
                nextCharUniqueCount = getMaxLowerUnseenCount(currentCharCount, uniqueCounts);
                  
                minDeletions += currentCharCount-nextCharUniqueCount;
                
                seenCounts.add(nextCharUniqueCount);
                uniqueCounts.add(nextCharUniqueCount);
                charCountMap.put(currentChar, nextCharUniqueCount);
                
            }
            else {
            	seenCounts.add(currentCharCount);
            }
        }
        
        //System.out.println("After -> " + charCountMap);
        
        return minDeletions;
    }
    
    public static void main(String[] args) {
        
    	MinDeletionsForUniqueString1647 obj = new MinDeletionsForUniqueString1647();
        
        System.out.println(obj.minDeletions("hello")); // expected output = 2
        System.out.println(obj.minDeletions("aab")); // // expected output = 0
        System.out.println(obj.minDeletions("aaabbbcc")); // // expected output = 2
        System.out.println(obj.minDeletions("ceabaacb")); // // expected output = 2
        
        return;
        
    }

}
