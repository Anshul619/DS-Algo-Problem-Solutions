package leetCodeProblems.HashSearch;

/**
 * Interview Bit - https://www.interviewbit.com/problems/pair-with-given-difference/
 */
import java.util.*;

public class FindIfPairExistsWithGivenDifference {
	
	public int solve(int[] A, int B) {

        //Arrays.sort(A);

        HashMap<Integer, Integer> hashMap = new HashMap<Integer, Integer>();

        for(int i=0; i<A.length; i++) {
        	hashMap.put(A[i], i);
        }

        for(int i=0; i<A.length; i++) {

            int targetNumUsingSum = B+A[i];
            int targetNumUsingDiff = A[i]-B;
            
            System.out.println("B ->" + B);
            System.out.println("A[i] ->" + A[i]);
            
            System.out.println("targetNumUsingSum ->" + targetNumUsingSum);
            System.out.println("targetNumUsingDiff ->" + targetNumUsingDiff);
            
            if (hashMap.containsKey(targetNumUsingSum)) {
            	System.out.println("targetNumUsingSum is TRUE");
            }
            
            if (hashMap.containsKey(targetNumUsingDiff)) {
            	System.out.println("targetNumUsingDiff is TRUE");
            }
            
            if ((hashMap.containsKey(targetNumUsingSum) && hashMap.get(targetNumUsingSum) != i) || 
            	(hashMap.containsKey(targetNumUsingDiff) && hashMap.get(targetNumUsingDiff) != i)) {
                return 1;
            }
        }

        return 0;
    }
	
	public static void main(String[] args) {
		
		//int[] inputArray = {-533,-666,-500,169,724,478,358,-38,-536,705,-855,281,-173,961,-509,-5,942,-173,436,-609,-396,902,-847,-708,-618,421,-284,718,895,447,726,-229,538,869,912,667,-701,35,894,-297,811,322};
		//int targetDifference = 369;
		
		int[] inputArray = {90, 497, -411, 990, 145, 353, 314, -349, -260, -956, 258, -665, -241, 192, 605, 264, -819, -497, 829, 775, -392, 292, 997, 549, 556, 561, 627, -533, 541, -871, 240, 813, 174, -399, -923, -785};
		int targetDifference = -466;
				
		//int[] inputArray = {20, 500, 1000, 200};
		//int targetDifference = 0;
		
		FindIfPairExistsWithGivenDifference obj = new FindIfPairExistsWithGivenDifference();
		
		System.out.println(obj.solve(inputArray, targetDifference));
	}

}
