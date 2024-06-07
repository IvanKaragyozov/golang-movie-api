# Music API

This is a simple Web API for Music implemented in Golang. It provides RESTful endpoints for CRUD operations on music resources.

## Features
- Create, Read, Update, Delete (CRUD): Perform basic CRUD operations on music resources.

## Usage
- Get All Songs: GET /api/songs
- Get a Song by ID: GET /api/songs/{id}
- Create a Song: POST /api/songs
- Update a Song: PUT /api/songs/{id}
- Delete a Song: DELETE /api/songs/{id}

## Error Handling
- Appropriate HTTP status codes and messages are returned in case of errors such as invalid JSON data or song not found.

## Technologies Used
- Golang: Backend implementation.
- Gorilla Mux: Routing and HTTP request handling.
- UUID: Generating unique identifiers.

## How to Run
1. Ensure Go v1.22.2 is installed.
2. Clone the repository.
3. Navigate to the project directory.
4. Run `go mod download` to install dependencies.
5. Run `go run main.go` to start the server.
6. Endpoints can be accessed at http://localhost:8000/
