

//In this method we use the inbuilt method of java hashcode
/*
The hash code for a String object is computed(using int arithmetic) as −

s[0]*31^{(n - 1)} + s[1]*31^{(n - 2)} + ... + s[n - 1]s[0]∗31 
(n−1)
 +s[1]∗31 
(n−2)
 +...+s[n−1] , where s[i] is the ith character of the string, n is the length of the string.
 */

public class Codec {
	Map<Integer, String> map = new HashMap<>();
	public String encode(String longUrl) {
		map.put(longUrl.hashCode(), longUrl);
		return "https://tinyurl.com/" + longUrl.hashCode();
	}

	public String decode(String shortUrl) {
		return map.get(Integer.parseInt(shortUrl.replace("https://tinyurl.com/","")));
	}
}