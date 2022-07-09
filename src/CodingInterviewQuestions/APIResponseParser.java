package CodingInterviewQuestions;

/**
 * Arcesium Coding Test, 8-July-2022
 */

//import com.google.gson.Gson;
//import com.google.gson.GsonBuilder;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.ArrayList;
import java.util.List;

class APIResponseParser {

    static class GeoCode {
        String lat;
        String lng;
    }

    static class Address {
        String street;
        String suite;
        String city;
        String zipcode;
        GeoCode geo;
    }

    static class Company {
        String name;
        String basename;
    }

    static class User {
        int id;
        String name;
        String username;
        String email;
        Address address;
        Company company;
        String website;
    }

    /*
     * Complete the 'apiResponseParser' function below.
     *
     * The function is expected to return an INTEGER_ARRAY.
     * The function accepts following parameters:
     *  1. STRING_ARRAY inputList
     *  2. INTEGER size
     */

    private static User[] parseJSONString(String jsonString) {

        //GsonBuilder builder = new GsonBuilder();
        //builder.setPrettyPrinting();

        //Gson gson = builder.create();
        //User[] usersList = gson.fromJson(jsonString, User[].class);

        //return usersList;

        return new User[0];

    }
    private static User[] callAPI() {

        try {
            URL url = new URL("https://raw.githubusercontent.com/arcjsonapi/ApiSampleData/master/api/users");
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod("GET");

            BufferedReader br = new BufferedReader(new InputStreamReader(conn.getInputStream()));
            StringBuilder sb = new StringBuilder();
            String line;
            while ((line = br.readLine()) != null) {
                sb.append(line+"\n");
            }
            br.close();
            conn.disconnect();

            String apiResponse = sb.toString();

            User[] usersList = parseJSONString(apiResponse);
            return usersList;
        }
        catch(IOException exception) {

        }

        return new User[0];
    }

    private static String getFieldValue(User userObj, String inputField) {

        String userFieldValue = "";

        /*
        Reflection needs to be implemented here. Due to time limitation, commentout

        try {
            Field fieldValue = userObj.getClass().getDeclaredField(inputField);

            String value = (String) fieldValue.get(userObj);
            System.out.println(value);

            return value;
        }
        catch(IllegalAccessException exception) {

        }
        catch(NoSuchFieldException exception) {

        }*/


        switch(inputField) {
            case "name":
                userFieldValue = userObj.name;
                break;
            case "username":
                userFieldValue = userObj.username;
                break;
            case "email":
                userFieldValue = userObj.email;
                break;
            case "address.street":
                userFieldValue = userObj.address.street;
                break;
            case "address.suite":
                userFieldValue = userObj.address.suite;
                break;
            case "address.city":
                userFieldValue = userObj.address.city;
                break;
            case "address.zipcode":
                userFieldValue = userObj.address.zipcode;
                break;
            case "address.geo.lat":
                userFieldValue = userObj.address.geo.lat;
                break;
            case "address.geo.lng":
                userFieldValue = userObj.address.geo.lng;
                break;
            case "website":
                userFieldValue = userObj.website;
                break;
            case "company.name":
                userFieldValue = userObj.company.name;
                break;
            case "company.basename":
                userFieldValue = userObj.company.basename;
                break;
        }

        return userFieldValue;
    }

    public static List<Integer> apiResponseParser(List<String> inputList, int size) {

        List<Integer> output = new ArrayList<>();

        User[] usersList = callAPI();

        String inputField = inputList.get(0);
        String inputFieldOperation = inputList.get(1);
        String inputFieldValue = inputList.get(2);

        for (int i=0; i < usersList.length; i++) {

            String userFieldValue = getFieldValue(usersList[i], inputField);

            switch(inputFieldOperation) {
                case "EQUALS":
                    if (userFieldValue.equals(inputFieldValue)) {
                        output.add(usersList[i].id);
                    }
                    break;
                case "IN":
                    if (inputFieldValue.contains(userFieldValue)) {
                        output.add(usersList[i].id);
                    }
                    break;
                default:
                    //Nothing
            }
        }

        if (output.size() == 0) {
            output.add(-1);
        }

        return output;
    }

}
