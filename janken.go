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
	scpu = trans_hand(cpu)
	suser = trans_hand(user)
	fmt.Println("User", suser, "VS", "Cpu", scpu)
}

func trans_hand(input int) string {
	var str string

	switch input {
	case 0:
		str = "rock"
	case 1:
		str = "scissors"
	case 2:
		str = "paper"
	default:
		// nothing
	}

	return str
}

func trans_direction(input int) string {
	var str string

	switch input {
	case 0:
		str = "up"
	case 1:
		str = "down"
	case 2:
		str = "right"
	case 3:
		str = "left"
	default:
		// nothing
	}
	return str
}

func print_achimuitehoi(cpu int, user int) {
	var suser string
	var scpu string

	scpu = trans_direction(cpu)
	suser = trans_direction(user)

	fmt.Println("User", suser, "VS", "Cpu", scpu)
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
		fmt.Println("Round ", count)
		fmt.Println("Type \"q\" if you want to quit.")
		fmt.Println("---------------Janken-----------------")
		fmt.Println("\"r\"ock, \"s\"cissor, \"p\"aper")

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
			fmt.Println("Janken: You win")
		case 1:
			fmt.Println("Janken: You lose")
		case 2:
			fmt.Println("Janken: Draw")
			fmt.Println("----------------------------------------------")
			continue
		}

		//あっち向いてホイの部分
		fmt.Println("-------------Achimuitehoi-------------")
		fmt.Println("\"r\"ight, \"l\"eft, \"d\"own \"u\"p")

		fmt.Scanf("%s", &suser)
		user_achi = get_user_achimuitehoi(suser) //strで受け取ってintに直す関数
		if user_achi == -1 {
			break
		}
		cpu_achi = random_cpu(4) //ランダム生成関数
		print_achimuitehoi(cpu_achi, user_achi)
		// あっち向いてホイ: 勝ち負け判定                    //ジャンケン勝ち負け
		achi_judge := judge_win_lose_achimuitehoi(user_achi, cpu_achi) //同じ方向:0,違う方向:1
		if janken_judge == 0 && achi_judge == 0 {
			userwin += 1
			fmt.Println("Achimuitehoi: You win")
		} else if janken_judge == 1 && achi_judge == 0 {
			cpuwin += 1
			fmt.Println("Achimuitehoi: You lose")
		} else {
			draw += 1
			fmt.Println("Draw")
		}

		fmt.Println("----------------------------------------------")
		fmt.Printf("win:%d, lose:%d, draw:%d\n", userwin, cpuwin, draw)
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
	case "q":
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
	case "q":
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
	janken_loop()
}
