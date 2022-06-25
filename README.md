# Microservices in Golang and creating Docker image for easy deployment

## Why Microservices Architecture?
Monolithic architecture is considered to be a traditional way of building applications. A monolithic application is built as a single unit. It has a huge code base and it lacks modularity, due to which it is hard to scale the application. When it comes to implementing new technology the app is needed to be rewritten.

Microservices is a method of developing software in which an application is broken down into multiple core functions called services that are loosely coupled with each other and work together but are not dependent on each other. The services communicate among themselves using lightweight protocols.

The benefit of using microservices is that the development team can rapidly build new features and update existing features of apps without harming existing features. Microservices can be deployed, maintained, updated, and scaled independently of each other.

![0_-e77N6QTfAe0G4NG](https://user-images.githubusercontent.com/49842473/174866644-4f8cf2cb-a575-43eb-ac77-92c13ba09796.png)

## Why Golang?
Go is a fast, statically typed compiled language that feels like a dynamically typed interpreted language. It is meant to be a simple but powerful language. Go language’s syntax is simple like C language and learning it is easy. It directly converts the code into machine language. Go, has a built-in garbage collector, and it is used to build and scale cloud computing systems and web applications and is built for back-end applications. In a nutshell, Go is the go-to language for faster, scalable, and reliable development.

## Why Docker?
Docker is a tool to containerize applications. While developing an application, we work on a device having a particular Operating system and use a specific version of the programming language with its dependencies. While deploying the application we might not have the same environment available to run, and here comes the concept of Containerisation. Containers are the light-weight virtual environment created to run applications on them.

The virtual machine is an emulation of a whole computer system. In VMs we separate the resources from physical machines so that they can be dedicated only to VMs. It acts like a physical server and hence it is not needed to run a single service.

A Container is just an emulation of OS. It contains all the dependencies and libraries needed to run the application. All the required things are preserved in an image, and the image contains all the binaries. It is very small compared to a VM and it runs like a normal application on the machine. With containerization, developers don’t need to write applications on different machines having different configurations.

![0_QkvhJ7zFiu6oitql](https://user-images.githubusercontent.com/49842473/174866876-a66a17a3-d5dd-4c34-93b2-7ff64860f82b.png)

## Required Installations:
_Note: This article is based on Ubuntu 20.10 and Go 1.14.7_

**A. Installing Go:**

1. Download Golang from the website [golang.org].

2. Extract the archive you downloaded into /usr/local, creating a Go tree in /usr/local/go:
```
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.14.7.linux-amd64.tar.gz
```

3. Add /usr/local/go/bin to the PATH environment variable:
```
export PATH=$PATH:/usr/local/go/bin
```

4. Verify that Go is installed:
```
go version
```

**B. Installing Docker**

1. Update apt package
```
$ sudo apt update
```

2. Remove old version of Docker:
```
sudo apt-get remove docker docker-engine docker.io
```

3. Install Docker:
```
sudo apt install docker.io
```

4. Verify the installed Docker version:
```
docker --version
```

5. Start docker service and verify:
```
sudo systemctl start docker
sudo systemctl enable docker
sudo systemctl status docker
```

![1_mk_NSed6FbBRZysP2_fVIg](https://user-images.githubusercontent.com/49842473/174867346-3d817a3d-dadc-409d-b0ce-99a610fc01a0.png)

6. Create an account on Docker Hub for uploading the Docker image we will create.


## Building microservice:

Open a terminal and create a project folder

```
mkdir microservice
cd microservice
```
Enable dependency tracking for your code
```
go mod init microservice
```

![1_0PagT68fN_tjDnrCuUHIYw](https://user-images.githubusercontent.com/49842473/174869395-c1ee404b-aa27-48f6-b5cc-4415505764d8.png)

Open the folder in your favorite text editor and create a main.go file inside it.

Write a simple Hello World Program and try to run it.
```
go run main.go
```

![1_qWOzb5J1sZVy4c7jcmN5QA](https://user-images.githubusercontent.com/49842473/174869482-7e54202d-4395-4a3c-a748-252aa1ebe317.png)

If it runs properly then we are ready to go.



We are building a microservice that will return data as JSON.

So now we must define a type which is the data that our service is going to return.

In the project, create a models directory which will contain definitions of the kinds of data returned and add file person.go in that. We are returning the list of person objects.

```
package models

//Person is a type
type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

In the ***models*** package, we are defining the person model. It has an ID as an integer, Name as a string, and Age as an integer. Annotations on the right side of attributes are here to specify how this data will be in JSON format.

For now, we need to create a package that has code that will be used to add and return data from the service. Let’s name it repository.

Create ***repository.go***

```
package repository

import (
	"microserv/models"
)

var people []models.Person
var nextID = 1

func GetDetails() []models.Person {
	return people
}

func AddPerson(person models.Person) int {
	person.ID = nextID
	nextID++
	people = append(people, person)

	return person.ID
}
```

First, we imported our models package where we defined our data model.

Then created an array people of type Person. This is a simple example so we are not using any database service and just an array.

We need to return data from the database for that create GetPerson() function which will return currently available data.

For adding data, create AddPerson() function, in which we will append newly arrived data in the people array.

Back to main.go, we will use http protocol.

For our service, we want main function to listen to a network port and keep listening to the incoming requests.

```
package main

import (
	"log"
	"microserv/handlers"
	"microserv/models"
	"microserv/repository"
	"net/http"
)

func main() {
	repository.AddPerson(models.Person{
		Name: "Mike",
		Age:  27,
	})

	http.HandleFunc("/", handlers.HandleRequest)

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
```

Add data of person using AddPerson function from repository package, so when we will run application, we will be able to confirm our application is working or not.

Using the HTTP package from Go standard library, we will listen and serve on port number 8082, Then check if any error has occurred or not. It will return an error message if any.

We need to set up our routing, from HTTP package using hadleFunc() function. Here we will specify URL, which in our case is “/” means root, and then call our handler function. So any request comes in to root just call our handler function.

When a request comes in, we want to redirect according to the request. As we will receive ‘Get’ and ‘Post’ requests, we will process accordingly.

If someone makes GET request we will return a list of people as JSON, and if someone makes POST request we will receive a JSON of details of the person and add it to our data.

Create a handlers directory in which we are going to create a handlers.go file. Next, we’ll be writing functions in it which will handle received requests and display or add data accordingly.

```
package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"microserv/models"
	"microserv/repository"
	"net/http"
)

// handler is here
func HandleRequest(w http.ResponseWriter, r *http.Request) {

	log.Println("Request recived: ", r.Method)

	switch r.Method {
	case http.MethodGet:
		display(w, r)
		break
	case http.MethodPost:
		add(w, r)
		break
	default:
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		break
	}

}

func display(w http.ResponseWriter, r *http.Request) {
	products := repository.GetDetails()
	json, _ := json.Marshal(products) //ignored the error

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)

	log.Println("Response returned: ", 200)

}

func add(w http.ResponseWriter, r *http.Request) {
	payload, _ := ioutil.ReadAll(r.Body)

	var person models.Person
	err := json.Unmarshal(payload, &person)
	if err != nil || person.Name == "" || person.Age == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}
	person.ID = repository.AddPerson(person)

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(201)

	json, _ := json.Marshal(person)
	w.Write(json)

	log.Println("Response returned: ", 201)
}
```

Creating function which is going to receive HTTP response writer and http request. http.Request will give access to what request is received and http.ResponseWriter will allow us to write an outgoing response.

First, check the received HTTP request, then we will switch based on the request. If it is a GET request call display function that will write existing data as JSON, and if it is POST request then call add function which will insert data in the database.

That is it, we created our first Microservice in Go!

Now it's time to start service. Open terminal and run:
```
go run main.go
```

![1_VcJtLT6q7oDd8xKWNl8XPQ](https://user-images.githubusercontent.com/49842473/175775051-8227977e-81c8-4a30-9cb9-c8eb706a9fe4.png)

Open browser and in search bar type localhost:8082(because we used port 8082), you can see JSON data which we already added.

![1_QwmLAt4SlPUmDCUP-Ubw_Q](https://user-images.githubusercontent.com/49842473/175775063-785dca4f-5407-4804-9778-7c556ad24290.png)

Now we will try to make a POST request using Postman.

![1_kDQtyEmfu9dFC4RkY27HPw](https://user-images.githubusercontent.com/49842473/175775072-fdb29ca8-ba82-41f3-85b3-dde685435249.png)

Refresh the page. We successfully added new data.

![1_yXpNz3hmhKQ1iawL-G19gQ](https://user-images.githubusercontent.com/49842473/175775120-2967a9bf-1848-48d7-acaf-94df7d20345f.png)

Good Job!!! Your microservice is running flawlessly.


## Creating Docker Image:

Now that our service is running it is time to create its docker image so we can easily deploy it anywhere on the cloud server.

***A.*** Create Dockerfile in the project folder:
```
FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o microserv .
FROM alpine
RUN mkdir /app
ADD . /app/
COPY --from=builder /build/microserv /app
WORKDIR /app
CMD ["./microserv"]
```

1. Starting with the Golang Alpine image, this is just an image that already exists on the docker hub. It contains everything that we need to build and run the application written in golang.
2. Then we create a directory in our image.
3. We will copy the build directory in the Docker image.
4. Then define the working directory. It specifies that any command will be executed in this directory.
5. Now for executing the command for building our application use RUN and create an executable file of the app.

Now building part is over, we do not need to build an application every time we want to run it. Now we have everything we need to be able to build our app.

6. Now we will use alpine(Alpine is a small Linux distribution) distribution image.
7. Then repeat the process of creating and defining working directory for app.
8. After that copy the executable file we created earlier to this image
9. Now tell it to run our application using CMD.

Here, the image we will end up with will only contain what is necessary to execute the application.

![1__ouU39K66HawJBnZ_W2sbg](https://user-images.githubusercontent.com/49842473/175775280-4ca7abaf-65ff-4340-8cd3-d2e47899b47f.png)

***B.*** Build the Docker image:
Open project directory in terminal and run the following command:
```
sudo docker build -t [your doker hub username]/[name of your image]:[tag]
```

![1_g_X32q7ELlybCBOAB88b9g](https://user-images.githubusercontent.com/49842473/175775322-09724e0d-5baf-42e8-bd9c-6a36bed0aea7.png)

Our Docker image is built. Try to run it using the command:
```
sudo docker run -p [port]:[port] [your doker hub username]/[name of your image]:[tag]
```

![1_WTPwWG3agdMvZm6BWgkCcA](https://user-images.githubusercontent.com/49842473/175775345-ab114a6d-3733-441e-b161-943fee351489.png)

And it is working!!

***C.*** Pushing image on Docker Hub
1. Go to Docker hub and create a repository.

![1_7jiJvVz60Tnxn0ovbLsYHg](https://user-images.githubusercontent.com/49842473/175775376-8cbb55be-33e7-4e47-888d-e1bc49ccd5fe.png)

2. Open terminal and login to Docker:
```
sudo docker login
```

![1__ouU39K66HawJBnZ_W2sbg](https://user-images.githubusercontent.com/49842473/175775401-e02f4ee8-8331-46a2-b2b3-ea09c7baaeda.png)

3. Push docker image on docker hub repository:
```
sudo docker push [your doker hub username]/[name of your image]:[tag]
```

![1_FQFgvqTLU0r_iqTYwhvVXQ](https://user-images.githubusercontent.com/49842473/175775428-56fffc48-9baa-4f83-91c4-da629d360b25.png)

![1_CXyfZP5HuVQ451fBhMe4Ng](https://user-images.githubusercontent.com/49842473/175775436-94689e40-1fde-4cb1-ba18-633a6b56fc01.png)

DONE!!!

Now we can use this image to deploy my service on the cloud or run it on any device having docker installed. Just type command:
```
docker pull [image name]:[tag]
```
And run it as we did earlier.

## Project Tree:

![1_1SOvt-cMm-QdkD2VyL2EwA](https://user-images.githubusercontent.com/49842473/175775483-549dbd56-dbdb-4761-b255-34579750c019.png)





