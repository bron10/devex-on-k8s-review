# Appointments

This application is part of the Min Salus system and provides the functionality for managing appointments. It's part of the project built in the [Developer Experience on Kubernetes](#) book by [Mauricio Salatino](https://salaboy.com) and [Thomas Vitale](https://www.thomasvitale.com).

## HTTP API

| Endpoint	      | Method   | Req. body   | Status | Resp. body     | Description    		   	              |
|:---------------:|:--------:|:-----------:|:------:|:--------------:|:-------------------------------------|
| `/`             | `GET`    |             | 200    | String         | Welcome message.                     |
| `/appointments` | `GET`    |             | 200    | Appointment[]  | Get all the booked appointments.     |
| `/appointments` | `POST`   | Appointment | 201    | Appointment    | Book a new appointment.              |
| `/appointments` | `DELETE` |             | 204    |                | Delete all appointments.             |

Get the welcome message:

```shell script
http :8081
```

Book an appointment without a category:

```shell script
http :8081/appointments patientId=42 appointmentDate="2028-02-29T12:00:00Z"
```

Book an appointment with a category:

```shell script
http :8081/appointments patientId=42 category="cardiology" appointmentDate="2028-02-29T12:00:00Z"
```

Get all appointments:

```shell script
http :8081/appointments
```

Delete all appointments:

```shell script
http DELETE :8081/appointments
```

## Run

Run the application in development mode, with live reload:

```shell script
./gradlew bootRun
```

## Test

Run all unit and integration tests:

```shell script
./gradlew test
```

## Build

Build the application:

```shell script
./gradlew build
```

## Package (JAR)

Package the application as a JAR artifact:

```shell script
./gradlew bootJar
```

## Package (OCI)

Package the application as an OCI image:

```shell script
./gradlew bootBuildImage
```
