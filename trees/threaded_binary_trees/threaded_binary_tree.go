package main

type ThreadedBT struct {
	left  *ThreadedBT
	right *ThreadedBT
	lTag  int
	rTag  int
	data  int
}
