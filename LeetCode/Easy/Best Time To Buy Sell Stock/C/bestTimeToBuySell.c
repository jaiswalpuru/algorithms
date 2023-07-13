#define MAX(x, y) (((x) > (y)) ? (x) : (y))

int maxProfit(int* prices, int pricesSize){
    int minInd = 0;
    int profit = 0;
    for (int i=0; i<pricesSize; i++) {
        if (*(prices+minInd) > *(prices+i)) minInd = i;
        else profit = MAX(profit, *(prices+i)-*(prices+minInd));
    }
    return profit;
}
