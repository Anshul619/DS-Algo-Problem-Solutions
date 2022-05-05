package leetCodeProblems.Arrays;

public class PascalTriangleKthRow {

	public int[] getRow(int A) {
       int c = 1;
       A += 1; // Convert it to normal A

       int[] output = new int[A];
       
       for(int i=1; i<=A;i++){
       
    	   output[i-1] = c;
           c = c*(A-i)/i;
       }
       
       return output;
   }


}
