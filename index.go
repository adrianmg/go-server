package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
  http.HandleFunc("/", getTime)
  http.HandleFunc("/ip", getIp)

  http.ListenAndServe(getServerPort(), nil)
}

func getTime(res http.ResponseWriter, req *http.Request) {
  now := time.Now()

  fmt.Fprintf(res, "The time is %s\n", now.Format("3:04:05 PM"))
}

func getIp(res http.ResponseWriter, req *http.Request) {
  ip := req.RemoteAddr

  fmt.Fprintf(res, "Your IP is %s\n", ip)
}

func getServerPort() string {
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  return ":" + port
}