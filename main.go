package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jasontconnell/fizzbuzz/process"
)

func main() {
	start := time.Now()
	rnglow, rnghigh, webmode := parseArgs(os.Args)

	if !webmode {
		stg := process.GetRules(process.GetEstHour())
		out := process.RunStage(rnglow, rnghigh, stg)
		for _, s := range out {
			fmt.Println(s)
		}
		log.Println("finished", time.Since(start))
	} else {
		log.Println("starting listening on :5678")
		startWeb()
	}
}

func startWeb() {
	h := func(resp http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			q := req.URL.Query()
			start, err := strconv.Atoi(q.Get("start"))
			if err != nil {
				start = 1
			}

			end, err := strconv.Atoi(q.Get("end"))
			if err != nil {
				end = 100
			}

			stg := process.GetRules(process.GetEstHour())
			output := process.RunStage(start, end, stg)

			enc := json.NewEncoder(resp)
			enc.Encode(output)
		}
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/fizzbuzz", h)
	mux.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":5678", mux)
}

func parseArgs(args []string) (int, int, bool) {
	defStart, defEnd := 1, 100
	web := false
	if len(args) == 3 {
		defStart, _ = strconv.Atoi(args[1])
		defEnd, _ = strconv.Atoi(args[2])

		if defStart <= 0 {
			defStart = 1
		}

		if defEnd > 100 {
			defEnd = 100
		}
	} else if len(args) == 2 && args[1] == "web" {
		web = true
	}
	return defStart, defEnd, web
}
