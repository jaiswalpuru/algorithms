class Solution {
public:
    vector<int> finalPrices(vector<int>& prices) {
        vector<int> dis_prices(prices.size());
        stack<int> st;

        dis_prices[prices.size()-1] = prices[prices.size()-1];
        st.push(prices[prices.size()-1]);
        for (int i = prices.size()-2; i >= 0; i--) {
            while (!st.empty() && st.top() > prices[i]) st.pop();
            if (st.empty()) dis_prices[i] = prices[i];
            else dis_prices[i] = prices[i] - st.top();
            st.push(prices[i]);
        }

        return dis_prices;
    }
};
