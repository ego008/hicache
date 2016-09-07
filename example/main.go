package main

import (
	"fmt"

	"github.com/ego008/hicache"
)

type setting struct {
	Title string
	Count int
	URL   string
}

func main() {
	c := hicache.New()

	c.Set("a", 222)
	v, ok := c.Get("a")
	fmt.Println(v, ok) // 222 true

	c.Del("a")
	fmt.Println(c.Get("a")) // <nil> false

	fmt.Println(c.Incr("b", 9))  // 9
	fmt.Println(c.Incr("b", 1))  // 10
	fmt.Println(c.Incr("b", -3)) // 7

	c.Set("c", "cvalue")
	fmt.Println(c.Incr("c", 1)) // 1

	st := setting{"test", 101, "http://www.google.com"}
	c.Set("d", st)
	v, ok = c.Get("d")
	if ok {
		fmt.Println(v)
		s := v.(setting)     // {test 101 http://www.google.com}
		fmt.Println(s.Title) // test
		fmt.Println(s.URL)   // http://www.google.com
	}

	fmt.Println(c.Count()) // 3

	c.Flush()
	fmt.Println(c.Count()) // 0
}
