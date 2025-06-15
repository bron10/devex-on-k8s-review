# Appointments

This application is part of the Min Salus system and provides the functionality for managing appointments. It's part of the project built in the [Developer Experience on Kubernetes](https://www.manning.com/books/developer-experience-on-kubernetes) book by [Mauricio Salatino](https://salaboy.com) and [Thomas Vitale](https://www.thomasvitale.com).

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

Book an appointment:

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

## Run (Skaffold)

Run the application in development mode on Kubernetes, with live reload:

```shell script
skaffold dev --port-forward
```

The application will start on port `8081` by default and the process will keep running, watching for changes in the source code. When you're done, stop the application process with `Ctrl+C`.

## Run (JKube)

First, provision a PostgreSQL database in your local Kubernetes cluster:

```shell script
kubectl apply -f config/db.yml
```

Build the application as a container image:

```shell script
./gradlew bootBuildImage
```

Then, load the image to the local cluster:

```shell
kind load docker-image appointments:0.0.1-SNAPSHOT --name devex-cluster
```

Next, generate the Kubernetes manifests and deploy the application with JKube.

```shell
./gradlew k8sResource k8sApply
```

When you're done, you can undeploy the application as follows:

```shell
./gradlew k8sUndeploy
```

And finally, unprovision the PostgreSQL database:

```shell script
kubectl delete -f config/db.yml
```


## Run (Manifests)

First, provision a PostgreSQL database in your local Kubernetes cluster:

```shell script
kubectl apply -f config/db.yml
```

Build the application as a container image:

```shell script
./gradlew bootBuildImage
```

Then, load the image to the local cluster. Due to some issues with the `kind` CLI, you need to run the command twice to ensure the image is loaded correctly.

```shell
kind load docker-image appointments:0.0.1-SNAPSHOT --name devex-cluster
```

Next, deploy the application using the Kubernetes manifests provided in the `config` directory:

```shell
kubectl apply -f config
```

When you're done, you can undeploy the application as follows:

```shell
kubectl delete -f config
```

## Build (Spring Boot)

Build the application as a container image using the Buildpacks integration in Spring Boot:

```shell script
./gradlew bootBuildImage
```

## Build (Pack)

Build the application as a container image using the Pack CLI:

```shell script
pack build appointments:0.0.1-SNAPSHOT
```

## Clean

Clean the build directory:

```shell script
./gradlew clean
```
