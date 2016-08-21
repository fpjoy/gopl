package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	if len(s) < 3 {
		return s
	}

	for i, v := range s {
		if i%3 == 2 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%c", v)
	}
	return buf.String()
}
