package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MakeRandNum(n int) int{
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(n)
	return answer+1
}

func GameStart(wallet, answer, userChoice int) int{
	for true{
		//Wallet 잔고확인
		if wallet < 100 {
			fmt.Println("잔액이 모자랍니다 !")
			break
		}

		//게임시작
		fmt.Println("1~5 사이의 원하는 숫자를 입력해주세요 ")
		fmt.Scan(&userChoice)
		if userChoice > 5 || userChoice < 1 {
			fmt.Println("❌ 정확한 숫자를 입력해주세요. 입력값은 1 ~ 5 입니다. ❌")

		} else {
			if userChoice == answer {
				wallet = wallet+500
				fmt.Println("🎉 축하합니다 !  500원이 추가됩니다. 🎉  💵 :", wallet)
				break
			} else {
				wallet = wallet-100
				fmt.Println("😭 틀렸습니다 ! 100원이 차감됩니다. 😭 💵 :", wallet)
				continue
			}
		}
	}
	return wallet
}


func main(){
	wallet := 1000
	var userChoice int
	answer := MakeRandNum(5)
	GameStart(wallet, answer, userChoice)

	fmt.Println("게임이 종료되었습니다. 당신의 잔고는 :",wallet,"입니다.")
}
