package grpshuffle_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/korosuke613/grpshuffle/go/grpshuffle"
	"google.golang.org/grpc/health/grpc_health_v1"
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

type ResponseFormat string

const (
	FormatJson       = "json"
	FormatJsonPretty = "json-pretty"
	FormatReadable   = "readable"
)

func (r *ResponseFormat) checkType() error {
	switch *r {
	case FormatJson:
	case FormatJsonPretty:
	case FormatReadable:
	default:
		return fmt.Errorf("fmt parameter allows %v, %v and %v", FormatJson, FormatJsonPretty, FormatReadable)
	}
	return nil
}

// HttpServe is serve http server for grpshuffle-client
func HttpServe(port int) {
	addr := fmt.Sprintf(":%v", port)

	http.HandleFunc("/shuffle", shuffleHandler)
	http.HandleFunc("/", healthHandler)

	log.Printf("Launch grpshuffle-client http server %v...", port)
	log.Fatal(http.ListenAndServe(addr, logger(http.DefaultServeMux)))
}

func shuffleHandler(writer http.ResponseWriter, request *http.Request) {
	conn, err := Connect(Host, Port, NoTLS)
	if err != nil {
		newErrorResponse(writer, 500, "Internal Server Error")
		return
	}
	defer CloseConnect(conn)

	format := ResponseFormat(request.FormValue("fmt"))
	if format != "" {
		err = format.checkType()
		if err != nil {
			newErrorResponse(writer, 400, err.Error())
			return
		}
	}

	rawDivide := request.FormValue("divide")
	if rawDivide == "" {
		newErrorResponse(writer, 400, "divide parameter is required")
		return
	}
	divide, err := strconv.Atoi(rawDivide)
	if err != nil {
		newErrorResponse(writer, 400, "divide parameter allows numbers")
		return
	}

	groupName := request.FormValue("groupName")
	if groupName == "" {
		groupName = "Group"
	}

	rawTargets := request.FormValue("targets")
	if rawTargets == "" {
		newErrorResponse(writer, 400, "targets parameter is required")
		return
	}
	targets := strings.Split(rawTargets, ",")
	cc := grpshuffle.NewComputeClient(conn)
	result, err := Shuffle(&cc, uint64(divide), targets)
	if err != nil {
		newErrorResponse(writer, 500, fmt.Sprintf("[Internal Server Error] %s", err))
		return
	}

	var res []byte

	switch format {
	case FormatJson:
		{
			res, err = json.Marshal(HttpResponse{
				Status: 200,
				Msg:    "Ok",
				Result: result,
			})
			if err != nil {
				log.Print(err)
				return
			}
		}
	case FormatJsonPretty:
		{
			res, err = json.MarshalIndent(HttpResponse{
				Status: 200,
				Msg:    "Ok",
				Result: result,
			}, "", "  ")
			if err != nil {
				log.Print(err)
				return
			}
		}
	case FormatReadable:
		{
			var prettyResponse string
			for i, combination := range result {
				var targetsString string
				for _, target := range combination.Targets {
					if targetsString == "" {
						targetsString = target
					} else {
						targetsString = fmt.Sprintf("%v, %v", targetsString, target)
					}
				}
				groupNumber := i + 1 // It is for humans, so start at 1.
				prettyResponse = fmt.Sprintf("%v%v %v: %v\n", prettyResponse, groupName, groupNumber, targetsString)
			}
			res = []byte(prettyResponse)
		}
	default:
		{
			res, err = json.Marshal(HttpResponse{
				Status: 200,
				Msg:    "Ok",
				Result: result,
			})
			if err != nil {
				log.Print(err)
				return
			}
		}
	}

	_, err = writer.Write(res)
	if err != nil {
		log.Print(err)
		return
	}
}

func healthHandler(writer http.ResponseWriter, _ *http.Request) {
	conn, err := Connect(Host, Port, NoTLS)
	if err != nil {
		newErrorResponse(writer, 500, "Internal Server Error")
		return
	}
	defer CloseConnect(conn)

	hc := grpc_health_v1.NewHealthClient(conn)
	result, err := HealthCheck(&hc)
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

	httpResponse := HttpResponse{
		Status: statusCode,
		Msg:    msg,
		Result: nil,
	}

	res := new(bytes.Buffer)
	encoder := json.NewEncoder(res)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(httpResponse)

	if err != nil {
		log.Print(err)
		return
	}

	_, err = writer.Write(res.Bytes())
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
