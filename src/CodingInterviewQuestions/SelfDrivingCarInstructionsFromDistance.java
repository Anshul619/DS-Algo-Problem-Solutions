package CodingInterviewQuestions;

/**
 * Related LeetCode - https://leetcode.com/problems/race-car/
 */
public class SelfDrivingCarInstructionsFromDistance {

    private String instructions;

    private void calculateInstructions(double distance) {

        if (distance == 1) {
            instructions += "A";
            return;
        }

        if (distance == 0) {
            return;
        }

        double maxDistance = 0;
        int i=0;

        while ( maxDistance <= distance) {
            maxDistance = Math.pow(2, i);
            i++;

            if (maxDistance <= distance) {
                instructions += "A";
            }
        }

        double distanceDiff = maxDistance - distance;

        instructions += "R";

        calculateInstructions(distanceDiff);
    }

    public int racecar(int target) {

        instructions = "";

        calculateInstructions(target);

        return instructions.length();
    }

    public static void main(String[] args) {

        SelfDrivingCarInstructionsFromDistance obj = new SelfDrivingCarInstructionsFromDistance();

        System.out.println(obj.racecar(25));
    }
}
