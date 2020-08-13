# 課題2: テスト, io.Readerとio.Writer
課題1作成物を元にテストを踏まえてリファクタ

## テスト
### 要件について
* テストのしやすさを考えてリファクタリングする  
main.goを対象に、課題1ではflag関係の処理と、画像変換処理を一つのRun()にまとめていた。  
課題2ではflag関係の処理をRun()にし、その初期化の引数を受け取るConverter()に分けて画像変換処理をテストできるようにした。

* テストのカバレッジを取る  
package conv(cover率 85.7%), package main(cover率 46.2%)

```bash
$ make test | grep cov
coverage: 85.7% of statements
ok      github.com/sourjp/gopherdojo-studyroom/kadai2   0.139s  coverage: 85.7% of statements

coverage: 46.2% of statements
ok      github.com/sourjp/gopherdojo-studyroom/kadai2/cmd/conv  0.142s  coverage: 46.2% of statements
```

* テーブル駆動テストを行う  
対応した

* テストヘルパーを作る  
testDecodeAndEncode()としてヘルパー関数を用意し、t.Helper()を適用し動作の違いを確認した

### 基本操作
```bash
$ tree  testdata
testdata
├── t1.jpg
└── testdata2
    └── t2.jpg

1 directory, 2 files

$ file testdata/t1.jpg 
testdata/t1.jpg: JPEG image data, baseline, precision 8, 1431x901, frames 3

$ make build
$ ./converter testdata
Image convert has finished!

$ tree testdata
testdata
├── t1.jpg
├── t1.png
└── testdata2
    ├── t2.jpg
    └── t2.png

1 directory, 4 files

$ file testdata/t1.png
testdata/t1.png: PNG image data, 1431 x 901, 16-bit/color RGB, non-interlaced
```

## io.Readerとio.Writer
### io packageとは
* I/Oプリミティブへの基本的なインターフェースを提供
* 主な仕事は、package osのようなプリミティブな実装に、機能を抽象化する共有パブリックインタフェースとしてラップすることで、関連プリミティブになるようにすること
* これらのインタフェースとプリミイティブは低レベルな命令を様々な実装でラップするため、知らされない限り並行実行が安全と想定しない方が良い

参考: [Overview](https://golang.org/pkg/io/#Overview)

### io.Readerとは
* Readメソッドをラップするインタフェース
* 基礎となるデータストリームから引数pに読み込んだ結果を書き込む
* 返り値は読み込めたbyte数 0 <= n <= len(p)と、読み込み時のerrを返す
* そもそものデータのbyte数が少なければ、pで与えたsizeとは小さい場合があるため n == len(p)とは限らない

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

参考: [type Reader](https://golang.org/pkg/io/#Reader)

### io.Writerとは
* Writeメソッドをラップするインタフェース
* 基礎となるデータストリームに引数pを書き込む
* 返り値は書き込めたbyte数 0 <= n <= len(p)と、書き込み停止時のerrを返す
* n < len(p)ならerrは非nilのためerrをとれば書き込みが確認できる

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

参考: [type Writer](https://golang.org/pkg/io/#Writer)

### `io.Reader`と`io.Writer`による利点は？
参考の例が非常にわかりやすかった。
次のような関数がある場合に、引数としてFileでもReaderでも`io.Read`の実装を満たすため様々なオブジェクトを渡すことができるため、関数を変更する必要がない。

* `func (f *File) Read(b []byte) (n int, err error) {...}` 
* `func (b *Reader) Read(p []byte) (n int, err error) {...}`

```go
func readbyte(r io.Reader) (rune, error) {
    n, err := r.Read(p[:])
    if n > 0 {
        return rune(p[0]), nil
    }
    return 0, err
}
```

これらのことから、`type interface`を用いて`io.Reader`や`io.Writer`と抽象化したことで、様々なオブジェクトを同じように扱うことが可能になった。

参考：　[io.Readerをすこれ](https://qiita.com/ktnyt/items/8ede94469ba8b1399b12)
