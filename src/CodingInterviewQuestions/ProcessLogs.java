package CodingInterviewQuestions;

/**
 * Asked in Amazon Coding Round Demo, 29thMay2022.
 */

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

public class ProcessLogs {

    public static List<String> processLogs(List<String> logs, int threshold) {
        // Write your code here

        HashMap<String, Integer> map = new HashMap<>();
        List<String> output = new ArrayList<>();

        for(int i=0; i < logs.size(); i++) {

            String log = logs.get(i);

            String[] logsSplit = log.split(" ");

            String senderUserID = logsSplit[0];
            String receiverUserID = logsSplit[1];

            if (map.containsKey(senderUserID)) {
                map.put(senderUserID, map.get(senderUserID)+1);
            }
            else {
                map.put(senderUserID, 1);
            }

            if (!senderUserID.equals(receiverUserID)) {
                if (map.containsKey(receiverUserID)) {
                    map.put(receiverUserID, map.get(receiverUserID)+1);
                }
                else {
                    map.put(receiverUserID, 1);
                }
            }

        }

        //System.out.println(map);

        for (String key: map.keySet()) {

            if (map.get(key) >= threshold) {
                output.add(key);
            }
        }

        return output;
    }

    public static void main(String[] args) {

        List<String> logs = new ArrayList<>();

        logs.add("88 98 100");
        logs.add("98 88 90");
        logs.add("10 10 50");
        logs.add("88 99 100");
        System.out.println(ProcessLogs.processLogs(logs, 3));
    }
}
