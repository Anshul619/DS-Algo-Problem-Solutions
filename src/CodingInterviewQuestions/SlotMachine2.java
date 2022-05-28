package CodingInterviewQuestions;

/**
 * Asked in Twillo Coding Round, 28thMay2022
 */

import java.io.*;
import java.util.*;

public class SlotMachine2 {

    public static int slotWheels(List<String> history) {

        ArrayList<ArrayList<Integer>> intList = new ArrayList<>();

        int numberOfWheels = 0;

        if (history.size() > 0) {
            numberOfWheels = history.get(0).length();
        }

        for(int i=0; i < history.size(); i++) {

            ArrayList<Integer> temp = new ArrayList<>();

            for (int j=0; j < history.get(i).length(); j++) {

                //Integer.valueOf(
                temp.add(Character.getNumericValue(history.get(i).charAt(j)));

            }

            intList.add(temp);
        }
        //System.out.println("Before ->" + intList);

        for (int i=0; i < intList.size(); i++) {
            Collections.sort(intList.get(i));
        }

        //System.out.println("After ->" + intList);

        int totalMax = 0;

        for (int j=0; j <numberOfWheels;j++) {

            int localMax = 0;

            for (int i=0; i < intList.size(); i++) {
                localMax = Math.max(localMax, intList.get(i).get(j));
            }

            totalMax += localMax;

        }

        return totalMax;
    }

    public static void main(String[] args) throws IOException {

        List<String> history = new ArrayList<>();
        history.add("137");
        history.add("364");
        history.add("115");
        history.add("724");
        System.out.println("TC1 o/p ->" + SlotMachine2.slotWheels(history));  // Excepted o/p = 14;

        history.clear();
        history.add("712");
        history.add("246");
        history.add("365");
        history.add("312");
        System.out.println("TC2 o/p ->" + SlotMachine2.slotWheels(history));  // Excepted o/p = 15;

        history.clear();
        history.add("1112");
        history.add("1111");
        history.add("1211");
        history.add("1111");
        System.out.println("TC3 o/p ->" + SlotMachine2.slotWheels(history));  // Excepted o/p = 5;
    }
}

