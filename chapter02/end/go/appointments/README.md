# Appointments Service

This service allow users to schedule new appointments and list them all. 


## Run from source

In a new terminal run `docker compose up` to start PostgreSQL:
```
docker compose up
```

Now you can start the appointments service by running: 

```
go run appointments.go
```

## Interacting with the appointments service

List all appointments: 

```
http :8081/appointments/
```

Create a new appointment using `httpie`:

```
http :8081/appointments/ < new-appointment.json
```

Delete all appointments: 
```
http delete :8081/appointments/
```

When finished stop docker compose with `docker compose down`

## Testing

First start the the docker-compose for tests:
```
docker compose -f tests/docker-compose up
```

Then run all Go Tests:

```
go test
```
