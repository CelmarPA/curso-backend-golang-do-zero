package main

import (
	"fmt"
	"calculator/math"

	"github.com/labstack/echo/v4"
)

func main () {
	// Echo instance
	e := echo.New()

	fmt.Println(e)

	sum := math.Sum(1, 2)
	fmt.Println(sum)

	sub := math.Sub(1, 2)
	fmt.Println(sub)
}
