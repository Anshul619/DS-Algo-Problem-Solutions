package leetCodeProblems.PrefixSum;

import java.util.HashMap;

public class SubarraySumK560 {

    public int subarraySum(int[] nums, int k) {

        HashMap<Integer, Integer> map = new HashMap<>();
        map.put(0,1);

        int currentSum = 0;
        int ansCount = 0;

        System.out.println(map);

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

            System.out.println(map);
        }

        return ansCount;
    }

    public static void main(String[] args) {

        SubarraySumK560 obj = new SubarraySumK560();

        int[] input = {10, 2, -2, -20, 10};
        int k = -10;

        //int[] input = {1,1,1};
        //int k= 2;
        System.out.println(obj.subarraySum(input, k));

    }
}
