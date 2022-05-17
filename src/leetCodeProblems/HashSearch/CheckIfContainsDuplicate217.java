package leetCodeProblems.HashSearch;

import java.util.HashSet;

public class CheckIfContainsDuplicate217 {
    public boolean containsDuplicate(int[] nums) {

        HashSet<Integer> numbersHashSet = new HashSet<>();

        for(int i=0; i < nums.length; i++) {

            if (numbersHashSet.contains(nums[i])) {
                return true;
            }
            else {
                numbersHashSet.add(nums[i]);
            }
        }
        return false;
    }

    public static void main(String[] args) {

        //int[] inputArray = {1, 2, 3, 1};

        int[] inputArray = {1,2,3,4};

        CheckIfContainsDuplicate217 obj = new CheckIfContainsDuplicate217();

        System.out.println(obj.containsDuplicate(inputArray));

    }
}
