/*
Write a function to flatten a nested dictionary. Namespace the keys with a period.

For example, given the following dictionary:

{
    "key": 3,
    "foo": {
        "a": 5,
        "bar": {
            "baz": 8
        }
    }
}
it should become:

{
    "key": 3,
    "foo.a": 5,
    "foo.bar.baz": 8
}
You can assume keys do not contain dots in them, i.e. no clobbering will occur.
*/

package main

import (
	"fmt"
	"reflect"
)

func go_deep(mp map[string]interface{}, k string, res *[]map[string]int, str string) {
	val := mp[k]
	num := 0
	if reflect.TypeOf(val) == reflect.TypeOf(0) {
		return
	}
	str += k
	temp := mp[k].(map[string]interface{})
	for key, value := range temp {
		if reflect.TypeOf(value) == reflect.TypeOf(0) {
			num = value.(int)
			*res = append(*res, map[string]int{str + "." + key: num})
		} else {
			str += "."
			go_deep(temp, key, res, str)
		}
	}
}

func nested_dictionary(mp map[string]interface{}) map[string]int {
	res := make(map[string]int)
	res_ := make([]map[string]int, 0)
	for k, val := range mp {
		go_deep(mp, k, &res_, "")
		if len(res_) == 0 {
			res[k] = val.(int)
		} else {
			for i := 0; i < len(res_); i++ {
				for p, q := range res_[i] {
					res[p] = q
				}
			}
			res_ = nil
		}
	}
	return res
}

func main() {
	temp := make(map[string]interface{})
	temp2 := make(map[string]interface{})
	mp := make(map[string]interface{})
	mp["key"] = 3
	mp["foo"] = make(map[string]interface{})
	temp["a"] = 5
	temp["bar"] = make(map[string]interface{})
	temp2["baz"] = 8
	temp["bar"] = temp2
	mp["foo"] = temp
	fmt.Println(nested_dictionary(mp))
}
