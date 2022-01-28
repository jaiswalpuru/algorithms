/*
Ninja is planning this 'N' days-long training schedule. Each day he can perform any one of these three activities.
(Running, fighting practice, learning new moves.) Each activity has some merit points on each day. As ninja has to improve
all his skills, he can't do the same activity in two consequtive days.
You are given a 2D array of size N*3 'POINTS' with the points corresponding to each day and activity.
Your task is to calculate the maximum number of merit points that ninja can earn.
*/

package main

import "fmt"

//brute force will give TLE------------------------------
func ninja(arr [][]int) int {
	n := len(arr)
	return _ninja(n-1, 3, arr)
}

func _ninja(day, last int, arr [][]int) int {
	if day == 0 {
		_max := 0
		for task := 0; task < 3; task++ {
			if task != last {
				_max = max(_max, arr[0][task])
			}
		}
		return _max
	}

	_max := 0
	for task := 0; task < 3; task++ {
		if task != last {
			points := arr[day][task] + _ninja(day-1, task, arr)
			_max = max(_max, points)
		}
	}

	return _max
}

//--------------------------------------------------

//------------------------------ninja memo -----------------------------------------
func ninja_memo(arr [][]int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 4)
		dp[i][0] = -1
		dp[i][1] = -1
		dp[i][2] = -1
		dp[i][3] = -1
	}

	return memo(arr, &dp, n-1, 3)
}

//last is the last task performed
func memo(arr [][]int, dp *[][]int, day int, last int) int {
	if day == 0 {
		_max := 0
		for task := 0; task < 3; task++ {
			if task != last {
				_max = max(_max, arr[0][task])
			}
		}
		return _max
	}

	if (*dp)[day][last] != -1 {
		return (*dp)[day][last]
	}

	_max := 0
	for task := 0; task < 3; task++ {
		if task != last {
			points := arr[day][task] + memo(arr, dp, day-1, task)
			_max = max(_max, points)
		}
	}
	(*dp)[day][last] = _max
	return (*dp)[day][last]
}

//------------------------------------------------------------------------------------

//-----------------------------------------Tabulation------------------------------------
func dp_ninja(arr [][]int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 4)
	}

	dp[0][0] = max(dp[0][1], dp[0][2])
	dp[0][1] = max(dp[0][0], dp[0][2])
	dp[0][2] = max(dp[0][1], dp[0][0])
	dp[0][3] = max(dp[0][1], max(dp[0][2], dp[0][0]))

	for day := 0; day < n; day++ {
		for last := 0; last < 4; last++ {
			dp[day][last] = 0
			_max := 0
			for task := 0; task < 3; task++ {
				if task != last {
					points := arr[day][task] + dp[day-1][task]
					_max = max(_max, points)
				}
			}
			dp[day][last] = _max
		}
	}
	return dp[n-1][3]
}

////------------------------------------------------------------------------------------

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(ninja_memo([][]int{
		{1, 2, 5},
		{3, 1, 1},
		{3, 3, 3},
	}))
}
