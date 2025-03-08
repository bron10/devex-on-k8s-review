# Appointments Service

This service allow users to schedule new appointments and list them all. 


## Build and run from source

In a new terminal run `docker compose up` to start PostgreSQL:
```
docker compose up
```

Now you can start the appointments service by running: 

```
go run appointments.go
```

## Running & calling


Using `httpie`:

```
http :8081/appointments/ < new-appointment.json
```
