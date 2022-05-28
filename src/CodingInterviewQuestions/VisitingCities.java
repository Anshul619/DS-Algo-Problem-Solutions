package CodingInterviewQuestions;

/**
 * Asked in Twillo Coding Round, 28thMay2022
 */

import java.util.ArrayList;
import java.util.List;

public class VisitingCities {
    public static List<Long> minimumCost(List<Integer> red, List<Integer> blue, int blueCost) {

        List<Long> output = new ArrayList<>();

        //output.add(Long.)
        output.add(Long.valueOf(0));

        String lastLine = "RED";
        int lastMinCost = 0;

        int localBlueCost = 0;
        int localRedCost = 0;

        for (int i=0; i < blue.size(); i++) {

            localRedCost = lastMinCost + red.get(i);
            localBlueCost = lastMinCost + blue.get(i);

            if (lastLine.equals("RED")) {
                localBlueCost += blueCost;
            }

            if (localRedCost < localBlueCost) {
                lastMinCost = localRedCost;
                output.add(Long.valueOf(localRedCost));
            }
            else {
                lastLine = "BLUE";
                lastMinCost = localBlueCost;
                output.add(Long.valueOf(localBlueCost));
            }
        }

        return output;
    }

    public static void main(String[] args) {

        VisitingCities obj = new VisitingCities();

        List<Integer> red = new ArrayList<>();
        red.add(2);
        red.add(3);
        red.add(4);

        List<Integer> blue = new ArrayList<>();
        blue.add(3);
        blue.add(1);
        blue.add(1);

        int blueCost = 2;

        System.out.println(obj.minimumCost(red, blue, blueCost));
    }
}
