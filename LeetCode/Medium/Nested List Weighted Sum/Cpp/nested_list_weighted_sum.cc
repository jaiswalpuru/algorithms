#include <iostream>
#include <vector>

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * class NestedInteger {
 *   public:
 *     // Constructor initializes an empty nested list.
 *     NestedInteger();
 *
 *     // Constructor initializes a single integer.
 *     NestedInteger(int value);
 *
 *     // Return true if this NestedInteger holds a single integer, rather than a nested list.
 *     bool isInteger() const;
 *
 *     // Return the single integer that this NestedInteger holds, if it holds a single integer
 *     // The result is undefined if this NestedInteger holds a nested list
 *     int getInteger() const;
 *
 *     // Set this NestedInteger to hold a single integer.
 *     void setInteger(int value);
 *
 *     // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 *     void add(const NestedInteger &ni);
 *
 *     // Return the nested list that this NestedInteger holds, if it holds a nested list
 *     // The result is undefined if this NestedInteger holds a single integer
 *     const vector<NestedInteger> &getList() const;
 * };
 */
//-------------------------dfs---------------------------
typedef std::vector<NestedInteger> vni;

class Solution {
public:
    int _depth(vni& nes_list, int d) {
        int sum = 0;
        int n = nes_list.size();
        for (int i=0; i<n; i++) {
            if (!nes_list[i].isInteger()) sum += _depth(nes_list[i].getList(), d+1);
            else sum += (d*nes_list[i].getInteger());
        }
        return sum;
    }
    int depthSum(vni& nes_list) {
        return _depth(nes_list, 1);
    }
};
//-------------------------dfs---------------------------

//-------------------------bfs---------------------------
typedef std::vector<NestedInteger> vni;
typedef std::queue<NestedInteger> qni;

class Solution {
public:
    int _depth(vni& nes_list, int d) {
        int sum = 0, depth = 1;
        qni q;
        for (int i=0; i<nes_list.size(); i++) {
            q.push(nes_list[i]);
        }
        while(!q.empty()) {
            size_t n = q.size();
            for (int i=0;i<n;i++) {
                NestedInteger p = q.front();
                q.pop();
                if (p.isInteger()) sum += d*p.getInteger();
                else {
                    for (NestedInteger ni : p.getList()) {
                        q.push(ni);
                    }
                }
            }
            d++;
        }
        return sum;
    }
    int depthSum(vni& nes_list) {
        return _depth(nes_list, 1);
    }
};
//-------------------------bfs---------------------------
