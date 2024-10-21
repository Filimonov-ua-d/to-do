# total english app
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

Also add envs

## Usage

To start the application, run the following command:
```bash
$ go run cmd/main.go
```

By default, the application will listen on port 8080.

## API Endpoints

### API HOST: 
```
https://to-do-g938.onrender.com
```

The application provides the following API endpoints:

### POST /register

This endpoint is used to authenticate users. The following request body is expected:

##### Example Input: 
```json
{
    "name":"test",
    "email": "test@gmail.com",
    "password":"test"
}
```

If the provided credentials are correct, the server will respond with a JWT token:

##### Example Response: 
```json
{
	"token": "eyJhbGcimYmY5Zjk3YjUifXgw6yb6If9aB8imGhxMOjsk"
} 
```


### POST /login

This endpoint is used to authenticate users. The following request body is expected:

##### Example Input: 
```json
{
  "name": "test"
  "email": "test",
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

### POST /contact-us

This endpoint is used to save contact us req. The following request body is expected:

#### Example Input
```json
{
    "name": "Stepan",
    "phone": "+380730737373", (OPTIONAL)
    "messanger": "+380730737373", (OPTIONAL)
    "email": "StepanGiga@gmail.com",
    "message": "Give me course for free pls", (OPTIONAL)
    "course": "Basic english"
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
200 OK
```

### PUT /api/update-profile

This endpoint is used to update a user's profile. The following request body is expected:

###### Example Request:
```json
{
    "name": "John Doe",
    "email": "johndoe@gmail.com",
    "password": "123456",
    "image_url": "Advanced English"
}
```

Also, you need to add a header:

#### Example Header:
```
key: 'Authorization'
```
```
value: 'Bearer <JWT TOKEN>'
```

###### Example Response:
```json
{
    200 OK
}
```

### POST /api/upload-video

This endpoint is used to upload a new video lesson. The following request body is expected:

###### Example Request:
```json
{
    "course_id": 12,
    "url": "https://example.com/video/lesson1.mp4",
    "comment": "Introduction to the course" // (OPTIONAL)
}
```

Also, you need to add a header:

#### Example Header:
```
key: 'Authorization'
```
```
value: 'Bearer <JWT TOKEN>'
```

###### Example Response:
```json
{
    200 OK

    "id": 1,
    "course_id": 12,
    "url": "https://example.com/video/lesson1.mp4",
    "comment": "Introduction to the course" // (OPTIONAL)
}
```

### DELETE /api/delete-video/:course_id

This endpoint is used to delete ALL video lessons by course ID.

You need to add a header:

#### Example Header:
```
key: 'Authorization'
```
```
value: 'Bearer <JWT TOKEN>'
```

###### Example Response:
```json
{
    200 OK
}
```

### GET /api/video/:id

This endpoint is used to get a video lesson by course ID.

You need to add a header:

#### Example Header:
```
key: 'Authorization'
```
```
value: 'Bearer <JWT TOKEN>'
```

###### Example Response:
```json
{
    "id": 1,
    "course_id": 12,
    "url": "https://example.com/video/lesson1.mp4",
    "comment": "Introduction to the course"
}
```

### GET /api/videos

This endpoint is used to get a video lessons.

You need to add a header:

#### Example Header:
```
key: 'Authorization'
```
```
value: 'Bearer <JWT TOKEN>'
```

###### Example Response:
```json
[
{
    "id": 1,
    "course_id": 12,
    "url": "https://example.com/video/lesson1-1.mp4",
    "comment": "Introduction to the course p.1"
},
{
    "id": 2,
    "course": 12,
    "url": "https://example.com/video/lesson1-2.mp4",
    "comment": "Introduction to the course p.2"
}
]
```

### PUT /api/upload-image/:id

SEND user_id in "/:id" param.
IN RESPONSE U WILL GET BASE64 STRING

This endpoint is used to upload a new user image. The following request body is expected:

###### Example Request:
<img width="615" alt="image" src="https://github.com/user-attachments/assets/35555820-5557-4a53-85e5-4a79e221fd81">

Also, you need to add a header:

#### Example Header:
```
key: 'Authorization'
```
```
value: 'Bearer <JWT TOKEN>'
```

###### Example Response:
```json
{
    200 OK

    "encoded_file": "base64 string",
}
