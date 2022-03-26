# go-gle

google.golang.org/api/* のラッパ

## 使い方

とりあえず[_example/print-range.go](https://github.com/pen/go-gle/blob/main/_example/print-range.go)参照。


## 認証情報

環境変数に入れておくと自動的に使用する。

```shell
echo "GOOGLE_APY_KEY=\"$(cat json.key | xz -9 -c | base64)\"" >> .env
direnv allow
rm json.key
```
