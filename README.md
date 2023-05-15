
## Getting started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

## Presiquites
#### Golang
You need to have Go v1.19 installed on your machine. Follow the official installation guide to install Go. Or, follow managing installations guide to have multiple Go versions on your machine.

### MySQL
This service has dependency with MySQL. For development environment, you need to have a MySQL server running on your machine.

## installation

1. clone this repository :  


```http
git clone https://github.com/Alfeenn/online-learning.git
```


2. Build binaries using go.build :

go build

## Running

Execute binary to start the service :

cmd : online-learning.exe

## Note 

Run auto migration first to migrate table and set up RBAC

cmd : online-learning.exe -migrate up


## API Documentation

Use Postman API Documentation :

```http
  https://www.postman.com/solar-water-440274/workspace/test/collection/20757492-1e249602-d084-42f5-91b0-a02e727510a9?action=share&creator=20757492
```


## API Reference

#### Register account using

```http
  GET /api/register
```

#### Login

```http
  GET /api/login
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `username`      | `string` | **Required**. |
| `password`      | `string` | **Required**.  |

After successful login response will looks like this :

```http
{
    "code": 200,
    "status": "OK",
    "data": {
        "Authorization": {
            "username": "mike@gmail.com",
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im1pa2VAZ21haWwuY29tIiwiZXhwIjoxNjgzNzE3ODI5fQ.mesCiuexL8U6BvdhIqkpGuwtaYbIhT5CNR7qPi5qHhY",
            "exp": "2023-05-10T18:23:49.4467804+07:00"
        }
    }
}
```
Copy token and send to authorization section with bearer token type


Use Token to protected api

#### Protected api

```http
  /api/user
```
```http
  /api/admin
```

see full documentation :

```http
  https://www.postman.com/solar-water-440274/workspace/test/collection/20757492-1e249602-d084-42f5-91b0-a02e727510a9?action=share&creator=20757492
```

