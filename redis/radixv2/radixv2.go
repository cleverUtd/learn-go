package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"io"

	"log"

	"runtime"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

const size = 300
const redisURI = "10.0.0.29:6671"

var (
	redisPool *pool.Pool
)

func init() {
	p, err := pool.New("tcp", redisURI, 100)
	if err != nil {
		fmt.Println("error create pool.", err)
	}
	redisPool = p

	rand.Seed(time.Now().UnixNano())
}

func redisDo(p *pool.Pool, cmd string, args ...interface{}) (reply *redis.Resp, err error) {
	reply = p.Cmd(cmd, args...)
	if err = reply.Err; err != nil {
		if err != io.EOF {
			fmt.Println("redis", cmd, args, "err is", err)
		}
	}
	return reply, err
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	// key := randString(10)
	reply, err := redisDo(redisPool, "GET", "foo")
	if err != nil {
		fmt.Println("redis error.", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, reply.String())
}

func main() {
	runtime.GOMAXPROCS(8)
	http.HandleFunc("/redisCmd", cmdHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
