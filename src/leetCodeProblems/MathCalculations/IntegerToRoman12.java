package leetCodeProblems.MathCalculations;

/**
 * LeetCode Link - https://leetcode.com/problems/integer-to-roman/
 * 
 * Interview Bit Link - https://www.interviewbit.com/problems/integer-to-roman/
 * 
 * About Roman Numbers - https://projecteuler.net/about=roman_numerals
 * 
 * @author anshul.agrawal
 *
 */
public class IntegerToRoman12 {
	private String getRomanSymbolAccordingToFraction(int fraction) {
        
        switch(fraction) {
            case 1000:
                return "M";
            case 900:
                return "CM";
            case 500:
                return "D";
            case 400:
                return "CD";
            case 100:
                return "C";
            case 90:
                return "XC";
            case 50:
                return "L";
            case 40:
                return "XL";
            case 10:
                return "X";
            case 9:
                return "IX";
            case 5:
                return "V";
            case 4:
                return "IV";
            case 1:
                return "I";
        }
        
        return "";
    }
    
	private String appendRomanSymbols(int num, int fraction) {
        
        String romanSymbol = "";
        
        if (num/fraction > 0) {
            int counts = num/fraction;
            
            for (int i=0; i<counts; i++) {
                romanSymbol += getRomanSymbolAccordingToFraction(fraction);
            }
        }
        
        return romanSymbol;
    }
    
	// Driver Method
    public String intToRoman(int num) {
        
        int leftOverNum = num;
        int[] fractionsArray = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
        
        String ansRomanNumber = "";
        
        for(int i=0; i<fractionsArray.length; i++) {
            ansRomanNumber += appendRomanSymbols(leftOverNum, fractionsArray[i]); 
            leftOverNum = leftOverNum%fractionsArray[i];
        }   
        
        return ansRomanNumber;
    }
    
    public static void main(String[] args) {
    	
    	int number = 1994;
    	
    	IntegerToRoman12 obj = new IntegerToRoman12();

    	System.out.println(obj.intToRoman(number));
    }
}
