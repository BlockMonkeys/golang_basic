package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	var nums [45] int
	var randomAry [6] int

	//45개 칸의 배열에 1 ~ 45까지의 숫자 삽입.
	for i:=0; i<len(nums); i++{
		nums[i] = i+1
	}

	//랜덤한 숫자 배열 6가지를 뽑는 반복문
	for i:=0; i<len(nums); i++{
		//랜덤한 숫자를 생성한다.
		rand.Seed(time.Now().UnixNano())
		randomNum := rand.Intn(len(nums))

		//만약 RandomAry의 마지막 값이 0이 아니라면 반복문을 종료한다.
		if randomAry[len(randomAry)-1] != 0 {
			break
		}

		//만약 nums Ary에서 뽑아온 값이 0이라면, i값을 --해서 현재값으로 유지한채 반복문을 다시 돌고,
		//nums Ary에서 뽑아온 값이 유효하다면 , 값을 뽑아서 randomAry에 삽입하고, nums배열에서는 0으로 바꿔준다.
		//Cuz, Golang의 Array는 고정길이로 생성하기 때문에 지우는 것 대신에 0으로 바꿔줬다.
		if nums[randomNum] != 0 {
			randomAry[i] = nums[randomNum]
			nums[randomNum] = 0
		} else {
			i--
		}
	}
	fmt.Println(randomAry)
}