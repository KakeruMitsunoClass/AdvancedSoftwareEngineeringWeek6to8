package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random_cpu(max_num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max_num)
}

func print_hand(cpu int, user int) {
	// xx vs xx
	var scpu string
	var suser string
	switch cpu {
	case 0:
		scpu = "rock"
	case 1:
		scpu = "scissors"
	case 2:
		scpu = "paper"
	}
	switch user {
	case 0:
		suser = "rock"
	case 1:
		suser = "scissors"
	case 2:
		suser = "paper"
	}
	fmt.Println("User", suser, "VS", "Cpu", scpu)
}

func print_cpu_achimuitehoi(input int) {
	switch input {
	case 0:
		fmt.Println("The CPU's direction is up")
	case 1:
		fmt.Println("The CPU's direction is down")
	case 2:
		fmt.Println("The CPU's direction is right")
	case 3:
		fmt.Println("The CPU's direction is left")
	default:
		// nothing
	}
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

func judge_win_lose_achimuitehoi(user int, cpu int) int {

	if user == cpu {
		return 0
	} else {
		return 1
	}
}

func janken_loop() (int, int, int) {
	var user int
	var cpu int
	var userwin int
	var cpuwin int
	var draw int
	var suser string
	var user_achi int
	var cpu_achi int
	count := 0

	for {
		count += 1
		fmt.Println("round ", count)
		fmt.Println("Type \"exit\" if you want to quit.")
		fmt.Println("Please enter the user's hand")

		//じゃんけん
		fmt.Scanf("%s", &suser)

		user = get_user(suser) //strで受け取ってintに直す関数
		if user == -1 {
			break
		}

		cpu = random_cpu(3) //ランダム生成関数
		print_hand(cpu, user)

		janken_judge := judge_win_lose(user, cpu)
		switch janken_judge {
		case 0:
			fmt.Println("User janken win")
		case 1:
			fmt.Println("Cpu janken win")
		case 2:
			fmt.Println("janken: Draw")
			fmt.Println("----------------------------------------------")
			continue
		}

		//あっち向いてホイの部分
		fmt.Println("Please enter the user's direction ")
		fmt.Println("\"r\"ight, \"l\"eft, \"d\"own \"u\"p")

		fmt.Scanf("%s", &suser)
		user_achi = get_user_achimuitehoi(suser) //strで受け取ってintに直す関数
		if user_achi == -1 {
			break
		}
		cpu_achi = random_cpu(4) //ランダム生成関数
		print_cpu_achimuitehoi(cpu_achi)
		// あっち向いてホイ: 勝ち負け判定
		tmp := judge_win_lose(user, cpu)                             //ジャンケン勝ち負け
		tmp_achi := judge_win_lose_achimuitehoi(user_achi, cpu_achi) //同じ方向:0,違う方向:1
		if tmp == 0 && tmp_achi == 0 {
			userwin += 1
			fmt.Println("User Achimuitehoi win")
		} else if tmp == 1 && tmp_achi == 0 {
			cpuwin += 1
			fmt.Println("Cpu Achimuitehoi win")
		} else {
			draw += 1
			fmt.Println("Draw")
		}
		fmt.Println("----------------------------------------------")
	}

	_ = judge(userwin, cpuwin, draw)

	return userwin, cpuwin, draw
}

func get_user(input string) int {
	var input_int int
	var suser string

	switch input {
	case "r":
		input_int = 0
	case "s":
		input_int = 1
	case "p":
		input_int = 2
	case "exit":
		input_int = -1
	default:
		fmt.Println("input correct hand")
		fmt.Scanf("%s", &suser)
		input_int = get_user(suser)
	}
	return input_int
}

func get_user_achimuitehoi(input string) int {
	var input_int int
	var suser string

	switch input {
	case "u":
		input_int = 0
	case "d":
		input_int = 1
	case "r":
		input_int = 2
	case "l":
		input_int = 3
	case "exit":
		input_int = -1
	default:
		fmt.Println("input correct direction")
		fmt.Scanf("%s", &suser)
		input_int = get_user_achimuitehoi(suser)

	}

	return input_int
}

func judge(userwin int, cpuwin int, draw int) int {
	// testのために誰が勝ったかを返す
	var win_lose int

	if userwin > cpuwin {
		fmt.Println("User win")
		fmt.Printf("win:%d, lose:%d, draw:%d\n", userwin, cpuwin, draw)
		win_lose = 0
	} else if cpuwin > userwin {
		fmt.Println("Cpu win")
		fmt.Printf("win:%d, lose:%d, draw:%d\n", userwin, cpuwin, draw)
		win_lose = 1
	} else {
		fmt.Println("Draw")
		fmt.Printf("user_win:%d, cpu_win:%d, draw:%d\n", userwin, cpuwin, draw)
		win_lose = 2
	}
	return win_lose
}

func main() {
	c := random_cpu(3)
	fmt.Println(c)
	fmt.Println("user input: rock = r, scissors = s, paper = p")
	janken_loop()
}
