package CodingInterviewQuestions;

/**
 * Asked in Amazon Coding Round, 29thMay2022.
 */

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class AmazonParcelDelivery {

    /*
     * Complete the 'getMinimumDays' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts INTEGER_ARRAY parcels as parameter.
     */

    public static int getMinimumDays(List<Integer> parcels) {
        // Write your code here

        int output = 0;

        int totalDelivered = 0;
        int lastDelivered = 0;

        int currentCount = 0;

        Collections.sort(parcels);

        //System.out.println(parcels);

        for (int i =0; i < parcels.size(); i++) {

            currentCount = parcels.get(i) - totalDelivered;

            if (parcels.get(i) != 0 && lastDelivered != parcels.get(i)) { // New min
                lastDelivered = parcels.get(i);
                totalDelivered += currentCount;
                output++;
            }
        }

        return output;
    }

    public static void main(String[] args) {
        List<Integer> input = new ArrayList<>();
        input.add(2);
        input.add(3);
        input.add(4);
        input.add(3);
        input.add(3);

        //System.out.println(AmazonParcelDelivery.getMinimumDays(input)); // expected output = 3

        input.clear();
        input.add(4);
        input.add(2);
        input.add(3);
        input.add(4);

        System.out.println(AmazonParcelDelivery.getMinimumDays(input)); // expected output = 3
    }
}
