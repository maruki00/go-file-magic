package main

import (
	"fmt"
	"strings"
)

func main() {
	b := []byte{0x01, 0x54, 0x71, 0x67, 255}

	var out strings.Builder
	for _, b := range b {
		out.WriteString(fmt.Sprintf("\\x%.2x", b))
	}
	fmt.Println(out.String())

}
