class Solution {
public:
    bool isIsomorphic(string s, string t) {
        int s_map[129];
        int t_map[129];
        fill(s_map, s_map+129, -1);
        fill(t_map, t_map+129, -1);
        for (int i=0; i<s.length(); i++) {
            int si = s[i];
            int ti = t[i];
            if (s_map[si] != -1 && s_map[si] != ti) return false;
            if (t_map[ti] != -1 && t_map[ti] != si) return false;
            s_map[si] = ti;
            t_map[ti] = si;
        }
        return true;
    }
};
