package main

import (
	"fmt"

	"job4j.ru/go-lang-base/internal/base"
)

func main() {
	first := 100
	second := 10
	res := base.Add(second, first)

	fmt.Println(fmt.Sprintf("%d + %d = %d\n", first, second, res))

	for {
		fmt.Println(1)
		break
	}

	nums := []int{1, 1, 1, 3}
	fmt.Println(base.Mono(nums))
}
