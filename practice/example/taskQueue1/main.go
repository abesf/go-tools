package main

import (
	"fmt"
	"sync"
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
		if i==69{
			close(jobChan)
			break
		}else{
			jobChan<-Job{i}
		}
	}
	wg.Wait()
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