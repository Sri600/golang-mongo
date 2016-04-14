package controllers

// created by H.G Nuwan Indika 

import (

	"log"
	
	"github.com/gin-gonic/gin"
	"../models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    //"strconv"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct {
		session *mgo.Session
	}
)

const (
    DB_NAME       = "sri600"
    DB_COLLECTION = "users"
)

// NewUserController provides a reference to a UserController with provided mongo session
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}



func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}

func messageTypeDefault(msg string,c *gin.Context) {
        content := gin.H{
            "status": "200",
            "result": msg,
        }
        c.Writer.Header().Set("Content-Type", "application/json")
        c.JSON(201, content)
}

func checkErrTypeOne(err error, msg string, status string , c *gin.Context) {
    if err != nil {
    	panic(err)
    	 log.Fatalln(msg, err)
        content := gin.H{
			"status": status,
            "result": msg,
        }
        c.Writer.Header().Set("Content-Type", "application/json")
        c.JSON(200, content)

    }
}

func checkErrTypeTwo(msg string, status string , c *gin.Context) {
        content := gin.H{
			"status": status,
            "result": msg,
        }
        c.Writer.Header().Set("Content-Type", "application/json")
        c.JSON(200, content)
}


// Get all Users
func (uc UserController) UsersList(c *gin.Context) {
	
	var results []models.User
    err := uc.session.DB(DB_NAME).C(DB_COLLECTION).Find(nil).All(&results)
    if err != nil{
		checkErrTypeOne(err, "Users doesn't exist","404",c)
		return
	}
  
    c.JSON(200, results)
}

// GetUser retrieves an individual user resource
func (uc UserController) GetUser(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
	
	 //Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		checkErrTypeTwo("ID is not a bson.ObjectId","404",c)
		return
	}
	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub user
	u := models.User{}
	err := uc.session.DB(DB_NAME).C(DB_COLLECTION).FindId(oid).One(&u)
	// Fetch user
	if err != nil {
		checkErrTypeTwo("Users doesn't exist","404",c)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
    c.JSON(200, u)
}

// CreateUser creates a new user resource
func (uc UserController) CreateUser(c *gin.Context) {

    var json models.User

    // This will infer what binder to use depending on the content-type header.
    c.Bind(&json) 

	u := uc.create_user(json.Name, json.Gender,json.Age,c)
    if u.Name == json.Name {
        content := gin.H{
            "result": "Success",
            "Name": u.Name,
            "Gender": u.Gender,
            "Age": u.Age,
        }
    
        c.Writer.Header().Set("Content-Type", "application/json")
        c.JSON(201, content)
    } else {
        c.JSON(500, gin.H{"result": "An error occured"})
    }
	
}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id){
		checkErrTypeTwo("ID is not a bson.ObjectId","404",c)
		return
	}
	// Grab id
	oid := bson.ObjectIdHex(id)
	
	// Remove user
	if err := uc.session.DB(DB_NAME).C(DB_COLLECTION).RemoveId(oid); err != nil{
		checkErrTypeOne(err,"Fail to Remove","404",c)
		return
	}

	messageTypeDefault("Success",c)
		
}


// RemoveUser removes an existing user resource
func (uc UserController) UpdateUser(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
    var json models.User

    // This will infer what binder to use depending on the content-type header.
    c.Bind(&json) 

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		checkErrTypeTwo("ID is not a bson.ObjectId","404",c)
		return
	}

	// Grab id
	
	u := uc.update_user(id,json.Name, json.Gender,json.Age,c)

    if u.Name == json.Name {
        content := gin.H{
            "result": "Success",
            "Name": u.Name,
            "Gender": u.Gender,
            "Age": u.Age,
        }
    
        c.Writer.Header().Set("Content-Type", "application/json")
        c.JSON(201, content)
    } else {
        c.JSON(500, gin.H{"result": "An error occured"})
    }

	// Write status
	//c.AbortWithStatus(200)
}

func (uc UserController) create_user(Name string, Gender string,Age int, c *gin.Context) models.User {
    user := models.User{
        Name:      Name,
        Gender:    Gender,
        Age:	Age,
    }
    // Write the user to mongo
	err := uc.session.DB(DB_NAME).C(DB_COLLECTION).Insert(&user)
    checkErrTypeOne(err, "Insert failed","403",c)
    return user
}

func (uc UserController) update_user(Id string,Name string, Gender string,Age int, c *gin.Context) models.User {

    user := models.User{

        Name:      Name,
        Gender:    Gender,
        Age:	Age,
    }

    
    oid := bson.ObjectIdHex(Id)
    // Write the user to mongo
    if err := uc.session.DB(DB_NAME).C(DB_COLLECTION).UpdateId(oid, &user); err != nil {
		checkErrTypeOne(err,"Update failed","403",c)
		
	}

    return user
}


