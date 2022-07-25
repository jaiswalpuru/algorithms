import java.util.Scanner;

public class Main {
    
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String s = sc.nextLine();
        int trueLength = sc.nextInt();
        urLify(s, trueLength);
    }

    public static void urLify(String s, int trueLength) {
        char[] str = s.toCharArray();
        int numOfSpaces = countCharacters(str, 0, trueLength, ' ');
        int newIndex = trueLength-1 + numOfSpaces*2;

        if (newIndex+1 < str.length) str[newIndex-1] = '\0';
        for (int i=trueLength-1; i>=0; i--){
            if (str[i] == ' ') {
                str[newIndex] = '0';
                str[newIndex-1] = '2';
                str[newIndex-2] = '%';
                newIndex-=3;
            }else{
                str[newIndex] = str[i];
                newIndex-=1;
            }
        }
        System.out.println(str);
    }

    public static int countCharacters(char[] arr, int start, int end, int target) {
        int cnt = 0;
        for (int i=start;i<end; i++){
            if (arr[i] == target){
                cnt++;
            }
        }
        return cnt;
    }
}
