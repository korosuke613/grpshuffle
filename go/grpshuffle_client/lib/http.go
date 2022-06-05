package grpshuffle_client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/korosuke613/grpshuffle/go/grpshuffle"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"io"
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
	FormatJson       ResponseFormat = "json"
	FormatJsonPretty ResponseFormat = "json-pretty"
	FormatReadable   ResponseFormat = "readable"
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
	http.HandleFunc("/repeat_shuffle", repeatShuffleHandler)
	http.HandleFunc("/", healthHandler)

	log.Printf("Launch grpshuffle-client http server %v...", port)
	log.Fatal(http.ListenAndServe(addr, logger(http.DefaultServeMux)))
}

func makeResponse(format ResponseFormat, shuffleResult []*grpshuffle.Combination, groupName string) (res []byte) {
	var err error

	switch format {
	case FormatJson:
		{
			res, err = json.Marshal(HttpResponse{
				Status: 200,
				Msg:    "Ok",
				Result: shuffleResult,
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
				Result: shuffleResult,
			}, "", "  ")
			if err != nil {
				log.Print(err)
				return
			}
		}
	case FormatReadable:
		{
			var prettyResponse string
			for i, combination := range shuffleResult {
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
				Result: shuffleResult,
			})
			if err != nil {
				log.Print(err)
				return
			}
		}
	}

	return res
}

type shuffleOptions struct {
	ResponseFormat ResponseFormat
	Divide         int
	GroupName      string
	Targets        []string
	RepeatInterval int
	RepeatTimes    int
}

func getOptionsFromFormValue(request *http.Request) (*shuffleOptions, error) {
	format := ResponseFormat(request.FormValue("fmt"))
	if format != "" {
		err := format.checkType()
		if err != nil {
			return nil, err
		}
	}

	rawDivide := request.FormValue("divide")
	if rawDivide == "" {
		return nil, fmt.Errorf("divide parameter is required")
	}
	divide, err := strconv.Atoi(rawDivide)
	if err != nil {
		return nil, fmt.Errorf("divide parameter allows numbers")
	}

	var repeatInterval int
	rawRepeatInterval := request.FormValue("interval")
	if rawRepeatInterval != "" {
		repeatInterval, err = strconv.Atoi(rawRepeatInterval)
		if err != nil {
			return nil, fmt.Errorf("interval parameter allows numbers")
		}
	}

	var repeatTimes int
	rawRepeatTimes := request.FormValue("times")
	if rawRepeatTimes != "" {
		var err error
		repeatTimes, err = strconv.Atoi(rawRepeatTimes)
		if err != nil {
			return nil, fmt.Errorf("times parameter allows numbers")
		}
	}

	groupName := request.FormValue("groupName")
	if groupName == "" {
		groupName = "Group"
	}

	rawTargets := request.FormValue("targets")
	if rawTargets == "" {
		return nil, fmt.Errorf("targets parameter is required")
	}
	targets := strings.Split(rawTargets, ",")

	return &shuffleOptions{
		format,
		divide,
		groupName,
		targets,
		repeatInterval,
		repeatTimes,
	}, nil
}

// WIP
func repeatShuffleHandler(writer http.ResponseWriter, request *http.Request) {
	conn, err := Connect(Host, Port, NoTLS)
	if err != nil {
		newErrorResponse(writer, 500, "Internal Server Error")
		return
	}
	defer CloseConnect(conn)

	options, err := getOptionsFromFormValue(request)
	if err != nil {
		newErrorResponse(writer, 400, err.Error())
		log.Print(err)
		return
	}

	cc := grpshuffle.NewComputeClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repeatShuffleStream, err := RepeatShuffle(ctx, &cc, uint64(options.RepeatTimes), uint64(options.RepeatInterval), uint64(options.Divide), options.Targets)
	if err != nil {
		newErrorResponse(writer, 504, "Gateway Timeout")
		log.Print(err)
		return
	}

	count := uint64(0)
	for {
		if options.RepeatTimes != 0 && uint64(options.RepeatTimes) < count {
			return
		}

		resp, err := repeatShuffleStream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}

			if status.Code(err) == codes.Canceled {
				break
			}
			log.Print(fmt.Errorf("receiving response: %w", err))
			return
		}

		res := makeResponse(options.ResponseFormat, resp.Combinations, options.GroupName)
		log.Print(res)

		_, err = writer.Write(res)
		if err != nil {
			log.Print(err)
			return
		}

		count += 1
	}
}

func shuffleHandler(writer http.ResponseWriter, request *http.Request) {
	conn, err := Connect(Host, Port, NoTLS)
	if err != nil {
		newErrorResponse(writer, 500, "Internal Server Error")
		return
	}
	defer CloseConnect(conn)

	options, err := getOptionsFromFormValue(request)
	if err != nil {
		newErrorResponse(writer, 400, err.Error())
		log.Print(err)
		return
	}

	cc := grpshuffle.NewComputeClient(conn)
	result, err := Shuffle(&cc, uint64(options.Divide), options.Targets)
	if err != nil {
		newErrorResponse(writer, 504, "Gateway Timeout")
		log.Print(err)
		return
	}

	var res = makeResponse(options.ResponseFormat, result, options.GroupName)

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
