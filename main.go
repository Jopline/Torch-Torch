package main

import (
	"fmt"
	"strings"
)

func less(a, b []string) bool {
	la, lb, min := len(a), len(b), 0
	if la == 0 && lb == 0 {
		return false
	} else if la == 0 && lb > 0 {
		return true
	} else if la > 0 && lb == 0 {
		return false
	}

	if min = la; la > lb {
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

// {[]string{"a", "b", "c"}, []string{"a", "b", "c"}, false},
func strSlice(a []string) string {
	var sb strings.Builder
	sb.WriteString(`[]string{`)

	for _, v := range a {
		sb.WriteString(fmt.Sprintf("%q,", v))
	}

	sb.WriteString(`}`)
	return sb.String()
}

func main() {
	ll := [][]string{
		{"a", "b", "c"},
		{"a", "b"},
		{"a", "a"},
		{"a", "c"},
		{"b"},
		{"b", "a"},
		{"c"},
		{""},
		{},
	}
	for i := range ll {
		for j := range ll {
			var op string
			if less(ll[i], ll[j]) {
				op = "true"
			} else {
				op = "false"
			}

			// fmt.Printf("%v %s %v\n", ll[i], op, ll[j])
			fmt.Printf("{%s, %s, %s},\n", strSlice(ll[i]), strSlice(ll[j]), op)
		}
	}
}
