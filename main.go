package main

import (
	"fmt"
	"log"
	"os"
	"pinfo/pkg"
	"strconv"
)

func main() {
	kernel, err := pkg.ParseKernel()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(kernel)

	if len(os.Args) != 2 {
		log.Fatal("参数错误！")
	}

	pid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	status, err := pkg.ParseStatus(pid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(status)

	others, err := pkg.ParseOther(pid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(others)

	io, err := pkg.ParseIO(pid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(io)

	limits, err := pkg.ParseLimits(pid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(limits)
}
