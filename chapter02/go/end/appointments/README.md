# Appointments Service

This service allow users to schedule new appointments and list them all. 


## Run from source

In a new terminal, start PostgreSQL:

```shell
docker compose up -d
```

Now you can start the appointments service by running: 

```shell
go run appointments.go
```

## Interacting with the appointments service

List all appointments: 

```shell
http :8081/appointments/
```

Create a new appointment using `httpie`:

```shell
http :8081/appointments/ < new-appointment.json
```

Delete all appointments:

```shell
http delete :8081/appointments/
```

When finished stop docker compose with `docker compose down`

## Testing

First start the test database:

```shell
docker compose -f tests/compose.yml up
```

Then run all Go Tests:

```shell
go test
```

When you're done, stop the test database:

```shell
docker compose -f tests/compose.yml down
```
