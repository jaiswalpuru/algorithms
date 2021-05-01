//The area of a circle is defined as πr^2. Estimate π to 3 decimal places using a Monte Carlo method.

//Hint: The basic equation of a circle is x2 + y2 = r2.

/*
Monte Carlo methods or monte carlo experiments are a broad set of computational algorithms that rely on random
sampling to obtain numerical results. The underlying concept is to use randomness to solve problems that might be
deterministic in principle.

Monte Carlo methods vary but tend to follow a particular pattern :
1. Define a domain of possible inputs.
2. Generate inputs randomly from probability distribution over domain.
3. Perform a deterministic computation on the inputs.
4. Aggregate the results.

Estimation of π , consider a quadrant inscribed in a unit square. r=1/2
1. Draw a square then inscribe a quadrant within it.
2. Uniformly scatter the points over the square.
3. Count the number of points in the quadrant that is having a distance from origin of less than 1.
4. The ratio of the inside-count and the total-sample-count is an estimate of the ratio of the two areas, π/4
   Multiply the result by 4 to estimate π.

area(circle)    no. of points generated inside the cirlce
------------ =  ------------------------------------------
area(square)    total no. of points generated inside square

        no. of points inside cirlce
π = 4 * ---------------------------
        no. of points inside square

*/

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	INTERVAL = 100000
)

func main() {

	var pi, randX, randY, distFromOrigin float64
	rand.Seed(time.Now().UTC().UnixNano())
	circlePoints, squarePoints := 0, 0

	for i := 0; i < INTERVAL; i++ {
		randX = float64(rand.Int()%(INTERVAL+1)) / float64(INTERVAL)
		randY = float64(rand.Int()%(INTERVAL+1)) / float64(INTERVAL)
		distFromOrigin = math.Sqrt(randX*randX + randY*randY)

		//check if the point lies inside the circle with r = 1
		if distFromOrigin <= 1 {
			circlePoints++
		}
		squarePoints++

		pi = float64(4*circlePoints) / float64(squarePoints)

		fmt.Println(randX, randY, circlePoints, squarePoints, " - ", pi)

	}

	fmt.Printf("π := %f\n", pi)

}
