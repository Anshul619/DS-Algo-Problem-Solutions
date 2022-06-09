package CodingInterviewQuestions;

import java.util.*;

/**
 * Asked in AthenaHealth interview@9June2022
 */

class TwoStrings {

    /*
     * Complete the 'commonSubstring' function below.
     *
     * The function accepts following parameters:
     *  1. STRING_ARRAY a
     *  2. STRING_ARRAY b
     */

    public static void commonSubstring(List<String> a, List<String> b) {

        // Write your code here
        for(int i=0; i < a.size(); i++) {

            String aString = a.get(i);
            String bString = b.get(i);
            int j=0;

            for(j=0; j < aString.length(); j++) {

                if (aString.charAt(j) == bString.charAt(j)) {
                    System.out.println("YES");
                    break;
                }
            }

            if (j == aString.length()) {
                System.out.println("NO");
            }
        }
    }

    public static void main(String[] args) {

        List<String> a = new ArrayList<>();
        a.add("hello");
        a.add("why");

        List<String> b = new ArrayList<>();
        b.add("he");
        b.add("bye");

        TwoStrings.commonSubstring(a, b);
    }

}

