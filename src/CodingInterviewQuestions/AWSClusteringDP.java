package CodingInterviewQuestions;

/**
 * Asked in Amazon Coding Round, 29thMay2022.
 *
 * To-Be-Completed
 */

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class AWSClusteringDP {

    private static int getBootingPowerMax(List<Integer> bootingPower, int startIndex, int endIndex) {

        int maxBooting = 0;

        for (int i=startIndex; i < endIndex; i++) {
            maxBooting = Math.max(maxBooting, bootingPower.get(i));
        }

        return maxBooting;

    }


    public static int findMaximumSustainableClusterSize(List<Integer> processingPower, List<Integer> bootingPower, long powerMax) {
        // Write your code here

        int ansCount = 0;
        int length = processingPower.size();

        int lastSuccessK = 0;

        int[][] dp = new int[length][length];

        for (int i=0; i < length; i++) {
            dp[i][i] = processingPower.get(i);
        }

        //System.out.println("DP ->" + Arrays.toString(dp));
        System.out.println("Length ->" + length);
        for (int k=2; k <= length; k++) {

            for (int i=0; i <= length-k; i++) {

                int minPower = 0;

                minPower = dp[i][i+k-1] + processingPower.get(i+k);

                dp[i][i+k-1] = minPower;

                System.out.println("k ->" + k);

                minPower = minPower*k;
                minPower += getBootingPowerMax(bootingPower, i, i+k);
                System.out.println("minPower ->" + minPower);

                if (minPower <= powerMax) {
                    ansCount++;
                    lastSuccessK = k;
                    break;
                }
            }

            //If it can't happen for k, it won't happen for greater than k.
            if (lastSuccessK != k) {
                break;
            }

        }

        return ansCount;

    }

    public static void main(String[] args) {

        List<Integer> bootingPower = new ArrayList<>();
        bootingPower.add(3);
        bootingPower.add(6);
        bootingPower.add(1);
        bootingPower.add(3);
        bootingPower.add(4);

        List<Integer> processingPower = new ArrayList<>();
        processingPower.add(2);
        processingPower.add(1);
        processingPower.add(3);
        processingPower.add(4);
        processingPower.add(5);

        int max = 25;
        int out = 0;

        //int out = AWSClustering.findMaximumSustainableClusterSize(bootingPower, processingPower, max); //expected O/p = 2

        bootingPower.clear();
        bootingPower.add(4);
        bootingPower.add(1);
        bootingPower.add(4);
        bootingPower.add(5);
        bootingPower.add(3);

        processingPower.clear();
        processingPower.add(8);
        processingPower.add(8);
        processingPower.add(10);
        processingPower.add(9);
        processingPower.add(12);

        max = 41;

        out = AWSClusteringDP.findMaximumSustainableClusterSize(processingPower, bootingPower, max); //expected O/p = 3

        System.out.println(out);
    }
}
