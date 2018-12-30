# Go - Mongo - Gin REST API CRUD
A Simple Go mongo gin crud

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites
1. golang
2. glide / dep / vgo

### Installing

Clone the Repo

```bash
go get github.com/alamin-mahamud/go-mongo-gin
cd $GOPATH/src/github.com/alamin-mahamud/go-mongo-gin
```

Install the requirements

```bash
glide up
```

## Running the tests

```bash
go test ./...
```

## Running the Application
```bash
go run server.go
```

## Testing all the endpoints
```bash
curl -v -H "Accept: application/json" -H "Content-type: application/json" -X GET http://localhost:8000/users 
curl -v -H "Accept: application/json" -H "Content-type: application/json" -X GET http://localhost:8000/users/{oid}
curl -v -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"name":"Amila","gender":"Female", "Age":34}' http://localhost:8000/users
curl -v -H "Accept: application/json" -H "Content-type: application/json" -X PUT -d '{"name":"Amila","gender":"male", "Age":34}' http://localhost:8000/users/{oid}
curl -i -H "Accept: application/json" -H "Content-type: application/json" -X DELETE http://localhost:8000/users/{oid}
```

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors
* **H.G Nuwan Indika** - *Initial work* 
* **Alamin Mahamud** - [alamin.rocks](https://alamin-rocks.herokuapp.com)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* Hat tip to anyone whose code was used