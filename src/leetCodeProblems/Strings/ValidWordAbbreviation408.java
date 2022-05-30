package leetCodeProblems.Strings;

/**
 * LeetCode - https://leetcode.com/problems/valid-word-abbreviation/
 * TimeComplexity - O(n)
 * Space Complexity - O(1)
 *
 */
public class ValidWordAbbreviation408 {

    public boolean validWordAbbreviation(String word, String abbr) {

        String lastAbbr = "";
        int wordPointer = 0;

        for (int i=0; i < abbr.length(); i++) {

            if (Character.isDigit(abbr.charAt(i))) {
                lastAbbr += abbr.charAt(i);
            }
            else {

                if (!lastAbbr.isEmpty()) {

                    if (lastAbbr.charAt(0) == '0') {
                        return false;
                    }

                    wordPointer += Integer.parseInt(lastAbbr);
                    lastAbbr = "";
                }

                if ( wordPointer < word.length()){

                    if(abbr.charAt(i) != word.charAt(wordPointer)) {
                        return false;
                    }

                    wordPointer++;
                }
                else {
                    return false;
                }
            }
        }

        if (!lastAbbr.isEmpty()) {

            if (lastAbbr.charAt(0) == '0') {
                return false;
            }

            wordPointer += Integer.parseInt(lastAbbr);
        }

        if ( wordPointer != word.length()){
            return false;
        }

        return true;
    }

    public static void main(String[] args) {

        ValidWordAbbreviation408 obj = new ValidWordAbbreviation408();

        boolean output = false;

        /*output = obj.validWordAbbreviation("internationalization", "i12iz4n");
        System.out.println(output);

        output = obj.validWordAbbreviation("apple", "a2e");
        System.out.println(output);

        output = obj.validWordAbbreviation("a", "2");
        System.out.println(output);

        output = obj.validWordAbbreviation("internationalization", "i5a11o1");
        System.out.println(output);*/
        output = obj.validWordAbbreviation("a", "01");
        System.out.println(output);
    }
}
