package main

type NumMatrix struct {
    prefix_sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
    dp := make([][]int, len(matrix)+1)
    for i:=0;i<=len(matrix);i++{
        dp[i] = make([]int, len(matrix[0])+1)
    }
    for i:=0;i<len(matrix);i++{
        for j:=0;j<len(matrix[0]);j++{
            dp[i][j+1] = dp[i][j]+matrix[i][j]
        }
    }
    return NumMatrix{prefix_sum:dp}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    sum := 0
    for i:=row1; i<=row2; i++{
        sum += this.prefix_sum[i][col2+1] - this.prefix_sum[i][col1]
    }
    return sum
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */


