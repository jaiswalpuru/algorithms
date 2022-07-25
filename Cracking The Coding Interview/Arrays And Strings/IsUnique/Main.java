import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String s;
        s = sc.nextLine();
        System.out.println(isUnique(s));
    }

    public static boolean isUnique(String s) {
        if (s.length() > 128) return false;
        
        boolean visited[] = new boolean[128];
        for (int i=0; i<s.length(); i++) {
            int val = s.charAt(i);
            if (visited[val]) return false;
            visited[val] = true;
        }
        return true;
    }

    //this method assumes that all the characters are lowercase
    public static boolean isUniqueTwo(String s) {
        int checker = 0;
        for (int i =0;i <s.length(); i++) {
            int val = s.charAt(i)-'a';
            if ((checker & (1 << val)) > 0) return false;
            checker |= (1<<val);
        }
        return true;
    }

}
