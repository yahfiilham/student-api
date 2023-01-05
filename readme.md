## Overview

student API is a simple RESTful API using golang without any framework for routing and dmbs (database engine). This repository implemented golang [standards layout](https://github.com/golang-standards/project-layout)

It includes:

|             |                                 |
| ----------- | ------------------------------- |
| function    | Health Check                    |
| endpoint    | GET http://{host}:{port}/health |
| description | Check health server             |

|             |                                   |
| ----------- | --------------------------------- |
| function    | Register Student                  |
| endpoint    | POST http://{host}:{port}/student |
| description | Register student into the system  |
| payload     | { name: “budi”, age: 5 }          |

|             |                                  |
| ----------- | -------------------------------- |
| function    | List Student                     |
| endpoint    | GET http://{host}:{port}/student |
| description | List student                     |

|             |                                       |
| ----------- | ------------------------------------- |
| function    | Get Student                           |
| endpoint    | GET http://{host}:{port}/student/{id} |
| description | Get student by ID                     |

|             |                                       |
| ----------- | ------------------------------------- |
| function    | Update Student                        |
| endpoint    | PUT http://{host}:{port}/student/{id} |
| description | Update student by ID                  |
| payload     | { “name: “budi kurniawan” }           |

|             |                                          |
| ----------- | ---------------------------------------- |
| function    | Delete Student                           |
| endpoint    | DELETE http://{host}:{port}/student/{id} |
| description | Delete student by ID                     |

## Prerequisite:

- Install golang
- Clone this [repository](https://github.com/yahfiilham/student-api)
- Install dependencies
  > go mod tidy
- To run the repository locally :
  > make run
- To run the coverage of unit test :
  > make coverage

---

[postman collection](https://api.postman.com/collections/17037134-0b2b3ef7-66c1-4b0a-8d8b-c88faedcf7b5?access_key=PMAT-01GP1VC237XE84BEF7HZEYNQR1)
