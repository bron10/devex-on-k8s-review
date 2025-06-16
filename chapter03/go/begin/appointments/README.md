# Appointments Service

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

## Build (ko)

Build the application as a container image using the ko CLI:

```shell script
ko build --local -B -t 0.0.1-SNAPSHOT
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

## Deploy on Kubernetes

Create a Kind cluster with: 

```shell
kind create cluster
```

You should see something like this: 

```shell
Creating cluster "kind" ...
 ✓ Ensuring node image (kindest/node:v1.27.3) 🖼
 ✓ Preparing nodes 📦  
 ✓ Writing configuration 📜 
 ✓ Starting control-plane 🕹️ 
 ✓ Installing CNI 🔌 
 ✓ Installing StorageClass 💾 
Set kubectl context to "kind-kind"
You can now use your cluster with:

kubectl cluster-info --context kind-kind

Not sure what to do next? 😅  Check out https://kind.sigs.k8s.io/docs/user/quick-start/
```

Then we can build, package and deploy the appointments service to our cluster by running: 


```shell
ko apply -f kubernetes/
```

This also should fail on Kubernetes, you can check the logs: 

```
k logs -f appointments-788bcd88fd-5nr78
2025/03/29 10:27:02 Starting Appointments Service in Port: 8081
2025/03/29 10:27:02 Connecting to Database: postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable.
2025/03/29 10:27:02 dial tcp [::1]:5432: connect: connection refused
```

Installing PostgreSQL using Helm:

```shell
helm install postgresql oci://registry-1.docker.io/bitnamicharts/postgresql
```