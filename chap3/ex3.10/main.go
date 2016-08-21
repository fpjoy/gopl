package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n < 3 {
		return s
	}

	for i, v := range s {
		if n%3 == 0 && i%3 == 0 && i != 0 {
			buf.WriteString(",")
		}

		if n%3 == 1 && i%3 == 1 {
			buf.WriteString(",")
		}

		if n%3 == 2 && i%3 == 2 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%c", v)
	}
	return buf.String()
}
