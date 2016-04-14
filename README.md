# gin-rest-mongodb-sample

	## you need to import following packages

  		1.go get github.com/gin-gonic/gin

  		2.go get gopkg.in/mgo.v2

  		3.go get gopkg.in/mgo.v2/bson

	## To run the server

  		1.go run server.go


  ## Routs

  		1.curl -v -H "Accept: application/json" -H "Content-type: application/json" -X GET http://localhost:8000/users 

  		2.curl -v -H "Accept: application/json" -H "Content-type: application/json" -X GET http://localhost:8000/users/{oid}

  		3.curl -v -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"name":"Amila","gender":"Female", "Age":34}' http://localhost:8000/users
  		
		  4.curl -v -H "Accept: application/json" -H "Content-type: application/json" -X PUT -d '{"name":"Amila","gender":"male", "Age":34}' http://localhost:8000/users/{oid}

  		5.curl -i -H "Accept: application/json" -H "Content-type: application/json" -X DELETE http://localhost:8000/users/{oid}

