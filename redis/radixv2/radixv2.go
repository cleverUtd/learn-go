package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"time"

	"io"

	"log"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"

	_ "net/http/pprof"
)

const size = 300
const redisURI = "10.0.0.29:6671"

var (
	redisPool *pool.Pool
)

func init() {
	p, err := pool.New("tcp", redisURI, size)
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
	_, err := redisDo(redisPool, "GET", "foo")
	if err != nil {
		fmt.Println("redis error.", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// runtime.GOMAXPROCS(8)
	flag.Parse()

	numGroutines()

	http.HandleFunc("/redisget", cmdHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// pprof

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func cpuprof() {
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
}

func memprof() {
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}

func numGroutines() {
	go func() {
		http.HandleFunc("/goroutines", func(w http.ResponseWriter, r *http.Request) {
			num := strconv.FormatInt(int64(runtime.NumGoroutine()), 10)
			w.Write([]byte(num))
		})
		http.ListenAndServe(":9091", nil)
	}()
}
