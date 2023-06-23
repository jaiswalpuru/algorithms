class Solution {
    public String simplifyPath(String path) {
        String[] str = path.split("/");
        ArrayDeque<String> res = new ArrayDeque<>();

        for (String s : str) {
            if (s.equals("..")) {
                if (res.size() > 0) res.removeLast();
            } else if (s.equals(".") || s.length() == 0) continue;
            else res.add(s);
        }
        
        StringBuilder sb = new StringBuilder();
        Iterator it = res.iterator();
        while(it.hasNext()) {
            String s = it.next().toString();
            if(s.length() > 0) {
                sb.append(s + "/");
            }
        }
        if (sb.length() > 1) 
            sb.deleteCharAt(sb.length()-1);
        return "/" + sb.toString();
    }
}
