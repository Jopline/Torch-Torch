package main

import "fmt"

func less(a, b []string) bool {
	la, lb := len(a), len(b)
	if la == 0 && lb == 0 {
		return false
	} else if la == 0 && lb > 0 {
		return true
	} else if la > 0 && lb == 0 {
		return false
	}

	min := la
	if lb < la {
		min = lb
	}
	for i := 0; i < min; i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}
	if la == lb {
		return false
	}

	return la < lb
}

func main() {
	ll := [][]string{
		{"a", "b", "c"},
		{"a", "b"},
		{"b"},
		{"b", "a"},
		{"c"},
		{""},
	}
	for i := range ll {
		for j := range ll {
			var op string
			if less(ll[i], ll[j]) {
				op = "<"
			} else {
				op = "≥"
			}

			fmt.Printf("%v %s %v\n", ll[i], op, ll[j])
		}
	}
}
