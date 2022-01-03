package grpshuffle_client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// HttpResponse is response http server for grpshuffle-client
type HttpResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

// HttpServe is serve http server for grpshuffle-client
func HttpServe(port int) {
	addr := fmt.Sprintf(":%v", port)

	http.HandleFunc("/shuffle", shuffleHandler)
	http.HandleFunc("/", healthHandler)

	log.Fatal(http.ListenAndServe(addr, logger(http.DefaultServeMux)))
}

func shuffleHandler(writer http.ResponseWriter, request *http.Request) {
	conn, err := Connect(Host, Port)
	if err != nil {
		newErrorResponse(writer, 500, "Internal Server Error")
		return
	}
	defer CloseConnect(conn)

	rawPartition := request.FormValue("partition")
	if rawPartition == "" {
		newErrorResponse(writer, 400, "partition parameter is required")
		return
	}
	partition, err := strconv.Atoi(rawPartition)
	if err != nil {
		newErrorResponse(writer, 400, "partition parameter allows numbers")
		return
	}

	rawTargets := request.FormValue("targets")
	if rawTargets == "" {
		newErrorResponse(writer, 400, "targets parameter is required")
		return
	}
	targets := strings.Split(rawTargets, ",")
	result, err := Shuffle(conn, int32(partition), targets)
	if err != nil {
		newErrorResponse(writer, 504, "Gateway Timeout")
		log.Print(err)
		return
	}

	res, err := json.Marshal(HttpResponse{
		Status: 200,
		Msg:    "Ok",
		Result: result,
	})
	if err != nil {
		log.Print(err)
		return
	}

	_, err = writer.Write(res)
	if err != nil {
		log.Print(err)
		return
	}
}

func healthHandler(writer http.ResponseWriter, _ *http.Request) {
	conn, err := Connect(Host, Port)
	if err != nil {
		newErrorResponse(writer, 500, "Internal Server Error")
		return
	}
	defer CloseConnect(conn)

	result, err := HealthCheck(conn)
	if err != nil {
		newErrorResponse(writer, 504, "Gateway Timeout")
		log.Print(err)
		return
	}

	res, err := json.Marshal(HttpResponse{
		Status: 200,
		Msg:    "Ok",
		Result: result,
	})
	if err != nil {
		log.Print(err)
		return
	}

	_, err = writer.Write(res)
	if err != nil {
		log.Print(err)
		return
	}
}

func newErrorResponse(writer http.ResponseWriter, statusCode int, msg string) {
	writer.WriteHeader(statusCode)
	res, err := json.Marshal(HttpResponse{
		Status: statusCode,
		Msg:    msg,
		Result: nil,
	})

	if err != nil {
		log.Print(err)
		return
	}

	_, err = writer.Write(res)
	if err != nil {
		log.Print(err)
		return
	}
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rAddr := r.RemoteAddr
		method := r.Method
		path := r.URL
		fmt.Printf("%v remote: %s [%s] %s\n",
			time.Now().Format("2006-01-02 15:04:05"),
			rAddr, method, path)
		h.ServeHTTP(w, r)
	})
}
