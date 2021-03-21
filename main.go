package main

import "fmt"

type Student struct {
	name string
}

func (s Student) calAvg(data []float64) (avgResult float64) {
	sum := 0.0

	for i := 0 ; i< len(data);i++{
		sum += data[i]
	}
	result := sum/float64(len(data))
	return result
}

func (s Student) judge(avg float64) (judgeResult string){
	if avg >= 60{
		judgeResult = "passed"
	} else{
		judgeResult = "failed"
	}
	return judgeResult
}

func main()  {
	s := Student{"eiji"}
	data := []float64{40,100,100,100,30}

	avg := s.calAvg(data)
	result := s.judge(avg)
	fmt.Println(result)
}