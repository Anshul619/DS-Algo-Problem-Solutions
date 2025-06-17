package leetCodeProblems.MathCalculations;

public class RomanToInteger13 {

	public int romanToInt(String s) {
        
        int answerInt = 0;
        
        for (int i=0; i < s.length(); i++) {
            
            boolean twoCharPairsMatched = false;
            
            if (i+1 < s.length()) {
                
                StringBuilder sb = new StringBuilder();
                sb.append(s.charAt(i));
                sb.append(s.charAt(i+1));
                
                String twoCharsPairs = sb.toString();
                
                switch (twoCharsPairs) {
                    case "CM":
                        answerInt += 900;
                        twoCharPairsMatched = true;
                        break;
                    case "CD":
                        answerInt += 400;
                        twoCharPairsMatched = true;
                        break;
                    case "XC":
                        answerInt += 90;
                        twoCharPairsMatched = true;
                        break;
                    case "XL":
                        answerInt += 40;
                        twoCharPairsMatched = true;
                        break;
                    case "IX":
                        answerInt += 9;
                        twoCharPairsMatched = true;
                        break;
                    case "IV":
                        answerInt += 4;
                        twoCharPairsMatched = true;
                        break;
                    default:
                        //Nothing
                        
                }
            }
            
            if (twoCharPairsMatched) {
                i++;
            }
            else {
                switch(s.charAt(i)) {
                    case 'M':
                        answerInt += 1000;
                        break;
                    case 'D':
                        answerInt += 500;
                        break;
                    case 'C':
                       answerInt += 100;
                       break;
                    case 'L':
                        answerInt += 50;
                        break;
                    case 'X':
                        answerInt += 10;
                        break;
                    case 'V':
                       answerInt += 5;
                       break;
                    case 'I':
                       answerInt += 1;
                       break;
                    default:
                        //Nothing
               } 
            }
        }
        
        return answerInt;
        
    }
	
	public static void main(String[] args) {
		
		String roman = "MCMXCIV";
		
		RomanToInteger13 obj = new RomanToInteger13();
		
		System.out.println(obj.romanToInt(roman));
	}
}
