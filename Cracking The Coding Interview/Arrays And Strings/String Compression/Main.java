import java.util.Scanner;

public class Main {
    
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String s = sc.nextLine();
        System.out.println(stringCompressionTwo(s));
    }

    //this method is O(n * k^2) because string concatenation is O(n^2)
    public static String stringCompressionOne(String s) {
        String res = "";
        int cnt = 0;
        for (int i=0; i<s.length();i++){
            cnt++;
            if (i+1 >= s.length() || s.charAt(i) != s.charAt(i+1)){
                res += "" + s.charAt(i) + cnt;
                cnt = 0;
            }
        }
        return res.length() < s.length() ? res : s;
    }

    //the time complexity problem can be solved using the string builder as it creates a resizable array
    public static String stringCompressionTwo(String s) {
        StringBuilder builder = new StringBuilder();
        int cnt = 0;
        for (int i=0;i<s.length();i++){
            cnt++;
            if (i+1 >= s.length() || s.charAt(i)!=s.charAt(i+1)){
                builder.append(s.charAt(i));
                builder.append(cnt);
                cnt = 0;
            }
        }
        return builder.length() < s.length() ? builder.toString() : s;
    }
}
