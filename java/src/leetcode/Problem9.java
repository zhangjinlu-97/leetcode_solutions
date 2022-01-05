package leetcode;

public class Problem9 {
    public boolean isPalindrome(int x) {
        if (x < 0) {
            return false;
        }
        int num = x;
        int rx = 0;
        while (num > 0) {
            rx *= 10;
            rx += num % 10;
            num /= 10;
        }
        return x == rx;
    }

    public static void main(String[] args) {
        boolean ans = new Problem9().isPalindrome(121);
        System.out.println(ans);
    }
}
