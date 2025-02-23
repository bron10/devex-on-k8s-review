package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	api "github.com/salaboy/devex-on-k8s/app/go/hospital-app/appointments/api"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

const (
	ApplicationJson = "application/json"
	ContentType     = "Content-Type"
)

// Event struct to encode events data
type Event struct {
	Id      string `json:"id"`
	Payload string `json:"payload"`
	Type    string `json:"type"`
}

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

type Appointment struct {
	Id           string    `json:"id"`
	PatientId    string    `json:"patientId"`
	DepartmentId string    `json:"departmentId"`
	Date         time.Time `json:"dateAndTime"`
}

func (s Appointment) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

var (
	VERSION             = getEnv("VERSION", "1.0.0")
	SOURCE              = getEnv("SOURCE", "https://github.com/")
	POD_NAME            = getEnv("POD_NAME", "N/A")
	POD_NAMESPACE       = getEnv("POD_NAMESPACE", "N/A")
	POD_NODENAME        = getEnv("POD_NODENAME", "N/A")
	POD_IP              = getEnv("POD_IP", "N/A")
	POD_SERVICE_ACCOUNT = getEnv("POD_SERVICE_ACCOUNT", "N/A")
	APP_PORT            = getEnv("APP_PORT", "8081")

	appointments = []Appointment{}
)

// respondWithJSON is a helper function to write a JSON response.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	chiServer := NewChiServer()

	// Start the server; this is a blocking call
	err := http.ListenAndServe(":"+APP_PORT, chiServer)
	log.Printf("Starting Appointments Service in Port: %s", APP_PORT)
	if err != http.ErrServerClosed {
		log.Panic(err)
	}
}

// getEnv returns the value of an environment variable, or a fallback value if not set.
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

// NewChiServer creates a new Chi server.
func NewChiServer() *chi.Mux {
	log.Printf("Starting Appointments Service in Port: %s", APP_PORT)

	// create new router
	r := chi.NewRouter()

	// add middlewares
	r.Use(middleware.Logger)

	// create new server
	server := NewServer()

	// add openapi spec
	OpenAPI(r)

	// add routes
	r.Mount("/", api.Handler(server))

	// add health check
	r.HandleFunc("/health/{endpoint:readiness|liveness}", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	return r
}

// server is the API server struct that implements api.ServerInterface.
type server struct {
}

// NewServer creates a new api.ServerInterface.
func NewServer() api.ServerInterface {
	return &server{}
}

// OpenApi returns a handler that serves the OpenAPI spec as JSON.
func OpenAPI(r *chi.Mux) {
	fs := http.FileServer(http.Dir(os.Getenv("KO_DATA_PATH") + "/docs/"))
	r.Handle("/openapi/*", http.StripPrefix("/openapi/", fs))
}

// GetAllAppointments returns all appointments.
func (s *server) GetAllAppointments(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, appointments)
}

// CreateAppointment creates a new appointment.
func (s *server) CreateAppointment(w http.ResponseWriter, r *http.Request) {

	var appointment Appointment
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		log.Printf("There was an error decoding the request body into the struct: %v", err)
	}

	appointment.Id = uuid.New().String()

	appointments = append(appointments, appointment)

	respondWithJSON(w, http.StatusOK, appointment)
}

// GetServiceInfo returns service information.
func (s *server) GetServiceInfo(w http.ResponseWriter, r *http.Request) {
	var info ServiceInfo = ServiceInfo{
		Name:              "APPOINTMENTS",
		Version:           VERSION,
		Source:            SOURCE,
		PodName:           POD_NAME,
		PodNodeName:       POD_NODENAME,
		PodNamespace:      POD_NAMESPACE,
		PodIp:             POD_IP,
		PodServiceAccount: POD_SERVICE_ACCOUNT,
	}
	w.Header().Set(ContentType, ApplicationJson)
	json.NewEncoder(w).Encode(info)
}
