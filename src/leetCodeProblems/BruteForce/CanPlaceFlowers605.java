package leetCodeProblems.BruteForce;

/**
 * LeetCode - https://leetcode.com/problems/can-place-flowers/submissions/
 *
 * TimeComplexity - O(n)
 * SpaceComplexity - O(1)
 */
public class CanPlaceFlowers605 {

    public boolean canPlaceFlowers(int[] flowerbed, int n) {

        int availablePositions = 0;

        for (int i=0; i < flowerbed.length; i++) {

            if (flowerbed[i] == 0) {

                boolean emptyLeft = false;
                boolean emptyRight = false;

                if (i == 0) {
                    emptyLeft = true;
                } else if (flowerbed[i-1] != 1) {
                    emptyLeft = true;
                }

                if (i == flowerbed.length-1) {
                    emptyRight = true;
                }
                else if (flowerbed[i+1] != 1) {
                    emptyRight = true;
                }

                //System.out.println("emptyRight ->" + emptyRight);

                if (emptyLeft && emptyRight) {
                    availablePositions++;
                    i++;
                }
            }
        }

        if (availablePositions >= n) {
            return true;
        }

        return false;
    }

    public static void main(String[] args) {

        //int[] flowerbed = {1,0,0,0,1};
        //int n = 1;

        //int[] flowerbed = {1,0,0,0,0,1};
        //int n = 2;

        //int[] flowerbed = {1,0,0,0,1,0,0};
        //int n = 2;

        int[] flowerbed = {0,0,1,0,0};
        int n = 1;

        CanPlaceFlowers605 obj = new CanPlaceFlowers605();
        System.out.println(obj.canPlaceFlowers(flowerbed, n));

    }
}
