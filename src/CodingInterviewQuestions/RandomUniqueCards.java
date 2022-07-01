package CodingInterviewQuestions;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.Random;

/**
 * Asked in TouchNote, 1stJuly, 2022.
 */

/*
Touchnote cards are made up of illustrations (artwork) and templates (a json object describing the location and size of illustrations and the location and size of the customer's photo). Together these are called 'Themes'
Themes are configured manually and sometimes duplicates are added in error

We'd like to provide clients with a random ordered list of unique themes.
Create a function/functions that takes a raw list of themes and returns a randomly ordered list of unique themes

Themes json array:

{ "illustrationId": "117", templateId: "auidk-oiwd", tag: ["birthday", "thankyou"] }
{ "illustrationId": "245", templateId: "bdapq-aioy", tag: ["eid"] }
{ "illustrationId": "117", templateId: "auidk-oiwd", tag: ["birthday", "thankyou"] }
{ "illustrationId": "117", templateId: "cdkjk-adsf", tag: ["mothers day"] }
{ "illustrationId": "120", templateId: "cdkjk-adsf", tag: ["mothers day"] }
{ "illustrationId": "324", templateId: "auidk-oiwd", tag: ["fathers day"] }

Response:

{ "illustrationId": "117", templateId: "auidk-oiwd", tag: ["birthday", "thankyou"] }
{ "illustrationId": "245", templateId: "bdapq-aioy", tag: ["eid"] }
{ "illustrationId": "117", templateId: "cdkjk-adsf", tag: ["mothers day"] }
{ "illustrationId": "120", templateId: "cdkjk-adsf", tag: ["mothers day"] }
{ "illustrationId": "324", templateId: "auidk-oiwd", tag: ["fathers day"] }

*/

public class RandomUniqueCards {
    static class Theme {
        String illustrationId;
        String templateId;
        String[] tag;

        Theme(String illustrationId, String templateId, String[] tag) {
            this.illustrationId = illustrationId;
            this.templateId = templateId;
            this.tag = tag;
        }
    }

    public ArrayList<Theme> getRandomThemes(ArrayList<Theme> input) {

        HashSet<String> set = new HashSet();

        ArrayList<Theme> tempList = new ArrayList();
        ArrayList<Theme> output = new ArrayList();

        for (int i = 0; i < input.size(); i++) {

            String uniqueKey = input.get(i).illustrationId + "_" + input.get(i).templateId;

            if (!set.contains(uniqueKey)) {
                set.add(uniqueKey);
                tempList.add(input.get(i));
            }
        }

        for (int i = 0; i < tempList.size(); i++) {

            int randomNumber = new Random().nextInt(tempList.size());

            output.add(tempList.get(randomNumber));
            tempList.remove(randomNumber);
        }

        return output;
    }

}