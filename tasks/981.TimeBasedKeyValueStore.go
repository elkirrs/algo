package tasks

import (
	"fmt"
)

func TimeBasedKeyValueStoreRun() {
	fmt.Println("981. Time Based Key-Value Store")
	fmt.Println("https://leetcode.com/problems/time-based-key-value-store/description/")
	fmt.Println("")

	operations := []string{
		"TimeMap",
		"set",
		"get",
		"get",
		"set",
		"get",
		"get",
	}

	arguments := [][]any{
		{},
		{"foo", "bar", 1},
		{"foo", 1},
		{"foo", 3},
		{"foo", "bar2", 4},
		{"foo", 4},
		{"foo", 5},
	}

	timeBasedKeyValueStorePrint(operations, arguments)
}

type TimeMap struct {
	key map[string][]Data
}

type Data struct {
	timestamp int
	value     string
}

func Constructor() TimeMap {
	return TimeMap{
		key: make(map[string][]Data),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.key[key] = append(this.key[key], Data{timestamp: timestamp, value: value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, exist := this.key[key]; !exist {
		return ""
	}

	data := this.key[key]

	l, r := 0, len(data)-1

	for l <= r {
		mid := l + (r-l)/2

		if data[mid].timestamp <= timestamp {
			if mid == len(data)-1 || data[mid+1].timestamp > timestamp {
				return data[mid].value
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ""
}

func timeBasedKeyValueStorePrint(operations []string, arguments [][]any) {
	fmt.Println("Input:")
	fmt.Println("Operations:", operations)
	fmt.Println("Arguments:", arguments)
	obj := Constructor()

	result := make([]any, 0)

	for i, op := range operations {
		switch op {
		case "set":
			obj.Set(
				arguments[i][0].(string),
				arguments[i][1].(string),
				arguments[i][2].(int),
			)

		case "get":
			out := obj.Get(
				arguments[i][0].(string),
				arguments[i][1].(int),
			)

			result = append(result, out)
		}
	}

	fmt.Println("Output:", result)
}
