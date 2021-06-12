# split-slice-go

## 概要

スライスを複数のスライスに分割する。
このとき分割後の各スライスの要素の“長さの合計”をできるだけ均等にする。

### 用例

```
The sun shines bright in the old Kentucky home
```
を空白の一で折り返し、3行でボックスに表示したい(等幅フォントを仮定)。
このときボックスの幅を可能な限り狭くしたい。

だめな解:
```
The sun shines bright in the
old
Kentucky home
```

ベストの解(この文字列の場合は2個ある):
```
The sun shines
bright in the
old Kentucky home
```
```
The sun shines
bright in the old
Kentucky home
```

## 使い方

```go
import split "github.com/pen/split-slice-go"

    // :

result := split.Sentence(
    "I have a pen. I have an apple.",  // 分割したい文字列
    2,                                 // 分割数
    false,                             // 長い行を早めに出すか?
)
```
3番目の引数は、前節の例において2つの解のどちらを返すかを指定することにあたる。

これにより結果が大きく異る場合もある。
```go
package main

import (
    "fmt"

    split "github.com/pen/split-slice-go"
)

func main() {
    sentence := `Are hackers a threat?` +
        ` The degree of threat presented by any conduct,` +
        ` whether legal or illegal,` +
        ` depends on the actions and` +
        ` intent of the individual and the harm they cause.`

    fmt.Printf("%s\n", split.Sentence(sentence, 7, false))
    fmt.Println("--------------------------")
    fmt.Printf("%s\n", split.Sentence(sentence, 7, true))
}
```
```console
$ go run sample.go 
Are hackers a threat?
The degree of threat
presented by any conduct,
whether legal or illegal,
depends on the actions and
intent of the individual
and the harm they cause.
--------------------------
Are hackers a threat? The
degree of threat presented
by any conduct, whether
legal or illegal, depends
on the actions and intent
of the individual and the
harm they cause.
$
```

### 整数のスライス

`IntSlice` は分割ポイントの入ったスライスを返すので、自力でもとのスライスを切り出す必要がある。

```go
origin := []int{3, 1, 4, 1, 5, 6, 5, 3, 5}

indice := split.IntSlice(origin, 3, false) //=> [0 3 6 9]
nPart := len(indice) - 1;

result := make([][]int, nPart)

// indiceを使ってもとのスライスから切り出す
for i := 0; i < nPart; i++ {
    result[i] = origin[indice[i]:indice[i+1]]
}

//=> [[3 1 4 1] [5 6] [5 3 5]]
```

### 一般化

`Slice` は []interface{} を分割する。要素からint型の“長さ”をとりだす関数を渡す必要がある。

```go
indice := split.Slice(
    slice,
    func(i int) int {
        // slice[i] の長さを返す
    },
    3, false
)
```

## アルゴリズム

我流。

この問題を効率よく解く既存のアルゴリズムやそれを実装したライブラリ(言語問わず)をご存知であれば教えていただけると幸いです。
