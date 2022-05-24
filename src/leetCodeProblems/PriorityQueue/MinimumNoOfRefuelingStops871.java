package leetCodeProblems.PriorityQueue;

/**
 * https://leetcode.com/problems/minimum-number-of-refueling-stops/submissions/
 */

import java.util.Comparator;
import java.util.HashSet;
import java.util.PriorityQueue;

public class MinimumNoOfRefuelingStops871 {

    static class GasStation {
        int position;
        int fuel;

        GasStation(int position, int fuel) {
            this.position = position;
            this.fuel = fuel;
        }
    }

    static class PriorityQueueCompare implements Comparator<GasStation> {

        public int compare(GasStation station1, GasStation station2) {

            return station2.fuel - station1.fuel;
        }
    }

    private int isMaxAlreadyTraversed(HashSet<Integer> alreadyTraversedStations, int currentStationPosition) {

        for (Integer station: alreadyTraversedStations) {
            if (station >= currentStationPosition) {
                return station;
            }
        }

        return -1;
    }

    private void rebuiltQueue(PriorityQueue<GasStation> gasStationQueue, PriorityQueue<GasStation> tempQueue) {

        for (GasStation gasStation: gasStationQueue) {
            tempQueue.add(gasStation);
        }
    }

    public int minRefuelStops(int target, int startFuel, int[][] stations) {

        if (startFuel >= target) {
            return 0;
        }

        if (stations.length == 0) {
            return -1;
        }

        PriorityQueue<GasStation> gasStationQueue =
                new PriorityQueue<GasStation>(stations.length, new PriorityQueueCompare());


        for(int i=0; i < stations.length; i++) {

            GasStation station = new GasStation(stations[i][0], stations[i][1]);

            gasStationQueue.add(station);
        }

        /*System.out.println(gasStationQueue.remove().fuel);
        System.out.println(gasStationQueue.remove().fuel);
        System.out.println(gasStationQueue.remove().fuel);
        System.out.println(gasStationQueue.remove().fuel);*/

        int currentFuel = startFuel;
        int currentTarget = target;
        int currentPosition = 0;

        HashSet<Integer> alreadyTraversedStations = new HashSet<Integer>();

        int ans = 0;
        int stationsTraversed = 0;


        while (currentTarget > currentFuel) {

            if (gasStationQueue.size() == 0) {
                return -1;
            }

            PriorityQueue<GasStation> tempQueue =
                    new PriorityQueue<GasStation>(gasStationQueue.size(), new PriorityQueueCompare());

            rebuiltQueue(gasStationQueue, tempQueue);

            while (tempQueue.size() > 0) {

                GasStation maxStation = tempQueue.remove();

                int neededFuel = maxStation.position - currentPosition;

                System.out.println("maxStation.position =>" + maxStation.position);
                System.out.println("neededFuel => " + neededFuel);
                if (neededFuel <= currentFuel) {

                    int maxAlreadyTraversed = isMaxAlreadyTraversed(alreadyTraversedStations, maxStation.position);

                    if ( maxAlreadyTraversed!= -1) {
                        currentFuel = currentFuel + maxStation.fuel;
                        currentTarget = target - maxAlreadyTraversed;
                        currentPosition = maxAlreadyTraversed;
                    }
                    else {
                        currentFuel = currentFuel - neededFuel + maxStation.fuel;
                        currentTarget = target - maxStation.position;
                        currentPosition = maxStation.position;
                    }

                    System.out.println("currentFuel => " + currentFuel);
                    System.out.println("currentTarget => " + currentTarget);
                    System.out.println("currentPosition => " + currentPosition);
                    gasStationQueue.remove(maxStation);

                    alreadyTraversedStations.add(maxStation.position);

                    ans++;
                    break;
                }

                else {
                    stationsTraversed++;
                }
            }

            if (stationsTraversed == stations.length) {
                return -1;
            }
        }

        if (ans == 0) {
            ans = -1;
        }

        return ans;
    }

    public static void main(String[] args) {

        MinimumNoOfRefuelingStops871 obj = new MinimumNoOfRefuelingStops871();

        //int[][] stations = {{10,60}, {20,30}, {30,30}, {60,40}};
        //int target = 100;
        //int startFuel = 10; // Expected Output = 2

        //int[][] stations = {{2,4}, {3,2}, {5,5}, {16,10}};
        //int target = 25;
        //int startFuel = 10; // Expected Output = 3

        /*int[][] stations = {{10,100}};
        int target = 100;
        int startFuel = 1;*/

        //int[][] stations = new int[0][0];
        //int target = 1;
        //int startFuel = 1;

        //int[][] stations = new int[0][0];
        //int target = 100;
        //int startFuel = 1;


        //int[][] stations = {{25,25}, {50,25}, {75,25}};
        //int target = 100;
        //int startFuel = 25;

        int[][] stations = {{25,30}};
        int target = 100;
        int startFuel = 50;

        System.out.println(obj.minRefuelStops(target, startFuel, stations));

    }
}
