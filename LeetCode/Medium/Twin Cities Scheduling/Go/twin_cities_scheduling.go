package main

import ("fmt"; "sort")

type I [][]int
func (a I) Len() int { return len(a) }
func (a I) Less(i, j int) bool { return a[i][0]-a[i][1] < a[j][0]-a[j][1] }
func (a I) Swap(i,j int) { a[i], a[j] = a[j], a[i] }


type I2 [][]int
func (a I2) Len() int { return len(a) }
func (a I2) Less(i, j int) bool { return a[i][1] < a[j][1] }
func (a I2) Swap(i,j int) { a[i], a[j] = a[j], a[i] }

func twin_cities_scheduling(cost [][]int) int {
	var a I
	a = cost
	sort.Sort(a)
	min_cost := 0
	n := len(a)
	i,j := 0, n-1
	for i < j {
		min_cost += a[i][0]
		min_cost += a[j][1]
		i++
		j--
	}
	return min_cost
}

func main() {
	fmt.Println(twin_cities_scheduling([][]int{{259,770},{448,54},{926,667},{184,139},{840,118},{577,469}}))
}