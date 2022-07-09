package leetCodeProblems.HashSearch;

/**
 * LeetCode - https://leetcode.com/problems/reconstruct-original-digits-from-english/submissions/
 *
 * TimeComplexity - O(n)
 * SpaceComplexity - O(n)
 */

import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;

public class ReconstructOriginalDigitsFromEnglish423 {

    HashMap<Character, Integer> frequencyMap;

    Character[] uniqueCharacters = {'z', 'w', 'u', 'x', 'g', 'o', 'h', 'f', 'v', 'n'};
    int[] uniqueCharNumber = {0, 2, 4, 6, 8, 1, 3, 5, 7, 9};

    private boolean isEnglishStringExistsInMap(String s) {

        for (int i=0; i < s.length(); i++) {
            char currentChar = s.charAt(i);

            if (!frequencyMap.containsKey(currentChar)) {
                return false;
            }
        }

        for (int i=0; i < s.length(); i++) {

            char currentChar = s.charAt(i);

            int count = frequencyMap.get(currentChar);
            count--;
            frequencyMap.put(currentChar, count);

            if (count == 0) {
                frequencyMap.remove(currentChar);
            }
        }

        return true;
    }

    private String getStringWithDigit(int i) {

        switch (i) {
            case 0:
                return "zero";
            case 1:
                return "one";
            case 2:
                return "two";
            case 3:
                return "three";
            case 4:
                return "four";
            case 5:
                return "five";
            case 6:
                return "six";
            case 7:
                return "seven";
            case 8:
                return "eight";
            case 9:
                return "nine";
        }

        return "";
    }

    private void generateFrequencyMap(String s) {
        for (int i=0; i < s.length(); i++) {

            char currentChar = s.charAt(i);

            if (frequencyMap.containsKey(currentChar)) {
                int count = frequencyMap.get(currentChar);
                count++;
                frequencyMap.put(currentChar, count);
            }
            else{
                frequencyMap.put(currentChar, 1);
            }
        }
    }
    public String originalDigits(String s) {

        frequencyMap = new HashMap<>();

        String output = "";
        List<Integer> outputList = new ArrayList<>();

        generateFrequencyMap(s);

        for(int i=0; i < uniqueCharacters.length; i++) {

            if (frequencyMap.containsKey(uniqueCharacters[i])) {

                int number = uniqueCharNumber[i];
                int frequency = frequencyMap.get(uniqueCharacters[i]);

                for(int j=0; j < frequency; j++) {
                    if (isEnglishStringExistsInMap(getStringWithDigit(number))) {
                        outputList.add(number);
                    }
                }
            }
        }

        Collections.sort(outputList);

        //System.out.println(outputList);

        for(int i=0; i < outputList.size(); i++) {
            output += String.valueOf(outputList.get(i));
        }

        return output;
    }

    public static void main(String[] args) {

        ReconstructOriginalDigitsFromEnglish423 obj = new ReconstructOriginalDigitsFromEnglish423();

        //System.out.println(obj.originalDigits("fviefuro"));
        //System.out.println(obj.originalDigits("f"));

        System.out.println(obj.originalDigits("owoztneoer"));
        System.out.println(obj.originalDigits("zerozero"));
    }
}
