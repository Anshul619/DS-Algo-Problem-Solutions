package leetCodeProblems.PriorityQueue;

/**
 * LeetCode - https://leetcode.com/problems/top-k-frequent-elements/submissions/
 *
 * TimeComplexity - O(n) + O(klogk), where k is the count of unique frequencies in the array
 * SpaceComplexity - O(n)
 */

import java.util.Arrays;
import java.util.Comparator;
import java.util.HashMap;
import java.util.PriorityQueue;

public class TopKFrequentElements347 {

    static class ElementFrequency {
        int element;
        int count;

        ElementFrequency(int element, int count) {
            this.element = element;
            this.count = count;

        }
    }

    static class ElementComparator implements Comparator<ElementFrequency> {

        public int compare(ElementFrequency elem1, ElementFrequency elem2) {
            return elem2.count - elem1.count;
        }
    }

    public int[] topKFrequent(int[] nums, int k) {

        HashMap<Integer, Integer> frequencyMap = new HashMap();
        PriorityQueue<ElementFrequency> frequencyPQ = new PriorityQueue<>(new ElementComparator());

        int[] output = new int[k];

        for (int num: nums) {
            frequencyMap.put(num,
                    frequencyMap.getOrDefault(num, 0) + 1);
        }

        for (int element: frequencyMap.keySet()) {
            ElementFrequency obj = new ElementFrequency(element, frequencyMap.get(element));
            frequencyPQ.add(obj);
        }

        for (int i=0; i <k; i++) {
            output[i] = frequencyPQ.poll().element;
        }

        return output;
    }

    public static void main(String[] args) {

        int[] input = {1,1,1,2,2,3};

        TopKFrequentElements347 obj = new TopKFrequentElements347();

        int[] output = obj.topKFrequent(input, 2);

        System.out.print(Arrays.toString(output));

    }
}
