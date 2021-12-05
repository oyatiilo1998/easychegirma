package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Add(time.Duration(+5) * time.Hour)

	fmt.Println("now:", now)

}
