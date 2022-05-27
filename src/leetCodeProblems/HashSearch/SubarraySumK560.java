package leetCodeProblems.HashSearch;

import java.util.HashMap;

public class SubarraySumK560 {

    public int subarraySum(int[] nums, int k) {

        HashMap<Integer, Integer> map = new HashMap<>();
        map.put(0,1);

        int currentSum = 0;
        int ansCount = 0;

        for (int i=0; i<nums.length; i++) {

            currentSum += nums[i];

            if (map.containsKey(currentSum - k)) {
                ansCount += map.get(currentSum - k);
            }

            if (map.containsKey(currentSum)) {
                map.put(currentSum, map.get(currentSum)+1);
            }
            else {
                map.put(currentSum, 1);
            }
        }

        return ansCount;
    }

    public static void main(String[] args) {

        SubarraySumK560 obj = new SubarraySumK560();

        int[] input = {10, 2, -2, -20, 10};
        int k = -10;

        System.out.println(obj.subarraySum(input, k));

    }
}
