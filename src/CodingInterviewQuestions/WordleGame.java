package CodingInterviewQuestions;

import java.util.HashMap;

/**
 * Asked in Google, 30thJune2022 - PreQualifier Round
 */

/*
   Wordle is a game where there's a secret 5-letter word.
   - The player has to identify the word in 6 tries.
   - Everytime they try a word, they get clues on how close the word is to the secret word.

   Example1
   - Secret word: CHALK
   - Player input: BLAST
   - Output: RYGRR

   Example2:
   - Secret word: CLEAR
   - Player input: BLESE
   - Output: RGGRR

   Example3:
   - Secret word: CLEER
   - Player input: BLESE
   - Output: RGGRR
*/

public class WordleGame {

    public String play(String secret, String playerInput) {

        HashMap<Character, Integer> map = new HashMap();

        String output = "";

        for (int i=0; i < secret.length(); i++) {

            if (playerInput.charAt(i) != secret.charAt(i)) {

                if (map.containsKey(secret.charAt(i))) {
                    int count = map.get(secret.charAt(i));
                    count++;
                    map.put(secret.charAt(i), count);
                }
                else {
                    map.put(secret.charAt(i), 1);
                }
            }
        }

        for (int i=0; i < playerInput.length(); i++) {

            //System.out.println(map);
            //System.out.println("Character ->" + playerInput.charAt(i));

            if ((playerInput.charAt(i) == secret.charAt(i))) {
                output += "G";
            }
            else if(map.containsKey(playerInput.charAt(i)) && map.get(playerInput.charAt(i)) >= 1) {
                output += "Y";
                int count = map.get(playerInput.charAt(i));
                count--;
                map.put(playerInput.charAt(i), count);
            }
            else {
                output += "R";
            }
        }
        return output;

    }

    public static void main(String[] args) {

        WordleGame obj = new WordleGame();
        System.out.println(obj.play("CHALK", "BLAST"));// O/P - RYGRR
        System.out.println(obj.play("CLEAR", "BLESE"));// O/P - RGGRR
        System.out.println(obj.play("CLEER", "BLESE"));// O/P - RGGRY
        System.out.println(obj.play("CLEER", "LLESE"));// O/P = RGGRY
    }
}




