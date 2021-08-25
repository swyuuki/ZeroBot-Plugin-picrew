package main

import (
	"cs/picrew"
	"fmt"
)

func main() {
	// for i, v := range picrew.Pic0["前髪"].Items[0][0] {
	// 	fmt.Print(i, ". ", v, " \n")
	// }

	fmt.Print(picrew.ToPic2())
}
