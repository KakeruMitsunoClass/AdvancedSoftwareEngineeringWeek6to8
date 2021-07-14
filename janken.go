package janken

import (
	"fmt"
)

func judge_win_lose(user int, cpu int) int {
	if user == cpu { //あいこ
		return 2
	} else if user == 0 && cpu == 1 { // user: グー, cpu:チョキ
		return 0
	} else if user == 0 && cpu == 2 { // user: グー, cpu:パー
		return 1
	} else if user == 1 && cpu == 0 { // user: チョキ, cpu:グー
		return 1
	} else if user == 1 && cpu == 2 { // user: チョキ, cpu:パー
		return 0
	} else if user == 2 && cpu == 0 { // user: パー, cpu:グー
		return 0
	} else if user == 2 && cpu == 1 { // user: パー, cpu:チョキ
		return 1
	} else {
		return 2
	}
}

func janken_loop() (int, int, int) {
	var user int
	var cpu int
	var userwin int
	var cpuwin int
	var draw int
	var suser string
	var result int

	for {
		fmt.Scanf("%s", &suser)
		if suser == "exit" {
			break
		}
		user = get_user(suser) //strで受け取ってintに直す関数
		cpu = get_cpu()        //ランダム生成関数
		tmp = judge_win_lose(user, cpu)
		switch tmp {
		case 0:
			userwin += 1
			fmt.Println("User win")
		case 1:
			cpuwin += 1
			fmt.Println("Cpu win")
		case 2:
			draw += 1
			fmt.Println("Draw")
		}
	}

	result = judge(userwin, cpuwin, draw)

	return userwin, cpuwin, draw
}

func main() {
	fmt.Println("user input: rock = r, scissors = s, paper = p")
	// janken_loop()
}
