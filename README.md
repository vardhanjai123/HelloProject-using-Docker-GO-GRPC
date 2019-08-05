# HelloProject-using-Docker-GO-gRPC
1. This is a basic containerised Hello world Client Server using GO.
2. I used gRPC protocol for data communication between client and server.
3. There are two containers one for client and one for server.
4. The client sends hello message every 5 seconds to server which it prints on its console output.
5. I wrote my code using GOlang.

# Initial Steps to run the Project
1. Download all the files from my repository
2. Start the Docker Quickstart terminal
3. Copy all the directory in the same hierarchy in one folder.Run commands from `HelloProject` directory
4. Run command `docker-compose build`.This builds both the dockerfile for client and server.
5. Run command `docker-compose up`.This is similar to running each docker image for client and server.

# Installations Required for starting the project
1. Install a Docker Desktop for Windows or Mac OS.But,if you don't have Windows 10,then Install Docker Toolbox.
2. Install protoc compiler for compiling protobuf messages.
3. Install Golang

# Basic Knowledge required
1. Basics of Golang
2. Basic understanding of gRPC like its advantages over xml,json
3. Basics understanding of docker technologies

# Acknowledgements
1. Followed Tutorials point for docker basic commands
2. Followed the freecodecamp for understanding how to implement it
3. Followed the Tensor Programming lectures to understand compilation commands and format of proto files
4. Followed Medium lecture to get concepts of gRPC working,its advantages over its counterpart xml,json
5. Followed official documentation of docker for pushing images to my docker hub
6. Followed Edureka tutorials to get understanding of docker and basics of GO
7. Took help from my friend Ashwini Abhishek whenever i got stuck up.

