package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/lucas59356/gedis/core"
)

var keys = make(chan (string))

func main() {
	times := flag.Int("n", 100, "Number of times of each iteration")
	flag.Parse()
	th := core.NewThread()
	for name, bench := range benchs {
		ti := time.Now()
		ta := bench(*times, th)
		tf := time.Now()
		fmt.Printf("Benchmark: %s: %v\n", name, tf.Sub(ti))
		println(ta.Show(name))
	}
	time.Sleep(100 * time.Second)
}

var benchs = map[string]func(int, *core.Thread) TimeADM{
	"benchSetString": func(nkeys int, th *core.Thread) TimeADM {
		keys = make(chan (string))
		// defer close(keys)
		go generateKeys(nkeys)
		time.Sleep(1 * time.Second)
		tadm := TimeADM{}
		for n := 1; n <= nkeys; n++ {
			select {
			case key := <-keys:
				ti := time.Now()
				_, _, err := th.Set(key, "Teste")
				tadm.Add(time.Now().Sub(ti))
				if err != nil {
					panic(err)
				}
			}
		}
		return tadm
	},
	"benchGetString": func(nkeys int, th *core.Thread) TimeADM {
		keys = make(chan (string))
		//	defer close(keys)
		go generateKeys(nkeys)
		time.Sleep(1 * time.Second)
		tadm := TimeADM{}
		for n := 1; n <= nkeys; n++ {
			select {
			case key := <-keys:
				ti := time.Now()
				_, _, err := th.Get(key)
				tadm.Add(time.Now().Sub(ti))
				if err != nil {
					println(err.Error())
				}
			}
		}
		return tadm
	},
	"benchDelString": func(nkeys int, th *core.Thread) TimeADM {
		keys = make(chan (string))
		go generateKeys(nkeys)
		time.Sleep(1 * time.Second)
		tadm := TimeADM{}
		for n := 1; n <= nkeys; n++ {
			select {
			case key := <-keys:
				ti := time.Now()
				err := th.Del(key)
				tadm.Add(time.Now().Sub(ti))
				if err != nil {
					panic(err)
				}
			}
		}
		return tadm
	},
}

func generateKeys(times int) {
	for i := 1; i <= times; i++ {
		keys <- fmt.Sprintf("Value%d", i)
	}
}
