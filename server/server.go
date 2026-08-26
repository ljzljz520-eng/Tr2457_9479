package server

import (
	"context"
	"encoding/json"
	"frontend_go/catalog"
	"frontend_go/events"
	"frontend_go/model"
	"net/http"
	"strings"
)

type API struct {
	catalog    *catalog.Service
	dispatcher *events.Dispatcher
}

func New(c *catalog.Service, d *events.Dispatcher) *API { return &API{catalog: c, dispatcher: d} }
func (a *API) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/records", a.records)
	mux.HandleFunc("/events", a.publish)
	return mux
}
func (a *API) records(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		q := r.URL.Query().Get("q")
		items, e := a.catalog.Search(q)
		if e != nil {
			http.Error(w, e.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(items)
		return
	}
	if r.Method == http.MethodPost {
		var rec model.Record
		if e := json.NewDecoder(r.Body).Decode(&rec); e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		if e := a.catalog.Register(rec); e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		w.WriteHeader(201)
		return
	}
	http.Error(w, "method not allowed", 405)
}
func (a *API) publish(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
		return
	}
	var e model.Event
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := a.dispatcher.Publish(r.Context(), e); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(202)
}
func Run(ctx context.Context, addr string, h http.Handler) error {
	srv := &http.Server{Addr: addr, Handler: h}
	go func() { <-ctx.Done(); srv.Shutdown(context.Background()) }()
	return srv.ListenAndServe()
}
func ParseID(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}
