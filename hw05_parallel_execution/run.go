package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}

	counter := &Counter{
		mu:           sync.Mutex{},
		errCount:     0,
		m:            int64(m),
		tasksChannel: make(chan Task),
		done:         make(chan struct{}, n),
		err:          nil,
	}

	go func() {
		defer func() {
			close(counter.tasksChannel)
			close(counter.done)
		}()

		for _, task := range tasks {
			select {
			case <-counter.done:
				return
			case counter.tasksChannel <- task:
			}
		}
	}()
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			for task := range counter.tasksChannel {
				err := task()
				if err != nil {
					counter.little()
				}
			}
		}(i)
	}
	wg.Wait()

	return counter.err
}

type Counter struct {
	mu           sync.Mutex
	errCount     int64
	m            int64
	tasksChannel chan Task
	done         chan struct{}
	err          error
	once         sync.Once
}

func (c *Counter) little() {
	c.mu.Lock()
	c.errCount++
	if c.errCount >= c.m {
		c.once.Do(func() {
			c.done <- struct{}{}
		})
		c.err = ErrErrorsLimitExceeded
	}
	c.mu.Unlock()
}
