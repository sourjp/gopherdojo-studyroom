# 課題3-1 タイピングゲームを作ろう
## 要件の対応
* 標準出力に英単語を出す（出すものは自由）
csvファイルで出力する単語を[]stringで読み取り、randでindexを指定して出力するようにしました。

* 標準入力から1行受け取る
`goroutine`を利用して受け取るようにしました。

* 制限時間内に何問解けたか表示する
正解時に`correct`をカウントして最後に表示しています。

## テストプレイ
ゲーム時間は15秒です。

```bash
$ make start

        Thank you for playing TYPING Games!
        Let's you type word as you can see on display withing 15 seocnds!

        Start Game in 3 seconds...

3...
2...
1...
> orange
orange
Correct!
> peach
peach
Correct!
> avocade
avocde
Bad...
> grape
grape
Correct!
> orange
orange
Correct!
> apple
apple
Correct!
> yuzu
yuzu
Correct!
> plum
TimeUp!!!
You got 6 points!
```
