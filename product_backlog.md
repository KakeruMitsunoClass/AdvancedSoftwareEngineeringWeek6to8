# Product Backlog
1. じゃんけんの出す手を入力する
2. cpuじゃんけんの手を出力する
3. 勝ち負けを判定する
4. 勝ち負けを保存する
5. ユーザーがexitするまでループする

# Sprint 1
## Sprint Backlog
1. じゃんけん関数を作る(入力: 二つのじゃんけんの手, 出力:どっちが勝ったか): judge_win_lose()
2. じゃんけんの入力，CPUの入力，勝ち負け保存関数(, 出力: 対戦結果): janken_loop()
3. main関数を作る(対戦結果を保存する): main()
4. テストファイル作成: janken_test()
## メモ
中身はまだ開発しない
大枠だけ
適当に返す
0: user
1: cpu
2: あいこ
r: グー (0)
s: チョキ (1)
p: パー (2)
## sprint review
ユーザーの入力を受ける関数
cpuのランダム関数
janken_loop()にreturn 追加
最終評価をする関数
judge_win_lose関数最後エラーを返す
testにパッケージ追加

# Sprint 2
## Sprint Backlog
1. ユーザーの入力を受けて，数字に変換して返す関数
2. cpuの手を生成する関数
3. janken_loopの修正: ユーザーの手決定関数の追加，cpuの手を決める関数の追加, sprint1で決まった修正
4. judge_win_lose(): sprint1で決まった修正
5. 最終評価をする関数の作成
6. test: sprint1で決まった修正
## メモ

# Sprint 3
## Sprint Backlog
1. 実行できるようにデバッグする
2. 動作を確認する
3. 動作がおかしいところを直す
4. cpuの手を表示する
5. 終了の仕方を表示する

# Sprint 4
## Sprint Backlog
1. 何回目の勝負か表示する
2. 入力に対するエラー処理をする
3. 勝負の手をわかりやすく表示する(ex. rock vs paper)
4. あっち向いてホイの入力の受け取り

# Sprint 5
## Sprint Backlog
1. 勝負の手をわかりやすく表示する(ex. rock vs paper)
2. じゃんけんの勝敗の出力
3. あっち向いてホイの入力の受け取り
4. あっちむいてほいの勝敗の出力

# Sprint 6
## Sprint Backlog
1. 表示をわかりやすく変更
2. コードのリファクタリング
3. 