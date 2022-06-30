package CodingInterviewQuestions;

import java.util.*;

/**
 * Asked in Surify Coding Interview, 30-June-2022
 */
// '''
// Given an array, print the Next Greater Element (NGE) for every element. The Next greater Element for an element x is the first greater
// element on the right side of x in the array. Elements for which no greater element exist, consider next greater element as -1.

// Examples:

// For an array, the rightmost element always has the next greater element as -1.
// For an array which is sorted in decreasing order, all elements have next greater element as -1.
// For the input array [4, 5, 2, 25], the next greater elements for each element are as follows.
// Element       NGE
//    4      -->   5
//    5      -->   25
//    2      -->   25
//    25     -->   -1
// d) For the input array [13, 7, 6, 12], the next greater elements for each element are as follows.

//   Element        NGE
//    13      -->    -1
//    7       -->     12
//    6       -->     12
//    12      -->     -1

// '''

//[4,5,2,25]

//Stack - [25], -1
//2 -> [25, 2], 5
//5 -> [25, 5], 25
//4 -> [25], 5

//[25, 22, 21, 20]
// 20 -> [20], -1
// 21 -> [21], -1
// 22 -> [22], -1
// 25 -> [25], -1


//Stack
//- Traverse from the end
//- push the element in the stack

public class NextGreaterElementArray {

    static int[] nextGreaterElement(int[] arr) {

        if (arr.length == 0) {
            return new int[0];
        }

        int[] out = new int[arr.length];
        Stack<Integer> stack = new Stack<Integer>();

        for(int i=arr.length-1; i >=0; i--){

            int nextGreaterElement = -1;

            while(stack.size() > 0 && stack.peek()<arr[i]) {
                stack.pop();
            }

            if (stack.size()>0) {
                nextGreaterElement = stack.peek();
            }

            out[i] = nextGreaterElement;
            stack.add(arr[i]);
        }

        return out;
    }

    public static void main(String[] args) {
        int[] input = {4,5,2,25};
        //int[] input = {13, 7, 6, 12};
        int[] output = NextGreaterElementArray.nextGreaterElement(input);
        System.out.println(Arrays.toString(output));
    }
}

