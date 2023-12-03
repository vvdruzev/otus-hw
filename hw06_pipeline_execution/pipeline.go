package hw06pipelineexecution

import "log"

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here.
	inn := make(chan interface{})
	st := stages[0](inn)

	go func() {
		for _, stage := range stages[1:] {
			st = stage(st)
		}
	}()
	keeping := true
	for keeping {
		select {
		case <-done:
			close(inn)
			keeping = false
			log.Println("keeping = false")
			break
		case v, ok := <-in:
			if !ok {
				close(inn)
				keeping = false
				log.Println("keeping = false")

				break
			}
			log.Printf("sdf %v\n", v)
			inn <- v
		}
	}

	return st
}
