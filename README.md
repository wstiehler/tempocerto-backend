| Branch                                                                                                                                           Code coverage |
| ------ | :------------------------------------------------------------------------------------------------------------------------------------------------------: | -------------------------------------------------------------------------------------------------------------------------------------------------------: |
| main | [![coverage report](https://github.com/wstiehler/tempocerto-backend/blob/main/covarage.txt)](https://github.com/wstiehler/tempocerto-backend/blob/main/covarage.txt) |


<p align="center"><img src="https://user-images.githubusercontent.com/1092882/86512217-bfd5a480-be1d-11ea-976c-a7c0ac0cd1f1.png" alt="goapp gopher" width="256px"/></p>

# tempocerto-backend
This project aims to create a microservices ecosystem for managing user CRUD operations, accommodating various roles.

## Execution Guide

To run this code, you need to have Makefile and Docker installed on your system.
***You must have ports free to run the APIs***

* 3001 (documentation), 
* 8080(tempocerto-api)

***Documentation and Function Requeriments***

The documentation has been generated using C4Builder tools. To access the documentation, simply execute the following command:

```bash
   make view-doc
```

### Steps for Execution

1. Clone this repository to your local machine.

2. In the terminal, navigate to the project root directory.

3. Run the following command to launch the application:

```bash
   make dev-start-with-db
```

After executing the above command, the documentation will be accessible at [http://localhost:3001](http://localhost:3001).


***To use the application methods, follow the instructions below:***

1. Open the Postman.
   
2. Import the provided request collection file (postman_collection.json).

3. [Link to insominia documentation and step by step to run](https://documenter.getpostman.com/view/31816718/2sA2xjyqTE)


***Local Tests***

inside the application folder, run the command (Is necessary that local docker run):
- Unit test
```bash
   make test-unit-verbose
```
----
- E2E tests
```bash
   make test-e2e-local
```

<hr>

## Directory structure

```bash
.
|____docs
| |____docs
|____cmd
| |____executor
| | |____main.go
|____scripts
|____tests
| |____e2e
|____internal
| |____environment
| |____infrastructure
| | |____database
| | |____logger
| |____api
| | |____middlewares
| | |____routes
| |____domain
| | |____domain_app
|____docs

```
<hr>

![Screenshot](/docs/docs-png/arch.jpg)