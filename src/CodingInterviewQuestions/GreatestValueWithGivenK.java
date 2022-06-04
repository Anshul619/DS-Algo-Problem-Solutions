package CodingInterviewQuestions;

/**
 * Asked in Vonage Coding Round, 4thJune2022.
 */
public class GreatestValueWithGivenK {

    private int currentK = 0;

    private int calculateFinalDigit(int currentDigit) {

        //System.out.println("Before ->" + currentDigit);
        if (currentK == 0) {
            return currentDigit;
        }

        int currentDigitK = currentK - currentDigit;
        int currentDigitNeededDiff = 9 - currentDigit;

        //System.out.println("currentDigitNeededDiff ->" + currentDigitNeededDiff);
        if (currentDigitNeededDiff <= currentDigitK) {
            currentDigit = 9;
            currentK -= currentDigitNeededDiff;
        }
        else {
            currentDigit += currentK;
            currentK = 0;
        }

        //System.out.println("After ->" + currentDigit);

        return currentDigit;
    }
    public int solution(int N, int K) {
        // write your code in Java SE 8

        currentK = K;
        int finalAnswer = 0;

        int firstDigit = N/100;
        finalAnswer += calculateFinalDigit(firstDigit)*100;

        //System.out.print();
        int secondDigit = (N - firstDigit*100)/10;
        finalAnswer += calculateFinalDigit(secondDigit)*10;

        int thirdDigit = (N - firstDigit*100);
        thirdDigit = (thirdDigit - (thirdDigit/10)*10);
        //System.out.println("thirdDigit ->" + thirdDigit);
        finalAnswer += calculateFinalDigit(thirdDigit)*1;

        return finalAnswer;

    }

    public static void main(String[] args) {

        //int N = 512;
        //int K = 10; // 972

        //int N = 191;
        //int K = 4; //591

        int N = 285;
        int K = 20; //999

        GreatestValueWithGivenK obj = new GreatestValueWithGivenK();
        System.out.print(obj.solution(N, K));

    }
}
