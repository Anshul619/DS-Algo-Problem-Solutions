package leetCodeProblems.HashSearch;

public class FourSum18 {
	
    List<List<Integer>> output = new ArrayList<List<Integer>>(); 
    
    public boolean isElementExistInOutputList(List<Integer> indexesWithTargetSum) {
        
        for(int i=0; i < output.size(); i++) {
            
            if (indexesWithTargetSum.equals(output.get(i))) {
                return true;
            }
        }
        
        return false;
    }
    
    public List<List<Integer>> fourSum(int[] nums, int target) {
        
        HashMap<Integer, Integer> hashMap = new HashMap<Integer, Integer>();
        
        Arrays.sort(nums);
            
        for(int i=0; i< nums.length; i++) {
            hashMap.put(nums[i], i);
        }
        
        for(int i=0; i < nums.length; i++) {
            
            for(int j=i+1; j < nums.length; j++) {
                
                for(int k=j+1; k < nums.length; k++) {
                    
                    int neededSum = target - nums[i] - nums[j] - nums[k];
                    
                    if (hashMap.containsKey(neededSum) && 
                        hashMap.get(neededSum) > i && 
                        hashMap.get(neededSum) > j && 
                        hashMap.get(neededSum) > k) {
                        
                        ArrayList<Integer> indexesWithTargetSum = new ArrayList<Integer>();
                        indexesWithTargetSum.add(nums[i]);
                        indexesWithTargetSum.add(nums[j]);
                        indexesWithTargetSum.add(nums[k]);
                        indexesWithTargetSum.add(neededSum);
                        
                        if (!isElementExistInOutputList(indexesWithTargetSum)) {
                            output.add(indexesWithTargetSum);
                        }
                    }
                }
            }
        }
        
        return output;
        
    }

}
