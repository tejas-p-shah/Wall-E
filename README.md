# Wall-E
Twitter clone in golang.


## About The Application

This is a Blogging web application. Through this application, a user can post, comment or respond on another user's wall. 


### Built with

Used Tools and Libraries :

* GoLang
* Go/Html Templates and Bootstrap for GUI
* MongoDB
* Github OAuth2 for Authentication.

### Prerequisites

To get this applcation running, following tools are required to be installed :

* Docker
* GoLang

### How to start the Application

* Check if GOPATH Env Variable is set.
* Clone this repository in $GOPATH/src/github.com/
* Get Mongo image :  
```
$ sudo docker run -d -p 27017:27017 -v ./mongodb:/data/db --name walle-mongo-container mongo:latest
```
* Start mongo Docker :
```
$ sudo docker start walle-mongo-container
``` 
* Run Application :
```
$ go run main/main.go
```




