# Web Service task


## To run the service 
```bash
go run main.go
```
or
```bash
go install && $GOBIN/webservicetask
```
## Lesson 6 
- Add profiling and benchmarking tools reafactor your application to make it more efficent and reducing memory allocations.

## Lesson 5
- Refactor your code into seperate packages if you havent already done so
- See if you can refactor your code to use you're favourite design pattern

## Lesson 4 
- Change the service to store its data in Mysql 
- Think about about how you can abstract away the database detail
- Change the rest of you're service to use the database code
- Run you're service with a static binary

## Lesson 3
- Change you're service to use context for requests
- When a developer is requested, search for that developers name on Google and return the first result as part of your result. Make your request to Google concurrent.

## Lesson 2 
We now need to improve our service a little. The service now needs to hold details of developers and expose two endpoints
- /developers - GET returns all developers POST adds a new developer
- /developers/1 - In the normal rest fashion you can retrieve specfic developer by their id.
We arent going to use any data store at the moment we are just going to hold our data store in memory. This has been created for you in the developers package

## Query the service
```bash
curl localhost:8000/developers
curl localhost:8000/developers/1
```

### Tasks
- Add the new endpoints to your service
- Split your program into packages where possible
- Add unit tests for all functions in all packages. Making sure to use the new 1.7 features. [https://golang.org/pkg/net/http/httptest/](https://golang.org/pkg/net/http/httptest/)
- Benchmark at least one function in each package

## Lesson 1
- Currently the service is very simple and just says hello to the name you give it

## Query the service
```bash
curl localhost:8000?name=charlie
```

### Tasks
- Change the output of the service from plain text to json.
- Allow a port to be passed in as a flag for the service to run on
- Create some simple validation for the name parameter with the following rules
    1. Check the name is present. Return an error message informing the client if not 
    2. The name must greater the one character long
- Create a seperate function which handles the sending of the response and supports multiple status codes to be passed in


