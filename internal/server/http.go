package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHTTPServer(addr string) *http.Server {
	httpSrv := newhttpServer()

	r := mux.NewRouter()
	r.HandleFunc("/", httpSrv.handleProduce).Methods(http.MethodPost)
	r.HandleFunc("/", httpSrv.handleConsume).Methods(http.MethodGet)

	return &http.Server{Addr: addr, Handler: r}
}

type httpServer struct {
	log *Log
}

func newhttpServer() *httpServer {
	return &httpServer{log: NewLog()}
}

type ProduceRequest struct {
	Record Record `json:"record"`
}

type ConsumeResponse struct {
	Record Record `json:"record"`
}

type ProduceResponse struct {
	Offset uint64 `json:"offset"`
}

type ConsumeRequest struct {
	Offset uint64 `json:"offset"`
}

func (httpSrv *httpServer) handleProduce(w http.ResponseWriter, r *http.Request) {
	var req ProduceRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	offset, err := httpSrv.log.Append(req.Record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := ProduceResponse{Offset: offset}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (httpSrv *httpServer) handleConsume(w http.ResponseWriter, r *http.Request) {
	var req ConsumeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	record, err := httpSrv.log.Read(req.Offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := ConsumeResponse{Record: record}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
