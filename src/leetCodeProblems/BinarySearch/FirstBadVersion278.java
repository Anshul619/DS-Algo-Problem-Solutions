package leetCodeProblems.BinarySearch;

/**
 * LeetCode - https://leetcode.com/problems/first-bad-version
 *
 * TimeComplexity - O(logn)
 * SpaceComplexity - O(1)
 */
public class Version278{

    int badVersion;

    Version278(int badVersion) {
        this.badVersion = badVersion;
    }

    private boolean isBadVersion(int version) {
        if (version == badVersion) {
            return true;
        }

        return false;
    }

    public int Version(int n) {

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

        Version278 obj = new Version278(2);

        System.out.println(obj.Version(5));
    }
}
