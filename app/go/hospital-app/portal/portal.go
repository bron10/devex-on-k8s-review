package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	api "github.com/salaboy/devex-on-k8s/app/go/hospital-app/portal/api"
)

var (
	Version = getEnv("VERSION", "1.0.0")

	Source                 = getEnv("SOURCE", "https://github.com/salaboy/devex-on-k8s/app/go/hospital-app/portal")
	PodName                = getEnv("POD_NAME", "N/A")
	PodNamespace           = getEnv("POD_NAMESPACE", "N/A")
	PodNodeName            = getEnv("POD_NODENAME", "N/A")
	PodIp                  = getEnv("POD_IP", "N/A")
	PodServiceAccount      = getEnv("POD_SERVICE_ACCOUNT", "N/A")
	AppointmentsServiceUrl = getEnv("APPOINTMENTS_SERVICE_URL", "http://localhost:8081")
	AppPort                = getEnv("APP_PORT", "8080")
	KoDataPath             = getEnv("KO_DATA_PATH", "kodata")
)

var events = []Event{}

const (
	ApplicationJson = "application/json"
	ContentType     = "Content-Type"
)

type ServiceInfo struct {
	Name              string `json:"name"`
	Version           string `json:"version"`
	Source            string `json:"source"`
	PodName           string `json:"podName"`
	PodNamespace      string `json:"podNamespace"`
	PodNodeName       string `json:"podNodeName"`
	PodIp             string `json:"podIp"`
	PodServiceAccount string `json:"podServiceAccount"`
}

type Event struct {
	Id      string `json:"id"`
	Payload string `json:"payload"`
	Type    string `json:"type"`
}

func main() {

	r := NewChiServer()

	log.Printf("Starting Patients Portal Go in Port: %s", AppPort)

	// Start the server; this is a blocking call
	err := http.ListenAndServe(":"+AppPort, r)
	if err != http.ErrServerClosed {
		log.Panic(err)
	}
}

func appointmentsServiceHandler(w http.ResponseWriter, r *http.Request) {
	proxyRequest("api/appointments", AppointmentsServiceUrl, w, r)
}

func proxyRequest(serviceName string, serviceUrl string, w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	r.Body = io.NopCloser(bytes.NewReader(body))

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(string(requestDump))

	url := fmt.Sprintf("%s%s", serviceUrl, r.RequestURI)
	if !strings.HasPrefix(url, "http") {
		url = fmt.Sprintf("http://%s", url)
	}

	log.Printf("Proxying request before replace to %s", url)
	// remove the service path
	url = strings.Replace(url, serviceName+"/", "", -1)

	log.Printf("Proxying request to %s", url)

	proxyReq, err := http.NewRequest(r.Method, url, bytes.NewReader(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	proxyReq.Header = make(http.Header)
	for h, val := range r.Header {
		proxyReq.Header[h] = val
	}

	resp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for h, val := range resp.Header {
		w.Header()[h] = val
	}

	w.WriteHeader(resp.StatusCode)

	log.Printf("Proxied request response code %s - %d", resp.Status, resp.StatusCode)

	_, err = w.Write(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// respondWithJSON is a helper function to write json response format.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// getEnv is a helper function to get environment variable or return a default value.
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

// OpenAPI OpenAPIHandler returns a handler that serves the OpenAPI documentation.
func OpenAPI(r *chi.Mux) {
	fs := http.FileServer(http.Dir(KoDataPath + "/docs/"))
	r.Handle("/openapi/*", http.StripPrefix("/openapi/", fs))
}

// server implements api.ServerInterface interface.
type server struct{}

// GetEventsWithPost gets all events from the in-memory store.
func (s *server) GetEventsWithPost(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, events)
}

// GetEventsWithGet gets all events from the in-memory store.
func (s *server) GetEventsWithGet(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, events)
}

// GetServiceInfo gets service information.
func (s *server) GetServiceInfo(w http.ResponseWriter, r *http.Request) {
	var info = ServiceInfo{
		Name:              "PORTAL",
		Version:           Version,
		Source:            Source,
		PodName:           PodName,
		PodNodeName:       PodNodeName,
		PodNamespace:      PodNamespace,
		PodIp:             PodIp,
		PodServiceAccount: PodServiceAccount,
	}
	w.Header().Set(ContentType, ApplicationJson)
	json.NewEncoder(w).Encode(info)
}

// NewServer creates a new api.ServerInterface server.
func NewServer() api.ServerInterface {
	return &server{}
}

// NewChiServer creates a new chi.Mux server.
func NewChiServer() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	fs := http.FileServer(http.Dir(KoDataPath))

	server := NewServer()

	OpenAPI(r)

	r.HandleFunc("/api/appointments/*", appointmentsServiceHandler)
	r.Mount("/api/", api.Handler(server))
	r.Handle("/*", http.StripPrefix("/", fs))

	// Add handlers for readiness and liveness endpoints
	r.HandleFunc("/health/{endpoint:readiness|liveness}", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	return r
}
