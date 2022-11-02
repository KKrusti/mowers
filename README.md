# Mowers challenge

Mowers is a Golang kata. The required tasks can be found at instructions.md file in this same source level
## Challenge assumptions
The following assumptions have been considered for the mowers to behave as expected:
* Since a Mower is placed by a worker in the plateau, mowers are unable to be placed outside the plateau dimensions and/or in the same position as another mower that is already in the plateau.
* Mowers can not collide with other mowers or move outside the plateau dimensions. To avoid this, if a mower tries to move to an unavailable position, the movement will be ignored, other movements will be processed as well as they're valid.
* To avoid verifications and simplify the code, configuration file is filled correctly as described on the instructions.
  * If the file format is correct, the orientation of mower initial position is flexible. In case an incorrect orientation is given (different to NESW) it will be defaulted to N. 

## Architecture
This is the first time using hexagonal architecture. The code has been divided into 4 layers: 
* application: where the use case is found, it acts as an orchestrator between the logic found on each both domain entities.
* domain: here we can find the two identified domain entities: mower and plateau. In addition, some valueobjects can be found here too
* infrastructure: this layer receives the call from the presentation layer to execute some logic that is not from domain one.
* ui/presentation: this is a small layer that calls infrastructure to process some info and then communicates with application laywer.

TDD has been applied in domain and application layer.

## Installation

Go 1.18+ is required to run the application since it's using some functionalities introduced in this version.\
There are three ways to run the application, only the first one requires golang 1.18+

### Local mode

To build the application run the following command from the root folder of the project

```bash
go build .
```
To run the tests and check the code coverage use the following command:
```bash
go test ./... --cover
```
Place a file on  project root folder with the instructions to be followed by the application *e.g: input.txt*\
To run the application use the following command(used input.txt as example of instructions file):
```bash
go run main.go input.txt
```

### Docker mode
One of the advantages of running the project in docker is that it can run even without go 1.18+ on your computer.

### Docker
Before building the docker image, edit the **input.txt** file that is found at root source folder and write there the desired instructions.
Once this has been done you can proceed to create the docker image.\
To build the docker image execute the following command
```
docker build -t <image-name> .
```
Once the docker image is built it can be executed using the command
```
docker run <image-name>
```
If it's desired to with another configuration for the mowers, edit the **input.txt** file and proceed to build and run process. 

### Docker-compose
To make it simpler the docker container can be run using docker compose at the project root folder.\
In this case the file used is the same than the one described in the docker process. Don't forget to edit **input.txt** with desired data.\
To execute using docker-compose use the following command: 
```
docker-compose up
```
To shut down the docker container use the following command
```
docker-compose down
```
In case the image is needed to rebuild use the commands
```
docker-compose build
docker-compose up
```

