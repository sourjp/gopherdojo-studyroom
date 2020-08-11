# conv - image converter(課題1:画像変換コマンド)

# 仕様
## ディレクトリの指定と、再帰的処理、JPGファイル->PNG変換（デフォルト）

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
Image convert was suceeded!

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

## 画像形式の指定（オプション）
`jpg,` `jpeg`, `png`, `gif`の相互変換に対応

```bash
$ ./converter --help
Usage of ./conv:
  -de string
        destinate extension to encode(jpg, jpeg, png, gif) (default "png")
  -se string
        source exetention to decode(jpg, jpeg, png, gif) (default "jpg")
```

## その他の対応
| 要件 | 対応内容 |
| --- | --- |
| mainパッケージと分離 | package convを作成 |
| 自作・標準・準標準パッケージのみ | 遵守 |
| ユーザー定義型 | Converterを作成 |
| GoDoc | godocで自作パッケージが見えることを確認 |
| Go Modules | go.mod, go.sumの作成 |
