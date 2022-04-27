package leetCodeProblems.TwoPointers;

public class StringReverseWordsInString {
	void reverse(char[] s, int start, int end ) {
        
        if (start<0 || end <0) {
            return;
        }
        
        int endIndex = end;
        
        for(int i=start;i<=(start+end)/2;i++){
            char temp = s[i];
            s[i] = s[endIndex];
            s[endIndex] = temp;
            endIndex --;
        }
    }

    String removeExtraSpaces(char [] s) {

        int index = 0;
        boolean cond = false;
        
        for (int i = 0; i < s.length; i++) {
            
            if (s[i] != ' ') {
                s[index] = s[i];
                index++;
                cond = true;
            } else {
                if (cond)
                    s[index++] = ' ';
                cond = false;
            }
            
        }
        
        if (index - 1 >= 0 && index - 1 < n && s[index - 1] == ' ')
            index--;
        
        return new String(s, 0, index);
        
    }
    
    public String reverseWords(String s) {
        
        char[] arr = s.toCharArray();
        int start = 0;
        
        for (int i=0; i<arr.length; i++) {
            
            if (arr[i] == ' ') {
                reverse(arr, start, i-1);
                start = i+1;
            }
        }
        
        reverse(arr, start, arr.length-1);
        
        reverse(arr, 0, arr.length-1);
        
        return removeExtraSpaces(arr);
    }

}
