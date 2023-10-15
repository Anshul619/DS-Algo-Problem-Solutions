package leetCodeProblems.HashSearch;

/**
 * LeetCode - https://leetcode.com/problems/3sum/
 */
import java.util.*;

public class ThreeSum19UsingHashSearch {
	
    List<List<Integer>> output = new ArrayList<List<Integer>>(); 
    
    private boolean isElementExistInOutputList(List<Integer> indexesWithTargetSum) {
        
        for(int i=0; i < output.size(); i++) {
            
            if (indexesWithTargetSum.equals(output.get(i))) {
                return true;
            }
        }
        
        return false;
    }
    
    public List<List<Integer>> threeSum(int[] nums, int target) {
        
        HashMap<Integer, Integer> hashMap = new HashMap<Integer, Integer>();
        
        Arrays.sort(nums);
            
        for(int i=0; i< nums.length; i++) {
            hashMap.put(nums[i], i);
        }
        
        for(int i=0; i < nums.length; i++) {
            
            for(int j=i+1; j < nums.length; j++) {
                
            	int neededSum = target - nums[i] - nums[j];
                
                if (hashMap.containsKey(neededSum) && 
                    hashMap.get(neededSum) > i && 
                    hashMap.get(neededSum) > j) {
                    
                    ArrayList<Integer> indexesWithTargetSum = new ArrayList<Integer>();
                    indexesWithTargetSum.add(nums[i]);
                    indexesWithTargetSum.add(nums[j]);
                    indexesWithTargetSum.add(neededSum);
                    
                    if (!isElementExistInOutputList(indexesWithTargetSum)) {
                        output.add(indexesWithTargetSum);
                    }
                }
            }
        }
        
        return output;
        
    }
    
    public static void main(String[] args) {
    	int[] inputArray = {-1, 0, 1, 2, -1, -4}; 
    	
    	ThreeSum19UsingHashSearch obj = new ThreeSum19UsingHashSearch();
    	
    	System.out.println(obj.threeSum(inputArray, 0));
    }

}
