package CodingInterviewQuestions;

/**
 * Asked in Traxn Interview, 27-June-2022
 *
 * Question - Check if A string is substring of B or not.
 *
 * TimeComplexity - O(n)
 * SpaceComplexity - O(1)
 */
public class IsSubString {

    public static boolean matchString(String A, String B) {

        if (A.length() == 0) {
            return false;
        }

        if (B.length() == 0) {
            return false;
        }

        int lastIteratedIndex = 0;

        for (int i = 0; i < B.length(); i++) {

            if (B.charAt(i) == A.charAt(0)) {

                i++;

                lastIteratedIndex = i;

                int j = 0;

                for (j = 1; j < A.length(); j++) {

                    if (i >= B.length()) {
                        return false;
                    }

                    if (B.charAt(i) != A.charAt(j)) {
                        break;
                    }

                    i++;
                }

                if (j == A.length()) {
                    return true;
                } else {
                    i = lastIteratedIndex;
                }
            }
        }

        return false;

    }

    public static void main(String[] args) {

        //String A = "I am";
        //String B = "My Name - I am Anshul Agrawal";

        //String A = "I am not";
        //String B = "My Name - I am Anshul Agrawal";

        String A = "am m1";
        String B = "My Name - I am am m1 Anshul Agrawal";

        System.out.println(IsSubString.matchString(A, B));
    }
}
