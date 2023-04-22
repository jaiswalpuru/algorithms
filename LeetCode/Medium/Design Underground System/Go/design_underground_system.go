package main

type User struct {
    check_in string
    checkin_time int
}

type Pair struct {
    total int
    n int
}

type UndergroundSystem struct {
    user_data map[int]User
    src_des map[string]map[string]Pair
}


func Constructor() UndergroundSystem {
    return UndergroundSystem{user_data:make(map[int]User), src_des:make(map[string]map[string]Pair)}
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
    this.user_data[id] = User{check_in : stationName, checkin_time:t}
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
    time_diff := t-this.user_data[id].checkin_time
    if _,ok:=this.src_des[this.user_data[id].check_in];!ok{
        this.src_des[this.user_data[id].check_in] = make(map[string]Pair)
    }
    val := this.src_des[this.user_data[id].check_in][stationName].total
    n := this.src_des[this.user_data[id].check_in][stationName].n
    this.src_des[this.user_data[id].check_in][stationName] = Pair{total:val+time_diff, n:n+1}
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    return float64(this.src_des[startStation][endStation].total)/float64(this.src_des[startStation][endStation].n)
}

func main () { }