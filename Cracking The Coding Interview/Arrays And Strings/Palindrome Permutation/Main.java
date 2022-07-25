import java.util.Scanner;

public class Main {
    
    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        String s = sc.nextLine();
        System.out.println(palindromePermutation(s));
    }

    public static boolean palindromePermutation(String s) {
        String commonString = s.toLowerCase();
        char[] str = commonString.toCharArray();
        int[] arr = new int[128];
        for (int i=0; i<str.length; i++){
            if (str[i]>='a' && str[i] <= 'z'){
                arr[str[i]]++;
            }
        }

        int oddCharCount = 0;
        for (int i=0;i<arr.length;i++){
            if (arr[i]%2==1) oddCharCount++;
        }
        if (oddCharCount>1) return false;
        return true;
    }
}
