package leetCodeProblems.GreedyAlgo;

/**
 * LeetCode - https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/submissions/
 *
 * TimeComplexity - O(k)
 */

import java.util.Arrays;

public class MaxPointsFromCards1423 {

    public int maxScore(int[] cardPoints, int k) {
        int n = cardPoints.length;

        int[] frontCards = new int[k + 1];
        int[] rearCards = new int[k + 1];

        for (int i = 0; i < k; i++) {
            frontCards[i + 1] = cardPoints[i] + frontCards[i];
            rearCards[i + 1] = cardPoints[n - i - 1] + rearCards[i];
        }

        System.out.println("frontSetOfCards ->" + Arrays.toString(frontCards));
        System.out.println("rearSetOfCards ->" + Arrays.toString(rearCards));

        int answerMaxScore = 0;

        for (int i = 0; i <= k; i++) {
            int currentScore = frontCards[i] + rearCards[k - i];
            answerMaxScore = Math.max(answerMaxScore, currentScore);
            System.out.println("Current Score ->" + currentScore);
            System.out.println("Max Score ->" + answerMaxScore);
        }

        return answerMaxScore;
    }

    public static void main(String[] args) {

        //int[] cardPoints = {1,2,3,4,5,6,1};
        //int k = 3; // answer = 12

        int[] cardPoints = {9,7,7,9,7,7,9};
        int k = 7; // answer = 12

        MaxPointsFromCards1423 obj = new MaxPointsFromCards1423();
        System.out.println(obj.maxScore(cardPoints, k));

    }
}
