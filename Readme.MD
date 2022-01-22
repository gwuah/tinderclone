## 1. How to run

- Clone `.env.local` into a `.env` file
- Copy the migration in `internal/core/migration/que.sql` and run it on your db. (We can find a better way later)
- Replace the variables accordingly to match your system
- Run `go run cmd/main.go` in your terminal

## 2. API Endpoints

### 2.1 Health check

**Request**

`GET /healthCheck`

**Response**

+ returns a `StatusOk(200)` with an `ok` response if the server is up and running.

### 2.2 Create User Account

**Sample Request**

`POST /createAccount`

`Content-Type`: application/json

`JSON body:` 

    {
    "phone_number" : "1234567"
    }


**Sample Response**

    {
    "data": {
        "id": 2,
        "phone_number": "1234567"
    },
    "message": "phone number succesfully added"
    }
