package leetCodeProblems.HashSearch;

import java.util.*;

/**
 * LeetCode - https://leetcode.com/problems/two-sum/
 * 
 * InterviewBit - https://www.interviewbit.com/problems/2-sum/
 * 
 * @author anshul.agrawal
 *
 */
public class TwoSum1 {
	
	static class Index2Sort implements Comparator<int[]> {

        public int compare(int[] pair1, int[] pair2) {

            return pair1[1] - pair2[1];
        }
    }
	
	public int[] twoSum(int[] nums, int target) {
        
        int[] indexesWithTargetSum = new int[2];
        
        HashMap<Integer, Integer> hashMap = new HashMap<Integer, Integer>();
        
        for(int i=0; i< nums.length; i++) {
            hashMap.put(nums[i], i);
        }
        
        for(int i=0; i < nums.length; i++) {
            
            int neededSum = target - nums[i];
            
            //System.out.println("Item - " + nums[i]+ ", Needed sum - " + neededSum);
            
            if (hashMap.containsKey(neededSum) && hashMap.get(neededSum) != i) {
                indexesWithTargetSum[0] = hashMap.get(neededSum);
                indexesWithTargetSum[1] = i;
                
                return indexesWithTargetSum;
            }
        }
        
        
        return indexesWithTargetSum;
    }
	
	// DO NOT MODIFY THE ARGUMENTS WITH "final" PREFIX. IT IS READ ONLY
    public int[] twoSumWithSavingMatchedPairs(final int[] A, int B) {

        int[] indexesWithTargetSum = new int[2];

        HashMap<Integer, ArrayList<Integer>> hashMap = new HashMap<Integer, ArrayList<Integer>>();
        
        ArrayList<int[]> matchedPairsList = new ArrayList<int[]>();

        for(int i=0; i< A.length; i++) {

            if (hashMap.containsKey(A[i])){
                hashMap.get(A[i]).add(i);
            }
            else {
                ArrayList<Integer> temp = new ArrayList<Integer>();
                temp.add(i);

                hashMap.put(A[i], temp);
            }
            
        }
        
        for(int i=0; i < A.length; i++) {
            
            int neededSum = B - A[i];

            if (hashMap.containsKey(neededSum)) {

                //System.out.println("Lower Index -> " + i + ", Higher Index Size -> " + hashMap.get(neededSum).size() + ", Needed Sum -> " + neededSum);
    
                ArrayList<Integer> temp = hashMap.get(neededSum);

                //System.out.println("neededSum - " + neededSum + ", Size -> " + temp.size());

                for(int j=0; j<temp.size(); j++) {

                    if (i < temp.get(j)) {
                        matchedPairsList.add(new int[]{i+1, temp.get(j)+1});
                    }
                }
            }
        }
        
        if(matchedPairsList.size() == 0) {
            return new int[0];
        }

        Collections.sort(matchedPairsList, new Index2Sort());

        /*for(int i=0; i < matchedPairsList.size(); i++) {
            System.out.println("Matched Pair, index1 -> " + matchedPairsList.get(i)[0] + ", index2 => " + matchedPairsList.get(i)[1]);
        }*/

        if (matchedPairsList.size() > 0) {
            indexesWithTargetSum[0] = matchedPairsList.get(0)[0];
            indexesWithTargetSum[1] = matchedPairsList.get(0)[1];
        }

        /*for(int i=1; i < matchedPairsList.size(); i++) {
            
            if (indexesWithTargetSum[1] == matchedPairsList.get(i)[1]) {
                indexesWithTargetSum[0] = Math.min(indexesWithTargetSum[0], matchedPairsList.get(i)[0]);
            }
            else {
                break;
            }
        }*/
        return indexesWithTargetSum;
    }

}
