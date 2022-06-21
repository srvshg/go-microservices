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



