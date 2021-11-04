package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	Id int `json:"id"`
}
func main()  {
	var wg sync.WaitGroup
	wg.Add(1)
	jobChan:=make(chan Job,100)
	go worker(jobChan,&wg)
	for i:=0;i<100;i++{
		jobChan<-Job{i}
		if i==99{
			close(jobChan)
			break
		}

	}
	//wg.Wait()
	//control exit time if time large then settime exit,if job finish before time limit just exit
	b:=WaitTimeout(&wg, 5*time.Second)
	fmt.Println(b)
}
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	ch := make(chan struct{})
	go func() {
		wg.Wait()
		close(ch)
	}()
	select {
	case <-ch:
		return true
	case <-time.After(timeout):
		return false
	}
}
func worker(jobChan <-chan Job,wg *sync.WaitGroup)  {
	defer wg.Done()
	for job:=range jobChan{
		process(job)
	}
}
func process(job Job)  {
	fmt.Printf("handle job which id is %+v\n", job.Id)
}