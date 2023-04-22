import java.util.HashMap;
import java.util.Map;

public class MaximumWordProd {
    public static void main(String[] args){
        System.out.println(maximumWordProd(new String[]{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}));
    }

    public static int maximumWordProd(String[] words) {
        int res =0;
        for (int i=0;i<words.length;i++){
            Map<Character, Boolean> mp = new HashMap<>();
            for (int j=0;j<words[i].length();j++){
                mp.put(words[i].charAt(j), true);
            }
            
            Boolean flag = false;
            for (int j=i+1;j<words.length;j++){
                for (int k=0;k<words[j].length();k++){
                    if (mp.containsKey(words[j].charAt(k))) {
                        flag=true;
                        break;
                    }
                }
                if (flag==false){
                    res = Math.max(res, words[j].length()*words[i].length());
                }
                flag= false;
            }
        }
        return res;
    }
}
