package janken

import (
	"fmt"
	"math/rand"
	"time"
)

func get_cpu() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3)
}

func judge_win_lose(user int, cpu int) int {

	if user == 0 && cpu == 0 { // user: グー, cpu:グー
		return 2
	} else if user == 0 && cpu == 1 { // user: グー, cpu:チョキ
		return 0
	} else if user == 0 && cpu == 2 { // user: グー, cpu:パー
		return 1
	} else if user == 1 && cpu == 0 { // user: チョキ, cpu:グー
		return 1
	} else if user == 1 && cpu == 1 { // user: チョキ, cpu:チョキ
		return 2
	} else if user == 1 && cpu == 2 { // user: チョキ, cpu:パー
		return 0
	} else if user == 2 && cpu == 0 { // user: パー, cpu:グー
		return 0
	} else if user == 2 && cpu == 1 { // user: パー, cpu:チョキ
		return 1
	} else if user == 2 && cpu == 2 { // user: パー, cpu:パー
		return 2
	} else { // 予期しない入力の時
		return -1
	}
}

func janken_loop() {
	var user int
	var cpu int
	var userwin int
	var cpuwin int
	var draw int

	user = 0 //strで受け取ってintに直す関数
	cpu = 1  //いったん決め打ち
	for {
		tmp = judge_win_lose(user, cpu)
		switch tmp {
		case 0:
			userwin += 1
		case 1:
			cpuwin += 1
		case 2:
			draw += 1
		}
		fmt.Printf("User win:%d lose:%d draw:%d\n", userwin, cpuwin, draw)
	}

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

func get_user(input string) int {
	var input_int int

	switch input {
	case "r":
		input_int = 0
	case "s":
		input_int = 1
	case "p":
		input_int = 2
	}

	return input_int
}

func judge(userwin int, cpuwin int, draw int) int {
	var win_lose int

	if userwin > cpuwin {
		fmt.Println("User win")
		fmt.Printf("win:%d lose:%d draw:%d\n", userwin, cpuwin, draw)
		win_lose = 0
	} else if cpuwin > userwin {
		fmt.Println("Cpu win")
		fmt.Printf("win:%d lose:%d draw:%d\n", userwin, cpuwin, draw)
		win_lose = 1
	} else {
		fmt.Println("Draw")
		fmt.Printf("user_win:%d cpu_win:%d draw:%d\n", userwin, cpuwin, draw)
		win_lose = 2
	}

	return win_lose
}

func main() {
	c := get_cpu()
	fmt.Println(c)
	fmt.Println("user input: rock = r, scissors = s, paper = p")
	janken_loop()
}
