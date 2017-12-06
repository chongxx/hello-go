package main

import (
	"fmt"
	"time"

	"github.com/chongxx/stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse("Hello, world.\n"))
	fmt.Printf("The time is ", time.Now())
}
