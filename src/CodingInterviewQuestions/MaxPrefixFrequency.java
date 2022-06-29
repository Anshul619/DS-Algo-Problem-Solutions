package CodingInterviewQuestions;

/**
 * Asked in Amex, 29thJune2022
 *
 * To-Be-Completed
 */

class MaxPrefixFrequency {
    int solution(int X, int Y, int[] A) {
        int N = A.length;
        int result = -1;
        int nX = 0;
        int nY = 0;
        for (int i = 0; i < N; i++) {
            if (i==0 && (A[i] != X && A[i] != Y))
                result = 0;
            if (A[i] == X)
                nX += 1;
            else if (A[i] == Y)
                nY += 1;
            if (nX == nY)
                result = i;
        }
        return result;
    }

    public static void main(String[] args) {

        MaxPrefixFrequency test = new MaxPrefixFrequency();

        int[] input = {1,2};

        System.out.println(test.solution(1,2, input));
    }
}

