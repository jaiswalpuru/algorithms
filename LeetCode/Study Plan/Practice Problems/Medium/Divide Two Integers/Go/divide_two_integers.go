package main

func divide_two_integers(dividend, divisor int) int {
	neg := 0
    
    if dividend == math.MinInt32 && divisor == -1 {
        return math.MaxInt32
    }
    
    if dividend < 0 {
        neg++
        dividend = - dividend
    }
    
    if divisor < 0 {
        neg++
        divisor = -divisor
    }
    
    sub := 0
    for dividend - divisor >= 0 {
        sub++
        dividend -= divisor
    }
    
    if neg == 1{
        sub=-sub
    }
    
    return sub
}

func main() {
	fmt.Println(divide_two_integers(10,3))

}