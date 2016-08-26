package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	loop int
	t    time.Duration
)

func init() {
	flag.IntVar(&loop, "n", 1, "how many times? if < 0 infinite")
	flag.DurationVar(&t, "t", time.Second, "sleep time")
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatalln("exec command required")
	}

	name := flag.Arg(0)
	arg := flag.Args()[1:]
	for i := 0; i < loop || loop < 0; i++ {
		cmd := exec.Command(name, arg...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
		time.Sleep(t)
	}
}
