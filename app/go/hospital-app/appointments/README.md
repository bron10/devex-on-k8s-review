# Appointments Service

This service allow users to schedule new appointments and list them all. 




## Build and run from source


```
go build -o appointments-service
```

To change the API and generate the code: 

```
oapi-codegen -generate chi-server -package api kodata/docs/openapi.yaml > api/api.go
```

## Running & calling

```
./appointments-service
```

Using `httpie`:

```
http :8081/appointments/ < new-appointment.json
```
