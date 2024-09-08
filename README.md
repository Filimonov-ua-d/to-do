# Test_task
This is a simple Golang application with PostgreSQL database and migrations using Goose.

## Installation
To start the application, you should clone the repository and install the required dependencies:

```bash 
$ git clone https://github.com/Filimonov-ua-d/test_task.git
```
```
$ cd test_task
```
```
$ go mod tidy
```
```
$ go install github.com/pressly/goose/v3/cmd/goose@latest
```

Next, you'll need to set up the database. Create a new PostgreSQL database and user, and grant the user permission to access the database. You can do this using the psql command-line tool, or any other PostgreSQL management tool:
```
CREATE DATABASE test_task;
```
```
CREATE USER test_task_user WITH PASSWORD 'your_password';
```
```
GRANT ALL PRIVILEGES ON DATABASE test_task TO test_task_user;
```

When the database is up and running, you need to apply the migrations using the Goose CLI tool.
To do that, run the following command from the project's root directory:
```bash
goose -dir ./ postgres "user=postgres password=postgres dbname=test_task sslmode=disable" up
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
```
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzEwMzgyMjQuNzQ0MzI0MiwidXNlc
  iI6eyJJRCI6IjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsIlVzZXJuYW1lIjoiemhhc2hrZXZ5Y2giLCJQYXNzd
  29yZCI6IjQyODYwMTc5ZmFiMTQ2YzZiZDAyNjlkMDViZTM0ZWNmYmY5Zjk3YjUifX0.3dsyKJQ-HZJxdvBMui0Mz
  gw6yb6If9aB8imGhxMOjsk"
} 
```

### POST /upload-images

This endpoint is used to upload images. The following request body is expected:

#### Example Input
```
form-data: file
```
```
file: image file
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
{
  "url": "http://localhost:8080/images/<filename>"
}
```

### GET /images
You need to add a header with the

#### Example Input
```
key 'Authorization'
``` 
```
value 'Bearer <JWT TOKEN>'
```

##### Example Response: 
```
{
    "Images": [
        {
            "id": 1,
            "user_id": 1,
            "image_path": "uploads/1234.png",
            "image_url": "http://localhost:8080/uploads/1234.png"
        }
    ]
}
```
