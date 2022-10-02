package leetCodeProblems.HashSearch;

/**
 * LeetCode - https://leetcode.com/problems/time-based-key-value-store/
 */

import java.util.Collections;
import java.util.HashMap;
import java.util.PriorityQueue;

public class TimeBasedKeyValueStore981 {

    HashMap<String, String> map;
    PriorityQueue<Integer> queue;

    public TimeBasedKeyValueStore981() {
        map = new HashMap<>();
        queue = new PriorityQueue<>(Collections.reverseOrder());
    }

    public void set(String key, String value, int timestamp) {

        String mapKey = key + "_" + timestamp;
        map.put(mapKey, value);
        queue.add(timestamp);
    }

    public String get(String key, int timestamp) {

        String mapKey = key + "_" + timestamp;

        if (map.containsKey(mapKey)) {
            return map.get(mapKey);
        }

        while (!queue.isEmpty()) {

            int peekStamp = queue.peek();

            if (peekStamp <= timestamp) {
                String mapKey1 = key + "_" + peekStamp;
                return map.get(mapKey);
            }
        }

        return "";
    }
}
