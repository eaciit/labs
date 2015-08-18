package main

import (
	"fmt"
	tk "github.com/eaciit/toolkit"
	"runtime"
	"time"
)

type ParallelResult struct {
	Data    []interface{}
	Success int
	Fail    int
	Errors  []string
}

func NewParallelResult() *ParallelResult {
	return &ParallelResult{
		Data:   []interface{}{},
		Errors: []string{}}
}

func RunParallel(keys []interface{}, f func(j <-chan interface{}, r chan<- tk.Result), workercount int) *ParallelResult {
	r := NewParallelResult()
	NumOfWork := len(keys)
	jobKeys := make(chan interface{}, NumOfWork)
	jobResult := make(chan tk.Result, NumOfWork)

	//--- pool the works
	for poolCount := 0; poolCount < workercount; poolCount++ {
		go f(jobKeys, jobResult)
	}

	//--- setting the key for work
	for keyId := 0; keyId < NumOfWork; keyId++ {
		jobKeys <- keys[keyId]
	}

	//--- collect the process
	for resultId := 0; resultId < NumOfWork; resultId++ {
		resultProcess := <-jobResult
		if resultProcess.Status == tk.Status_OK {
			r.Success++
			r.Data = append(r.Data, resultProcess.Data)
		} else {
			r.Fail++
			r.Errors = append(r.Errors, resultProcess.Message)
		}
	}

	return r
}
