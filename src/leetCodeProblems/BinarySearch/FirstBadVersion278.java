package leetCodeProblems.BinarySearch;

/**
 * LeetCode - https://leetcode.com/problems/first-bad-version
 *
 * TimeComplexity - O(nlogn)
 * SpaceComplexity - O(1)
 */
public class FirstBadVersion278{

    int badVersion;

    FirstBadVersion278(int badVersion) {
        this.badVersion = badVersion;
    }

    private boolean isBadVersion(int version) {
        if (version == badVersion) {
            return true;
        }

        return false;
    }

    public int firstBadVersion(int n) {

        int left = 1;
        int right = n;
        int mid = 0;

        while (left < right) {

            mid = left+(right-left)/2;

            //System.out.println("Mid ->" + mid);
            if (isBadVersion(mid)) {
                right = mid;
            }
            else {
                left = mid+1;
            }
        }

        return left;
    }

    public static void main(String[] args) {

        FirstBadVersion278 obj = new FirstBadVersion278(2);

        System.out.println(obj.firstBadVersion(5));
    }
}
