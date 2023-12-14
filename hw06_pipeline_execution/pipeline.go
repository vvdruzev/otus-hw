package hw06pipelineexecution

import (
	"log"
	"sync"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here.
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, stage := range stages {
			in = stage(in)
		}
		for {
			select {
			case <-done:
				//close(inn)
				log.Println("keeping = false")
				return
			case _, ok := <-in:
				if !ok {
					return
				}
				//log.Printf("sdf %v\n", v)
				//inn <- v
			}
		}
	}()
	wg.Wait()
	return in
}
