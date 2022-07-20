package CodingInterviewQuestions;

/***

 Asked in Google, 20July2022

 Consider a self-driving car that can only travel in two directions: back and forth.
 is pre-programmed with a sequence of instructions to travel a certain distance.
 executes each instruction immediately after the previous one finishes, then immediately comes to a complete stop.

 The instruction set:
 * A=Accelerate: move at the current speed and direction for one second, then double the speed. If this is the first instruction in the set, assume an initial speed of 1 meter/second.
 * R=Reverse: instantaneously switch direction, then reset the speed to 1 meter/second (in the new direction).

 The speed is 1 meter/second after a Reverse. The car does not travel during the Reverse phase.

 For example, the given the sequence "AAAAARAAARA",

 A -> 1m, 2m/sec
 A -> 2m, 4m/sec
 A -> 4m, 8m/sec
 A -> 8m, 16m/sec
 A -> 16m, 32m/sec, TOTAL = 31m
 R -> Switch, 1m/sec
 A -> -1m, 2m/sec, TOTAL = 30m
 A -> -2m, 4m/sec, TOTAL = 28m
 A -> -4m, 8m/sec,TOTAL = 24m
 R -> Switch, 1m/sec, TOTAL = 24m
 A -> 1m, 2m/sec, TOTAL = 25m

 String = "AAAAARAAARA"

 int totalDistanceTravelled
 int CurrentSpeed =
 boolean isForward -> true

 ***/

public class SelfDrivingCarDistanceFromInstructions {

    public int calculateDistance(String instructions) {

        int totalDistanceTravelled = 0;
        int currentSpeed = 1;//per second
        boolean isForwardDirection = true;

        for( int i=0; i< instructions.length();i++) {

            if (instructions.charAt(i) == 'A') {

                if (isForwardDirection) {
                    totalDistanceTravelled += currentSpeed;
                }
                else {
                    totalDistanceTravelled -= currentSpeed;
                }

                currentSpeed = currentSpeed*2;
            }
            else {
                if (isForwardDirection) {
                    isForwardDirection = false;
                }
                else {
                    isForwardDirection = true;
                }
                currentSpeed = 1;
            }
        }

        return totalDistanceTravelled;
    }

    public static void main(String[] args) {

        SelfDrivingCarDistanceFromInstructions obj = new SelfDrivingCarDistanceFromInstructions();

        String instructions = "AAAAARAAARA";

        System.out.println(obj.calculateDistance(instructions));
    }
}
