package core_store

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type HttpEndpointsFactory interface {
	ListMoviesEndpoint() func(w http.ResponseWriter, r *http.Request)
	GetMovieByIdEndpoint(idParam string) func(w http.ResponseWriter, r *http.Request)
	Register() func(w http.ResponseWriter, r *http.Request)
	Login() func(w http.ResponseWriter, r *http.Request)
}

type httpEndpointsFactory struct {
	coreService CoreService
}

type customError struct {
	Message string `json:"message"`
}

func NewHttpEndpoints(userService CoreService) HttpEndpointsFactory {
	return &httpEndpointsFactory{coreService: userService}
}

func (httpFac *httpEndpointsFactory) ListMoviesEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		listMovieReq := &ListMoviesCommand{}
		if r.Header.Get("Content-Type") == "application/json" {
			err := json.NewDecoder(r.Body).Decode(listMovieReq)
			if err != nil {
				respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
				return
			}
		}
		count, err := strconv.ParseInt(r.URL.Query().Get("count"), 10, 64)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
			return
		}
		listMovieReq.Count = count
		data, err := listMovieReq.Exec(httpFac.coreService)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
			return
		}
		respondJSON(w, http.StatusOK, data)
	}
}

func (httpFac *httpEndpointsFactory) GetMovieByIdEndpoint(idParam string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		movieReq := &GetMovieByIdCommand{}
		vars := mux.Vars(r)
		idStr, ok := vars[idParam]
		if !ok {
			respondJSON(w, http.StatusInternalServerError, &customError{"no token param"})
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, &customError{"no int"})
			return
		}
		movieReq.Id = id
		if r.Header.Get("Content-Type") == "application/json" {
			err := json.NewDecoder(r.Body).Decode(movieReq)
			if err != nil {
				respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
				return
			}
		}
		data, err := movieReq.Exec(httpFac.coreService)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
			return
		}
		respondJSON(w, http.StatusOK, data)
	}
}

func (httpFac *httpEndpointsFactory) Register() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		userCmd := &CreateUserCommand{}
		if r.Header.Get("Content-Type") == "application/json" {
			err := json.NewDecoder(r.Body).Decode(userCmd)
			if err != nil {
				respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
				return
			}
		}
		data, err := userCmd.Exec(httpFac.coreService)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
			return
		}
		respondJSON(w, http.StatusCreated, data)
	}
}

func (httpFac httpEndpointsFactory) Login() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token,Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "3600")
		userCmd := &LoginUserCommand{}
		if r.Header.Get("Content-Type") == "application/json" {
			err := json.NewDecoder(r.Body).Decode(userCmd)
			if err != nil {
				respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
				return
			}
		}
		data, err := userCmd.Exec(httpFac.coreService)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
			return
		}
		respondJSON(w, http.StatusOK, data)
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
