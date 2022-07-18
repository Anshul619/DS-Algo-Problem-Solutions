package leetCodeProblems.BruteForce;

/**
 * LeetCode - https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
 *
 * TimeComplexity - O(N) where N is string length
 * SpaceComplexity - O(N-D) where D is total length of all duplicates
 */
public class RemoveAllAdjacentDuplicatesInString1047 {

    public String removeDuplicates(String s) {

        StringBuilder sb = new StringBuilder();
        int sbLength = 0;

        for (char character: s.toCharArray()) {

            if (sbLength != 0 && character == sb.charAt(sbLength-1)) {
                sb.deleteCharAt(sbLength-1);
                sbLength--;
            }
            else {
                sb.append(character);
                sbLength++;
            }
        }

        return sb.toString();
    }

    /**
     * BruteForce technique - TimeLimitExceeding on LeetCode
     *
     * @param s
     * @return
     */
    public String removeDuplicatesBruteForce(String s) {

        String output = "";
        boolean iterateNextDuplicate = true;

        while (iterateNextDuplicate) {

            boolean duplicateFound = false;
            output = "";

            //System.out.println(s);

            for (int i=0; i < s.length(); i++) {
                if (i < s.length()-1 && s.charAt(i) == s.charAt(i+1)) {
                    i += 1;
                    duplicateFound = true;
                }
                else {
                    output += s.charAt(i);
                }
            }

            s = output;

            if (duplicateFound) {
                iterateNextDuplicate = true;
            }
            else {
                iterateNextDuplicate = false;
            }
        }

        return output;
    }

    public static void main(String[] args) {
        String input = "abbaca";

        RemoveAllAdjacentDuplicatesInString1047 obj = new RemoveAllAdjacentDuplicatesInString1047();
        String output = obj.removeDuplicates(input);

        System.out.println(output);
    }
}
