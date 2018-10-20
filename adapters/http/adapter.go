package http

import (
	"net/http"

	"context"
	"log"

	"encoding/json"
	"io/ioutil"

	"fmt"
	"github.com/ablce9/go-assignment/domain"
	"github.com/ablce9/go-assignment/engine"
	"github.com/gorilla/mux"
	"os"
)

// Adapter ...
type Adapter struct {
	Server *http.Server
}

// Start starts http server.
func (adapter *Adapter) Start() {
	// todo: start to listen
	go func() {
		// Make server proc background; thus we can shutdown
		// sevrer gracefully.
		log.Fatal(adapter.Server.ListenAndServe())
	}()

	http.DefaultServeMux.Handle("/", adapter.Server.Handler)
	log.Println("Start(): Server is listening on", adapter.Server.Addr)
}

// Stop shutdowns http server gracefully.
func (adapter *Adapter) Stop() {
	// todo: shutdown server
	if err := adapter.Server.Shutdown(context.Background()); err != nil {
		log.Printf("Shutdown(): error: %v", err)
	}
}

type handler struct {
	Engine engine.Engine
}

// NewAdapter ...
func NewAdapter(e engine.Engine) *Adapter {
	// todo: init your http server and routes
	h := handler{
		Engine: e,
	}
	router := mux.NewRouter()
	router.HandleFunc("/knight", h.knightHandler).Methods("POST", "GET")
	router.HandleFunc("/knight/{id}", h.getKnightsHandler).Methods("GET")
	router.Use(LoggingMiddleware)

	address := os.Getenv("GO_ASSIGNMENT_ADDR")
	if address == "" {
		address = ":5000"
	}

	srv := &http.Server{
		// TODO: Setup timeout for read/write
		Handler: router,
		Addr:    address,
	}

	return &Adapter{
		Server: srv,
	}
}

type provider interface {
	GetKnightRepository() engine.KnightRepository
}

func (h handler) knightHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		knight := domain.Knight{
			Name:        "",
			Strength:    -1,
			WeaponPower: -1,
		}
		if err := json.Unmarshal(buf, &knight); err != nil ||
			knight.Name == "" ||
			knight.Strength < 0 || knight.WeaponPower < 0 {
			reason := struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			}{
				http.StatusBadRequest,
				"bad data",
			}
			w.WriteHeader(http.StatusBadRequest)
			marshaled, _ := json.Marshal(reason)
			w.Write(marshaled)
			return
		}

		if p, ok := h.Engine.(provider); ok {
			repo := p.GetKnightRepository()
			repo.Save(&knight)
			w.WriteHeader(http.StatusCreated)
			return
		}
	} else if r.Method == "GET" {
		if p, ok := h.Engine.(provider); ok {
			repo := p.GetKnightRepository()
			knights := repo.FindAll()
			if len(knights) == 0 {
				reason := struct {
					Code    int    `json:"code"`
					Message string `json:"message"`
				}{
					http.StatusNotFound,
					"no data found",
				}
				w.WriteHeader(http.StatusNotFound)
				marshaled, _ := json.Marshal(reason)
				w.Write(marshaled)
				return
			}
			marshaled, _ := json.Marshal(knights)
			w.Header().Set("Content-Type", "application/json")
			w.Write(marshaled)
			return
		}
	}
	http.Error(w, "404 page not found", http.StatusNotFound)
}

func (h handler) getKnightsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id != "" {
		if p, ok := h.Engine.(provider); ok {
			repo := p.GetKnightRepository()
			knight := repo.Find(id)
			if knight == nil {
				message := fmt.Sprintf("Knight #%s not found.", id)
				reason := struct {
					Code    int    `json:"code"`
					Message string `json:"message"`
				}{
					http.StatusNotFound,
					message,
				}
				w.WriteHeader(http.StatusNotFound)
				marshaled, _ := json.Marshal(reason)
				w.Write(marshaled)
				return
			}
			marshaled, _ := json.Marshal(knight)
			w.Header().Set("Content-Type", "application/json")
			w.Write(marshaled)
			return
		}
	}
	http.Error(w, "404 page not found", http.StatusNotFound)
}
