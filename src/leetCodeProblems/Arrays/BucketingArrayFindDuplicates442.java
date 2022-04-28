package leetCodeProblems.Arrays;

public class BucketingArrayFindDuplicates442 {
	
	// DO NOT MODIFY THE ARGUMENTS WITH "final" PREFIX. IT IS READ ONLY
    public int findDuplicatesHashSet(final int[] A) {

        HashSet<Integer> hashSet = new HashSet<Integer>();

        for(int i=0; i<A.length; i++) {
            
            if (hashSet.contains(A[i])) {
                return A[i];
            }

            hashSet.add(A[i]);
        }

        return -1;
    }
	
	public List<Integer> findDuplicatesConstantSpace(int[] nums) {
        
        List<Integer> output = new ArrayList<Integer>();
        
        for(int i=0; i<nums.length; i++) {

            // This is important, to get original value at nums[i]. 
            // Because it might be changed due to array operations
            int hashMapIndexInArray = nums[i] % nums.length;
            
            if (hashMapIndexInArray >= 0) {
                //System.out.println (hashMapIndexInArray);
                nums[hashMapIndexInArray] = nums[hashMapIndexInArray] + nums.length;
            }
        }
        
        for(int i=0; i<nums.length; i++) {
            
            System.out.println(nums[i]);
            if (nums[i] > (nums.length*2)) {
                
                if (i == 0) {
                    output.add(nums.length);
                }
                else {
                    output.add(i);
                }
                
            }
        }
        
        return output;
    }

}
