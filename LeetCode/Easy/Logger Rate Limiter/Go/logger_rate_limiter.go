package main

type Logger struct{ hash_map map[string]int }

func Constructor() Logger { return Logger{hash_map: make(map[string]int)} }

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if v, ok := this.hash_map[message]; !ok {
		this.hash_map[message] = 10 + timestamp
		return true
	} else {
		if v <= timestamp {
			this.hash_map[message] = 10 + timestamp
			return true
		} else {
			return false
		}
	}
}

func main() {

}
