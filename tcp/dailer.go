package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
)

type Worker struct {
	addr string
}

type Result struct {
	Port int
	Open bool
}

var (
	host string
	port string
)

func init() {
	flag.StringVar(&host, "host", "scanme.nmap.org", "target host")
	flag.StringVar(&port, "port", "80", "valid values: INT, INT-INT")
}

func main() {
	flag.Parse()

	start := 0
	end := 0
	var convErr error

	if strings.Contains(port, "-") {
		ranges := strings.SplitN(port, "-", 2)
		if len(ranges) != 2 {
			panic("please specify just one range 50-70")
		}

		start, convErr = strconv.Atoi(ranges[0])

		if convErr != nil {
			panic(convErr)
		}

		end, convErr = strconv.Atoi(ranges[1])

		if convErr != nil {
			panic(convErr)
		}

		if start >= end {
			panic("start must be slower than end in port range")
		}

	}

	singlePort := start == 0 && end == 0

	if singlePort {
		if convErr != nil {
			panic(convErr)
		}
		target := fmt.Sprintf("%s:%s", host, port)
		_, err := net.Dial("tcp", target)
		if err != nil {
			panic(err)
		}

		log.Println(target)
		return
	}

	var wg sync.WaitGroup
	pool := make(chan int, end-start)
	results := make(chan Result, end-start)

	go func() {
		for i := start; i <= end; i++ {
			pool <- i
		}
	}()

	// 4 workers
	for i := 0; i <= 30; i++ {
		w := Worker{host}
		go w.Run(pool, results, &wg)
		wg.Add(1)
	}

	outcomes := make([]Result, 0, end-start)
	for r := range results {
		outcomes = append(outcomes, r)
		if len(outcomes) == end-start {
			close(pool)
			break
		}
	}

	close(results)

	for _, item := range outcomes {
		fmt.Println(item)
	}

}

func (w *Worker) Run(ports chan int, results chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for p := range ports {
		target := fmt.Sprintf("%s:%d", w.addr, p)
		r := Result{}
		r.Port = p
		_, err := net.Dial("tcp", target)
		closed := err != nil

		if closed {
			fmt.Println(err)
		}
		r.Open = !closed
		results <- r
	}
}
