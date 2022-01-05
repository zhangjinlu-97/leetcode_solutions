package leetcode;

import java.util.HashMap;

public class Problem166 {
    public String fractionToDecimal(int numerator, int denominator) {
        if (numerator == 0) {
            return "0";
        }
        StringBuilder ans = new StringBuilder();
        ans.append(((numerator > 0) ^ (denominator > 0)) ? "-" : "");  // + or -
        long num = Math.abs((long) numerator);
        long den = Math.abs((long) denominator);
        ans.append(num / den);
        num %= den;
        if (num != 0) {
            ans.append(".");
            HashMap<Long, Integer> indexMap = new HashMap<>();
            indexMap.put(num, ans.length());
            while (num != 0) {
                num *= 10;
                ans.append(num / den);
                num %= den;
                if (indexMap.containsKey(num)) {
                    ans.insert(indexMap.get(num), "(");
                    ans.append(")");
                    break;
                }
                indexMap.put(num, ans.length());
            }
        }
        return ans.toString();
    }

    public static void main(String[] args) {
        Problem166 p = new Problem166();
        System.out.println(p.fractionToDecimal(0, -5));
    }
}
