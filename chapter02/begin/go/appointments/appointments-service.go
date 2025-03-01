package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	api "github.com/salaboy/devex-on-k8s/app/go/hospital-app/appointments/api"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var (
	VERSION            = getEnv("VERSION", "1.0.0")
	SOURCE             = getEnv("SOURCE", "https://github.com/")
	APP_PORT           = getEnv("APP_PORT", "8081")
	PostgresqlHost     = getEnv("POSTGRES_HOST", "localhost")
	PostgresqlPort     = getEnv("POSTGRES_PORT", "5432")
	PostgresqlUsername = getEnv("POSTGRES_USERNAME", "postgres")
	PostgresqlPassword = getEnv("POSTGRES_PASSWORD", "postgres")
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

	// connect to database
	db := NewDB()

	// check if database is alive
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to PostgreSQL.")

	// create new server
	server := NewServer(db)
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
	DB *sql.DB
}

// NewServer creates a new api.ServerInterface.
func NewServer(db *sql.DB) api.ServerInterface {
	return &server{
		DB: db,
	}
}

// OpenApi returns a handler that serves the OpenAPI spec as JSON.
func OpenAPI(r *chi.Mux) {
	fs := http.FileServer(http.Dir(os.Getenv("KO_DATA_PATH") + "/docs/"))
	r.Handle("/openapi/*", http.StripPrefix("/openapi/", fs))
}

// GetAllAppointments returns all appointments.
func (s *server) GetAllAppointments(w http.ResponseWriter, r *http.Request) {
	var query = "SELECT id, patientId, departmentId, appointmentDate FROM Appointments a"
	var rows *sql.Rows
	var err error

	rows, err = s.DB.Query(query)

	if err != nil {
		log.Printf("There was an error executing the query %v", err)
	}

	defer rows.Close()
	var appointments []Appointment
	for rows.Next() {

		var appointment Appointment
		err = rows.Scan(&appointment.Id, &appointment.PatientId, &appointment.DepartmentId, &appointment.AppointmentDate)
		if err != nil {
			log.Printf("There was an error scanning the sql rows: %v", err)
		}
		appointments = append(appointments, appointment)

	}

	log.Printf("Appointments retrieved from Database: %d", len(appointments))
	respondWithJSON(w, http.StatusOK, appointments)
}

// CreateAppointment creates a new appointment.
func (s *server) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appointment Appointment
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		log.Printf("There was an error decoding the request body into the struct: %v", err)
		respondWithJSON(w, http.StatusInternalServerError, err)
		return
	}

	appointment.Id = uuid.New().String()

	insertStmt := `insert into Appointments(id, patientId, departmentId, appointmentDate) values($1, $2, $3, $4)`

	_, err = s.DB.Exec(insertStmt, appointment.Id, appointment.PatientId, appointment.DepartmentId, appointment.AppointmentDate)

	if err != nil {
		log.Printf("An error occurred while executing query: %v", err)
		respondWithJSON(w, http.StatusInternalServerError, err)
		return
	}

	log.Printf("Appointment Stored in Database: %v", appointment)

	respondWithJSON(w, http.StatusOK, appointment)

}

// GetServiceInfo returns service information.
func (s *server) GetServiceInfo(w http.ResponseWriter, r *http.Request) {
	var info ServiceInfo = ServiceInfo{
		Name:    "APPOINTMENTS",
		Version: VERSION,
		Source:  SOURCE,
	}
	w.Header().Set(ContentType, ApplicationJson)
	json.NewEncoder(w).Encode(info)
}

// Welcome returns a welcome message from the Appointments Service
func (s *server) Welcome(w http.ResponseWriter, r *http.Request) {
	var welcome Welcome = Welcome{
		Message: "Welcome: you are now talking to the Appointments Service",
	}
	w.Header().Set(ContentType, ApplicationJson)
	json.NewEncoder(w).Encode(welcome)
}

func NewDB() *sql.DB {
	connStr := "postgresql://" + PostgresqlUsername + ":" + PostgresqlPassword + "@" + PostgresqlHost + ":" + PostgresqlPort + "/postgres?sslmode=disable"
	log.Printf("Connecting to Database: %s.", connStr)
	// Connect to database

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

const (
	ApplicationJson = "application/json"
	ContentType     = "Content-Type"
)

type ServiceInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Source  string `json:"source"`
}

type Appointment struct {
	Id              string    `json:"id"`
	PatientId       string    `json:"patientId"`
	DepartmentId    string    `json:"departmentId"`
	AppointmentDate time.Time `json:"appointmentDate"`
}

type Welcome struct {
	Message string `json:"message"`
}

func (s Appointment) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}
