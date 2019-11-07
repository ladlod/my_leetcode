class leetcode9 {
    public static void main(String args[]) {
        int x = 121;
        System.out.println(isPalindrome(x));
    }

    public static boolean isPalindrome(int x) {
        if (x < 0)
            return false;
        int num = x;
        int _x = 0;
        while (x >= 10) {
            _x = _x * 10 + x % 10;
            x /= 10;
        }
        _x = _x * 10 + x % 10;
        if (_x == num)
            return true;
        else
            return false;
    }
}