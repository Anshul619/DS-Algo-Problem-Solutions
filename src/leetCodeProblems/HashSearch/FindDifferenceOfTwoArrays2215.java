package leetCodeProblems.HashSearch;

import java.lang.reflect.Array;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;

public class FindDifferenceOfTwoArrays2215 {

    public List<List<Integer>> findDifference(int[] nums1, int[] nums2) {

        HashSet set1 = new HashSet();
        HashSet set2 = new HashSet();

        List<List<Integer>> answer = new ArrayList<>();

        List answerList1 = new ArrayList();
        List answerList2 = new ArrayList();

        for(int i=0; i < nums1.length; i++) {
            set1.add(nums1[i]);
        }

        for(int i=0; i < nums2.length; i++) {
            set2.add(nums2[i]);
        }

        //System.out.println(set1);
        //System.out.println(set2);
        for (int i=0; i < nums1.length; i++) {

            if (!set2.contains(nums1[i]) && !answerList1.contains(nums1[i])) {
                answerList1.add(nums1[i]);
            }

        }

        answer.add(answerList1);

        for (int i=0; i < nums2.length; i++) {

            if (!set1.contains(nums2[i]) && !answerList2.contains(nums2[i])) {
                answerList2.add(nums2[i]);
            }

        }

        answer.add(answerList2);

        return answer;
    }

    public static void main(String[] args) {

        FindDifferenceOfTwoArrays2215 obj = new FindDifferenceOfTwoArrays2215();

        int[] nums1 = {1,2,3};
        int[] nums2 = {2,4,6};

        System.out.println(obj.findDifference(nums1, nums2));
    }
}
