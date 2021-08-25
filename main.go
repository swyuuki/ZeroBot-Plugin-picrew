package main

import (
	"cs/picrew"
	"fmt"
)

func main() {
		var url string
		if rand.Intn(3) == 0 {
			url = picrew.ToPic0()
		} else if rand.Intn(2) == 0 {
			url = picrew.ToPic1()
		} else {
			url = picrew.ToPic2()
		}

	fmt.Print(url)
}
