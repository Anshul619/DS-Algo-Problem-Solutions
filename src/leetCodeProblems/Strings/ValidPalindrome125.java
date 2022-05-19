package leetCodeProblems.Strings;


public class ValidPalindrome125 {

    public boolean isPalindrome(String s) {

        int leftPointer = 0;
        int rightPointer = s.length()-1;

        while (leftPointer < rightPointer) {

            if (Character.isLetterOrDigit(s.charAt(leftPointer))
                    && Character.isLetterOrDigit(s.charAt(rightPointer))) {

                if (Character.toLowerCase(s.charAt(leftPointer)) ==
                        Character.toLowerCase(s.charAt(rightPointer))) {

                    leftPointer++;
                    rightPointer--;

                }
                else {
                    return false;
                }
            }
            else if (!Character.isLetterOrDigit(s.charAt(leftPointer))) {
                leftPointer++;
            }
            else if (!Character.isLetterOrDigit(s.charAt(rightPointer))) {
                rightPointer--;
            }

        }

        return true;
    }

    // Driver Code
    public static void main(String[] args)
    {
        ValidPalindrome125 obj = new ValidPalindrome125();

        //String str = "A man, a plan, a canal: Panama";

        String str = "race a car";
        //String str =  " ";
        System.out.println(obj.isPalindrome(str));
    }
}
