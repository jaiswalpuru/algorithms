import java.util.Scanner;

public class Main {
    
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String s1 = sc.nextLine();
        String s2 = sc.nextLine();
        
        System.out.println(oneWay(s1,s2));
    }

    public static boolean oneWay(String str1, String str2) {
        String s1 = str1.length() < str2.length() ? str2 : str1;
        String s2 = str1.length() < str2.length() ? str1 : str2;
        int n = s1.length();
        int m = s2.length();
        int i = 0;
        int j = 0;
        int differ = 0;

        if (m-n > 1) return false;
        while(i < s1.length() && j < s2.length()) {
            if (s1.charAt(i) != s2.charAt(j)){
                differ++;
                if (differ > 1) return false;
                if (s1.length() == s2.length()) j++;
            }else {
                j++;
            }
            i++;
        }
        return true;
    }
}
