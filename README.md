# ServerSentEvent Server and Client sample

- server.go: ServerSentEventサーバー実装サンプル
- client.go: GopherJS用クライアントサンプル

server.goから
```go
broker.Push <- []byte("hoge")
```

とするとイベント待ちクライアント群に"hoge"がプッシュされる仕組み。

## test run

```sh
make build
make run
make open # other terminal
```
