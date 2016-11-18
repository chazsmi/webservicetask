# Web Service task

- Currently the service is very simple and just says hello to the name you give it

## To run the service 
```bash
go run main.go
```
or
```bash
go install && $GOBIN/webservicetask
```

## Query the service
```bash
curl localhost:8000?name=charlie
```

## Lesson 1 tasks
- Change the output of the service from plain text to json.
- Allow a port to be passed in as a flag for the service to run on
- Create some simple validation for the name parameter with the following rules
    1. Check the name is present. Return an error message informing the client if not 
    2. The name must greater the one character long
- Create a seperate function which handles the sending of the response and supports multiple status codes to be passed in
