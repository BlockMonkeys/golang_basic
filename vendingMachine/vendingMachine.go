package main

import "fmt"


func main(){
	var userMoney int
	var userSelect string

	beverages := map[string] int{
		"코카콜라" : 800,
		"환타" : 500,
		"제티" : 700,
	}
	for key, val := range beverages {
		fmt.Println("목록을 출력합니다 : ", key, val)
	}

	//사용자로 부터 돈을 입력받는다.
	fmt.Println("돈을 넣어주세요.")
	fmt.Scan(&userMoney)

	//입력받은 돈에 맞는 목록을 출력한다.
	fmt.Println("무엇을 구매하시겠나요? 음료 이름을 입력해주세요...")
	for key, val := range beverages {
		if userMoney >= val {
			fmt.Print(key, ":", val)
		}
	}

	//사용자 입력값에 맞는 것을 탐색한다.
	fmt.Scan(&userSelect)
	for key, val := range beverages {
		if key == userSelect {
			fmt.Println("요청하신 음료수", userSelect, "가 나왔습니다.")
			//거스름돈 반환
			var change int
			change = userMoney - val
			if change == 0 {
				fmt.Println("거스름돈은 없습니다.")
			} else {
				fmt.Println("거스름돈은 ", change, "원 입니다. 반환되었습니다.")
			}
		}
	}
	fmt.Println("거래가 종료되었습니다. 감사합니다 😆")
}