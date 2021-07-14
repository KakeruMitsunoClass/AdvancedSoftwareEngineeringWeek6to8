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