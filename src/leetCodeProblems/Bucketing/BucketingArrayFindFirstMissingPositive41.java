package leetCodeProblems.Bucketing;

public class BucketingArrayFindFirstMissingPositive41 {
	
	public int[] shiftPositveNumbersToNewArray(int[] nums, int negativeNumbersArrSize) {
        int[] positiveNumbersArr = new int[nums.length - negativeNumbersArrSize];
        
        int positiveNumbersStartPoint = negativeNumbersArrSize;
        
        for(int i = 0; i < positiveNumbersArr.length; i++) {
            positiveNumbersArr[i] = nums[positiveNumbersStartPoint+i];
        }
        
        return positiveNumbersArr;
    }
    
    // This would be done using 2-pointer approach.
    // This would return index where negative numbers array is ending.
    public int segregateNegativePositiveNumbers(int[] nums) {
        
        // Note - negativeNumbersStartPoint is equal to 0. And it won't be changed.
        int negativeNumbersArrSize = 0; 
            
        for (int i=0; i<nums.length;i++) {
            
            if (nums[i] <= 0) {
                int temp = nums[negativeNumbersArrSize];
                nums[negativeNumbersArrSize] = nums[i];
                nums[i] = temp;
                
                negativeNumbersArrSize++;
            }
        }
        
        return negativeNumbersArrSize;
    }
    
    public int findMissingPositiveInPositiveArray(int[] positiveNumbersArr){
        
        for(int i=0; i < positiveNumbersArr.length; i++) {
            
            // This is needed, because the value might be negative because of earlier swapping. Hence consider its positive value. 
            int indexValuePositive = Math.abs(positiveNumbersArr[i]);
            
            int indexToNegate = indexValuePositive-1;
            
            if (indexToNegate >=0 && 
                indexToNegate < positiveNumbersArr.length && 
                positiveNumbersArr[indexToNegate] > 0) {
                
                //System.out.println("Current index -> " + i +", Current index value -> " + positiveNumbersArr[i]+ ", Negate at index -> " + indexToNegate);
                positiveNumbersArr[indexToNegate] = -positiveNumbersArr[indexToNegate];
            }
        }
        
        //System.out.println("positiveNumbersArr, after negation- " + Arrays.toString(positiveNumbersArr));
        
        for(int i=0; i < positiveNumbersArr.length; i++) {
            if (positiveNumbersArr[i] > 0) {
                return i+1; // Since index is ZERO based, hence increment it by 1
            }
        }
        
        return positiveNumbersArr.length+1; // Since index is ZERO based, hence increment it by 1
    } 
    
    public int firstMissingPositive(int[] nums) {
        
        int negativeNumbersArrSize = segregateNegativePositiveNumbers(nums);
    
        int[] positiveNumbersArr = shiftPositveNumbersToNewArray(nums, negativeNumbersArrSize);
        
        //System.out.println("positiveNumbersArr, before negative - " + Arrays.toString(positiveNumbersArr));
        int firstMissingPositive = findMissingPositiveInPositiveArray(positiveNumbersArr);
        
        return firstMissingPositive;
        
    }
    
    // Driver code
    public static void main(String[] args)
    {
        int arr[] = { 0, 10, 2, -10, -20 };
        
        BucketingArrayFindFirstMissingPositive41 obj = new BucketingArrayFindFirstMissingPositive41();
        
        int missing = obj.firstMissingPositive(arr);
        System.out.println("The smallest positive missing number is " + missing);
    }

}
