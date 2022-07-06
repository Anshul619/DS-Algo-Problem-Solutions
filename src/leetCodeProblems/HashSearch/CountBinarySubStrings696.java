package leetCodeProblems.HashSearch;

/**
 * LeetCode - https://leetcode.com/problems/count-binary-substrings/
 *
 * TimeComplexity - O(n2)
 * SpaceComplexity - O(n)
 */

import java.util.HashMap;

public class CountBinarySubStrings696 {

    public int countBinarySubstrings(String s) {
        HashMap<String, Integer> map = new HashMap();

        int leftPointer = 0;
        int answer = 0;

        while(leftPointer < s.length()) {

            int rightPointer = leftPointer;
            int zerosCount = 0;
            int onesCount = 0;
            String subString = "";

            while(rightPointer < s.length()) {

                if (s.charAt(rightPointer) == '0') {
                    zerosCount++;
                }
                else if (s.charAt(rightPointer) == '1') {
                    onesCount++;
                }

                subString += s.charAt(rightPointer);

                if (zerosCount == onesCount) {

                    if (map.containsKey(subString)) {
                        int count = map.get(subString);
                        count++;
                        map.put(subString, count);
                    }
                    else {
                        map.put(subString, 1);
                    }

                    break;
                }

                rightPointer++;
            }

            leftPointer++;
        }

        for(Integer value: map.values()) {
            answer+= value;
        }

        return answer;
    }

    public static void main(String[] args) {

        // String input = "00110011";
        // String input = "10101";
        String input = "110011111011011001111111111111001100101111111110011011111111110101001001111000111111011110111111010111010000111100111001110101100011101110101111111010110010111111010110110111111100111111100001111011010111111111101010100111110001000000100110100110010111111011011111011111101001100111011110111000101110011001101000101110000000110100001110011111101011010111011000001111011011110101011001010111110101111111111111010011010111110111100110101110111111111110011101101101001011011111110111001101111111111000111011000010111111111110111001111101110001011111001111011111111011100011010101110110101110101101110001110110111111011100011111010110011011111100111111011111110110111001110111011111011111110100111110011101011101111001101110010110010101011000101110011011110111111011111111111001010110010011111010101111101010110110111111011111101101100100111100111100011101011110010011001111111110010011110001111101110000100110111101101111001010111000111101000000111111000101100011011011111100011000111010110011111110001";

        CountBinarySubStrings696 obj = new CountBinarySubStrings696();

        System.out.println(obj.countBinarySubstrings(input));
    }
}
