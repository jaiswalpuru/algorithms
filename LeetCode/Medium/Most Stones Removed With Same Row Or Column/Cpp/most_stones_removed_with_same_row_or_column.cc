#include <iostream>
#include <vector>
#include <map>

#define PB push_back

typedef std::vector<int> vi;
typedef std::vector<vi> vii;
typedef std::map<int, vii> mi;
typedef std::map<vi, bool> mb;

mi make_graph(vii &stones, bool is) {
    mi m;
    for (int i=0; i<stones.size(); i++) {
        if (is) m[stones[i][0]].PB(stones[i]);
        else m[stones[i][1]].PB(stones[i]);
    }
    return m;
}

void dfs(vii &stones, mi &xg, mi &yg, mb &visited, vi stone) {
    visited[stone] = true;
    for (int i=0;i<xg[stone[0]].size();i++) {
        if (!visited[xg[stone[0]][i]])
            dfs(stones, xg, yg, visited, xg[stone[0]][i]);
    }
    for (int i=0;i<yg[stone[1]].size();i++) {
        if (!visited[yg[stone[1]][i]])
            dfs(stones, xg, yg, visited, yg[stone[1]][i]);
    }
}

int most_stones_removed_with_same_row_or_col(vii stones) {
    mi xg = make_graph(stones, true);
    mi yg = make_graph(stones, false);
    int cnt = 0;
    mb visited;
    for (int i=0;i<stones.size(); i++) {
        if (!visited[stones[i]]) {
            cnt++;
            dfs(stones, xg, yg, visited, stones[i]);
        }
    }
    return stones.size() - cnt;
}

int main() {
    std::cout<<most_stones_removed_with_same_row_or_col(vii{
        {0,0},{0,1},{1,0},{1,2},{2,1},{2,2}
    })<<"\n";
}
