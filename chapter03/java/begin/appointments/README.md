# Appointments

This application is part of the Min Salus system and provides the functionality for managing appointments. It's part of the project built in the [Developer Experience on Kubernetes](https://www.manning.com/books/developer-experience-on-kubernetes) book by [Mauricio Salatino](https://salaboy.com) and [Thomas Vitale](https://www.thomasvitale.com).

## HTTP API

| Endpoint	      | Method   | Req. body   | Status | Resp. body     | Description    		   	              |
|:---------------:|:--------:|:-----------:|:------:|:--------------:|:-------------------------------------|
| `/`             | `GET`    |             | 200    | String         | Welcome message.                     |
| `/appointments` | `GET`    |             | 200    | Appointment[]  | Get all the booked appointments.     |
| `/appointments` | `POST`   | Appointment | 201    | Appointment    | Book a new appointment.              |
| `/appointments` | `DELETE` |             | 204    |                | Delete all appointments.             |

## Build (Pack)

Build the application as a container image using the Pack CLI:

```shell script
pack build appointments:0.0.1-SNAPSHOT
```

## Build (Spring Boot)

Build the application as a container image using the Buildpacks integration in Spring Boot:

```shell script
./gradlew bootBuildImage
```

## Run (Compose)

Run the application and its dependencies using Compose:

```shell script
podman compose up -d
```

Book an appointment:

```shell script
http :8081/appointments patientId=42 category="cardiology" appointmentDate="2028-02-29T12:00:00Z"
```

Stop the application and its dependencies:

```shell script
podman compose down
```

## Run (Kubernetes with Buildpacks)

First, provision a PostgreSQL database in your local Kubernetes cluster:

```shell script
kubectl apply -f config/db.yml
```

Load the container image built previously with Buildpacks:

```shell script
kind load docker-image appointments:0.0.1-SNAPSHOT --name devex-cluster
```

Deploy the application to Kubernetes:

```shell script
kubectl apply -f config
```

Validate the resources created:

```shell script
kubectl get all -l app=appointments
```

Book an appointment:

```shell script
http :9090/appointments patientId=42 category="cardiology" appointmentDate="2028-02-29T12:00:00Z"
```

Undeploy the application and its dependencies:

```shell script
kubectl delete -f config
```

## Run (Kubernetes with JKube)

First, provision a PostgreSQL database in your local Kubernetes cluster:

```shell script
kubectl apply -f config/db.yml
```

Load the container image built previously with Buildpacks:

```shell script
kind load docker-image appointments:0.0.1-SNAPSHOT --name devex-cluster
```

Generate the Kubernetes manifests and deploy the application with JKube.

```shell
./gradlew k8sResource k8sApply
```

Validate the resources created:

```shell script
kubectl get all -l app=appointments
```

Book an appointment:

```shell script
http :9090/appointments patientId=42 category="cardiology" appointmentDate="2028-02-29T12:00:00Z"
```

When you're done, you can undeploy the application as follows:

```shell
./gradlew k8sUndeploy
```

And finally, unprovision the PostgreSQL database:

```shell script
kubectl delete -f config/db.yml
```

## Run (Skaffold)

Run the application in development mode on Kubernetes, with live reload:

```shell script
skaffold dev --port-forward
```

The application will start on port `8081` by default and the process will keep running, watching for changes in the source code.

Book an appointment:

```shell script
http :8081/appointments patientId=42 category="cardiology" appointmentDate="2028-02-29T12:00:00Z"
```

When you're done, stop the application process with `Ctrl+C`.

## Clean

Clean the build directory:

```shell script
./gradlew clean
```
