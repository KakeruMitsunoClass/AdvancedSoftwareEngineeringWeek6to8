
func janken_loop(){
    var user int
    var cpu int
    var userwin int
    var cpuwin int
    var draw int
    
    user = 0 //strで受け取ってintに直す関数
    cpu = 1 //いったん決め打ち
    for{
        tmp = judge_win_lose(user,cpu)
        switch tmp {
            case 0 :
                userwin += 1
            case 1 :
                cpuwin += 1
            case 2 :
                draw += 1
        }
        fmt.Printf("User win:%d lose:%d draw:%d\n",userwin,cpuwin,draw)
    }
}