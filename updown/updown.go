package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MakeRandomNumber(n int) int{
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(n)
	return answer
}


func main(){
	userAnswer := 0
	tryngCount := 0
	answer := MakeRandomNumber(100)

	for true{
		if tryngCount == 10 {
			fmt.Println("정답은 :", answer, "입니다.", "시도한 횟수 :", tryngCount)
			break
		}

		if userAnswer == answer {
			fmt.Println("정답입니다 ! 정답은 :", answer, "입니다.", "시도한 횟수 :", tryngCount)
			break
		} else {
			fmt.Scanln(&userAnswer)
			tryngCount++
			if userAnswer > answer {
				fmt.Println("DOWN !!", answer)
				continue
			} else {
				fmt.Println("UP !!", answer)
				continue
			}
		}
	}
}
