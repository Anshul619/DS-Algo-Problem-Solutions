package leetCodeProblems.GreedyAlgo;

/**
 * LeetCode - https://leetcode.com/problems/buildings-with-an-ocean-view/
 * Time Complexity - O(n)
 * Space Complexity - O(n)
 */

import java.util.ArrayList;
import java.util.Arrays;

public class BuildingsWithOceanView1762 {
    public int[] findBuildings(int[] heights) {

        int maxHeight = 0;

        ArrayList<Integer> outputList = new ArrayList<>();

        for (int i = heights.length-1; i >=0; i--) {

            if (heights[i] > maxHeight) {
                outputList.add(i);
                maxHeight = heights[i];
            }
        }

        //System.out.println(queue.poll());
        //System.out.println(queue.poll());
        //System.out.println(queue.poll());
        //System.out.println(queue.toString());

        int[] output = new int[outputList.size()];

        for(int i=0; i < outputList.size(); i++){
            output[outputList.size()-i-1] = outputList.get(i);
        }
        return output;
    }

    public static void main(String[] args) {

        BuildingsWithOceanView1762 obj = new BuildingsWithOceanView1762();

        int[] input = {4,2,3,1};

        int[] output = obj.findBuildings(input);

        System.out.println(Arrays.toString(output));
    }


}
