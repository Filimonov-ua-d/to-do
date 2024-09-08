# to-do app
This is a simple Golang application with PostgreSQL database.

## Installation
To start the application, you should clone the repository and install the required dependencies:

```bash 
$ git clone https://github.com/Filimonov-ua-d/to-do.git
```
```
$ cd to-do
```
```
$ go mod tidy
```

Next, you'll need to set up the database. Create a new PostgreSQL database and user, and grant the user permission to access the database. You can do this using the psql command-line tool, or any other PostgreSQL management tool:
```
CREATE DATABASE to-do;
```
```
CREATE USER to_do_user WITH PASSWORD 'your_password';
```
```
GRANT ALL PRIVILEGES ON DATABASE to-do TO to_do_user;
```

## Usage

To start the application, run the following command:
```bash
$ go run cmd/main.go
```

By default, the application will listen on port 8080.

## API Endpoints

The application provides the following API endpoints:

### POST /login

This endpoint is used to authenticate users. The following request body is expected:

##### Example Input: 
```json
{
  "username": "test",
  "password": "123456"
}
```

If the provided credentials are correct, the server will respond with a JWT token:

##### Example Response: 
```json
{
	"token": "eyJhbGcimYmY5Zjk3YjUifXgw6yb6If9aB8imGhxMOjsk"
} 
```

### POST /create_task

This endpoint is used to create task. The following request body is expected:

#### Example Input
```json
{
    "title": "test2",
    "description": "test2"
}
```

Also, you need to add a header with the

#### Example Input
```
key 'Authorization'
``` 
```
value 'Bearer <JWT TOKEN>'
```

##### Example Response: 
```
201 CREATED
```

### GET /tasks
You need to add a header with the

#### Example Input
```
key 'Authorization'
``` 
```
value 'Bearer <JWT TOKEN>'
```

##### Example Response: 
```json
200 OK

[
    {
        "id": 1,
        "title": "test",
        "description": "",
        "is_done": true
    },
    {
        "id": 3,
        "title": "test2",
        "description": "test2",
        "is_done": false
    }
]
```

### GET /tasks/:id
You need to add a header with the

#### Example Input
```
key 'Authorization'
``` 
```
value 'Bearer <JWT TOKEN>'
```

##### Example Response: 
```json
200 OK

{
    "id": 1,
    "title": "test",
    "description": "",
    "is_done": false
}
```

### PUT /tasks/:id
You need to add a header with the

#### Example Input
```
key 'Authorization'
``` 
```
value 'Bearer <JWT TOKEN>'
```

#### Body
```json
{
    "id": 1,
    "title": "test",
    "desctiption": "test",
    "is_done": true
}
```

##### Example Response: 
```json
200 OK
```

### DELETE /tasks/:id
You need to add a header with the

#### Example Input
```
key 'Authorization'
``` 
```
value 'Bearer <JWT TOKEN>'
```

##### Example Response: 
```json
200 OK
```
