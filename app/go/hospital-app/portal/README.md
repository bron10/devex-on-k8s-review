# Portal Service

This service provides the backend and frontend for the Patient's portal. 




## Build and run from source


```
go build -o portal
```

To change the API and generate the code: 

```
oapi-codegen -generate chi-server -package api kodata/docs/openapi.yaml > api/api.go
```