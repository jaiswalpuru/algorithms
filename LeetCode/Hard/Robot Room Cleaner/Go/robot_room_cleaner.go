package main

/**
 * // This is the robot's control interface.
 * // You should not implement it, or speculate about its implementation
 * type Robot struct {
 * }
 *
 * // Returns true if the cell in front is open and robot moves into the cell.
 * // Returns false if the cell in front is blocked and robot stays in the current cell.
 * func (robot *Robot) Move() bool {}
 *
 * // Robot will stay in the same cell after calling TurnLeft/TurnRight.
 * // Each turn will be 90 degrees.
 * func (robot *Robot) TurnLeft() {}
 * func (robot *Robot) TurnRight() {}
 *
 * // Clean the current cell.
 * func (robot *Robot) Clean() {}
 */
type Pair struct {
	x, y int
}

var dir = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func backtrack(robot *Robot, r int, c int, d int, visited map[Pair]bool) {
	visited[Pair{r, c}] = true
	robot.Clean()
	for i := 0; i < 4; i++ {
		x, y := dir[d][0]+r, dir[d][1]+c
		if !visited[Pair{x, y}] && robot.Move() {
			backtrack(robot, x, y, d, visited)
		}
		d = (d + 1) % 4
		robot.TurnRight()
	}
	go_back(robot)
}

func go_back(robot *Robot) {
	robot.TurnRight()
	robot.TurnRight()
	robot.Move()
	robot.TurnRight()
	robot.TurnRight()
}

func cleanRoom(robot *Robot) {
	backtrack(robot, 0, 0, 0, map[Pair]bool{})
}
