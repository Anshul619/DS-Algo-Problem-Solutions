package leetCodeProblems.BruteForce;

/**
 * LeetCode - https://www.interviewbit.com/problems/implement-strstr/
 * InterviewBit - https://www.interviewbit.com/problems/implement-strstr/
 * 
 * @author anshul.agrawal
 *
 */
public class ImplementStrStr28 {
	
	public int strStr(String haystack, String needle) {
        
        if (needle.isEmpty()) {
            return 0;
        }
        
        int needleIndex = 0;
        int haystackMatchIterationIndex = 0;
        
        for(int haystackIndex=0; haystackIndex < haystack.length(); haystackIndex++) {
            
            //System.out.println("haystackIndex ->" + haystackIndex + ", needleIndex ->" + needleIndex);
            if (haystack.charAt(haystackIndex) == needle.charAt(needleIndex)) {
                
                haystackMatchIterationIndex = haystackIndex+1;
                
                for(needleIndex=1; needleIndex < needle.length(); needleIndex++) {
                    
                    if (haystackMatchIterationIndex < haystack.length() && haystack.charAt(haystackMatchIterationIndex) == needle.charAt(needleIndex)) {
                        haystackMatchIterationIndex++;   
                    }
                    else {
                        needleIndex = 0;
                        break;
                    }
                }
                
                // All characters are matching
                if (needleIndex == needle.length()) {
                    
                    //System.out.println("needleIndex -> "+ needleIndex);
                    //System.out.println("needle Length -> "+ needle.length());
                    return haystackIndex;
                }
                
            }
        }
        
        return -1;
        
    }
	
	public static void main(String[] args) {
		
		String hayStack = "hello";
		String needle = "ll";
		
		ImplementStrStr28 obj = new ImplementStrStr28();
		
		System.out.println(obj.strStr(hayStack, needle));
		
	}

}
