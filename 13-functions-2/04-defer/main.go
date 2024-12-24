package main

import (
	"fmt"
	"time"
)

func main() {
	metrics := &Metrics{}

	err := Execute(func() error {
		fmt.Println("Executing...")
		return nil
	}, metrics)

	fmt.Println(err)
}

/*
* Fill in the Execute function.

The function should execute the given f function, and call proper methods on metrics:

Call StoreExecution() before calling the function.
Call StoreSuccess() if the function f returns without an error.
Call StoreFailure() if the function f returns with an error.
Use defer to handle the err return value.
*/

func Execute(f func() error, metrics *Metrics) (err error) {
	metrics.StoreExecution()
	defer func() {
		if err == nil {
			metrics.StoreSuccess()
		} else {
			metrics.StoreFailure()
		}
	}()
	err = f()
	return
}

type Metrics struct {
	execution []time.Time
	success   []time.Time
	failure   []time.Time
}

func (m *Metrics) StoreExecution() {
	m.execution = append(m.execution, time.Now())
}

func (m *Metrics) StoreSuccess() {
	m.success = append(m.success, time.Now())
}

func (m *Metrics) StoreFailure() {
	m.failure = append(m.failure, time.Now())
}
