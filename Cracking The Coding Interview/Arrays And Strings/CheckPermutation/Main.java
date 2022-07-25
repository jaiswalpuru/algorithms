import java.util.Scanner;

public class Main {
    
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String s1 = sc.nextLine();
        String s2 = sc.nextLine();
        System.out.println(checkPermutationTwo(s1, s2));
    }

    public static boolean checkPermutation(String s1, String s2) {
        if (s1.length() != s2.length()) return false;
        int char_set[] = new int[128];
        for (int i=0; i<s1.length();i++ ){
            char_set[s1.charAt(i)]++;
        }

        for (int i=0;i<s2.length(); i++){
            int val = s2.charAt(i);
            if (char_set[val] > 0) {
                char_set[val]--;
                continue;
            }
            return false;
        }
        return true;
    }

    public static String sort(String s) {
        char[] arr = s.toCharArray();
        java.util.Arrays.sort(arr);
        return new String(arr);
    }

    public static boolean checkPermutationTwo(String s1, String s2) {
        if (s1.length()!= s2.length()) return false;
        return sort(s1).equals(sort(s2));
    }
}

