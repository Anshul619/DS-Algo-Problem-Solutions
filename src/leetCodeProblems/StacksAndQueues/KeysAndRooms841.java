package leetCodeProblems.StacksAndQueues;

import java.util.*;

/**
 * LeetCode - https://leetcode.com/problems/keys-and-rooms/
 *
 * Traversal Type
 * - If we use Queue, it would be BFS.
 * - If we use Stack, it would be DFS.
 *
 * TimeComplexity - O(N+E) - N is the number of rooms, E is total number of keys
 * SpaceComplexity - O(N) - Extra storage for Queue & VisitedRooms array
 */
public class KeysAndRooms841 {

    public static boolean canVisitAllRooms(List<List<Integer>> rooms) {

        int roomsCount = rooms.size();

        //Stack<Integer> unusedKeys = new Stack();
        Queue<Integer> unusedKeys = new LinkedList();
        boolean[] visitedRooms = new boolean[roomsCount];

        unusedKeys.add(0);

        while (!unusedKeys.isEmpty()) {

            //int roomKey = unusedKeys.pop(); // Stack operation
            int roomKey = unusedKeys.remove();

            if (!visitedRooms[roomKey]) {

                visitedRooms[roomKey] = true;
                List<Integer> newKeys = rooms.get(roomKey);

                for(int i=0; i<newKeys.size(); i++) {

                    if (!visitedRooms[newKeys.get(i)]) {
                        unusedKeys.add(newKeys.get(i));
                    }
                }
            }
        }

        for(int i=0; i < visitedRooms.length; i++) {
            if (!visitedRooms[i]) {
                return false;
            }
        }

        return true;
    }

    public static void main(String[] args) {

        ArrayList<Integer> firstList = new ArrayList<>();
        firstList.add(1);

        ArrayList<Integer> secondList = new ArrayList<>();
        secondList.add(2);

        ArrayList<Integer> thirdList = new ArrayList<>();
        thirdList.add(3);

        ArrayList<Integer> forthList = new ArrayList<>();

        List<List<Integer>> allRoomsList = new ArrayList<>();
        allRoomsList.add(firstList);
        allRoomsList.add(secondList);
        allRoomsList.add(thirdList);
        allRoomsList.add(forthList);

        /*ArrayList<Integer> firstList = new ArrayList<>();
        firstList.add(1);
        firstList.add(3);

        ArrayList<Integer> secondList = new ArrayList<>();
        secondList.add(3);
        secondList.add(0);
        secondList.add(1);

        ArrayList<Integer> thirdList = new ArrayList<>();
        thirdList.add(2);

        ArrayList<Integer> forthList = new ArrayList<>();
        forthList.add(0);

        ArrayList<ArrayList<Integer>> allRoomsList = new ArrayList<>();
        allRoomsList.add(firstList);
        allRoomsList.add(secondList);
        allRoomsList.add(thirdList);
        allRoomsList.add(forthList);*/

        boolean output = KeysAndRooms841.canVisitAllRooms(allRoomsList);

        System.out.println(output);

    }
}
