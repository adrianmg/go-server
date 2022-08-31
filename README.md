# http server in go

A tiny HTTP Go server for learning purposes. It's got a single endpoint:

- `/`: returns the current downloads in JSON format for https://adrianmato.art

Demo available on [adrianmg-go-server.fly.dev](https://adrianmg-go-server.fly.dev).

## dev

Use `go run .` or `go run index.go` to run locally on `localhost:8080`.

Deploy it for free on [Fly.io](https://fly.io) using `flyctl launch` and `flyctl deploy`. Alternatively, you can use [Render](https://render.com) to deploy it.
