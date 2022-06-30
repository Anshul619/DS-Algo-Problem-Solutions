package CodingInterviewQuestions;

/**
 * Asked in Google, 30thJune2022 - PreQualifier Round
 */
  /*Wordle is a game where there's a secret 5-letter word. The player has to identify the word in 6 tries. Everytime they try a word, they get clues on how close the word is to the secret word.

        For example:
        Secret word: CHALK
        Player input: BLAST

        Output: RYGRR

        Secret word: CLEAR
        Player input: BLESE

        HashSet - C, L, E, R

        Secret word: CLEER
        Player input: BLESE

        HashSet - C, L, E, R

        Ouput = RGGRR

        HashSet
        - Iterate through secret word, save every character, its position in the HashSet
        - Iterate player input, check every character in the HashMap.
        - its index is same as index in player input, then "G" in output.


public String wordleGame(String[] secret, String[] playerInput) {

        HashMap<Character, Integer> map = new HashMap();

        String output = "";

        for(int i=0; i < secret.length(); i++) {

        if (map.containsKey(secret.charAt(i)) {
        int count = map.get(secret.charAt(i);
        count++;
        map.put(secret.charAt(i), count);
        }
        else {
        map.put(secret.charAt(i), 1);
        }
        }

        for(int i=0; i < playerInput.length(); i++) {
        if (playerInput.charAt(i).equals(secret.charAt(i) && map.get(playerInput.charAt(i)) >= 1) {
        output += "G";
        int count = map.get(playerInput.charAt(i);
        count--;
        map.put(secret.charAt(i), count);
        }
        else if(map.containsKey(playerInput.charAt(i)) && map.get(playerInput.charAt(i)) >= 1) {
        output += "Y";
        int count = map.get(playerInput.charAt(i);
        count--;
        map.put(secret.charAt(i), count);
        }
        else {
        output += "R";
        }
        }

        return output;

        }*/

/*
1 2 10

Median
- Sort these


3 2 1
- min - 1
- max - 3

2 3 1

1 3 1
- min - 1
- max - 3

public int getMedian(int[] arr) {

    Arrays.sort(arr); // O(nlogn)

    int middle = arr[1];

    return middle;
}

 */


