#include <iostream>
#include <vector>
#include <queue>
#include <map>

struct P { int x, y, step; };

typedef std::vector<char> vc;
typedef std::vector<vc> vcc;
typedef std::queue<P> qp;
typedef std::vector<int> vi;
typedef std::vector<vi> vii;
typedef std::map<std::string, bool> mpb;

vii dir = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};

int nearest_exit_from_entrance_in_maze(vcc maze, vi ent) {
    int n = maze.size(), m = maze[0].size();
    qp q;
    q.push(P{ent[0], ent[1], 0});
	mpb visited;
	while(!q.empty()) {
        P curr = q.front();
        q.pop();
	    if (curr.x <0 || curr.y < 0 || curr.x >=n || curr.y >= m || visited[std::to_string
		(curr.x)+"->"+std::to_string(curr.y)] || maze[curr.x][curr.y] == '+') continue;
        if ((curr.x == 0 || curr.y == 0 || curr.x == n-1 || curr.y == m-1) &&
		(!(curr.x==ent[0]&&curr.y==ent[1]))) return curr.step;
    	visited[std::to_string(curr.x)+"->"+std::to_string(curr.y)] = true;
        for (int i=0; i<4; i++) {
            int x = curr.x+dir[i][0], y = curr.y+dir[i][1];
            if (x == ent[0] && y == ent[1]) continue;
            	q.push(P{x, y, curr.step+1});
        }
    }
    return -1;
}

int main() {
    std::cout<<nearest_exit_from_entrance_in_maze(vcc{{'+','+','.','+'},{'.','.','.','+'},{'+','+','+','.'}}, vi{1, 2})<<"\n"; 
}
