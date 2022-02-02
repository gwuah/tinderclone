# Tinder clone

This clone mirrors Tinder's geosocial networking application. It allows users anonymously like or dislike other users' profiles, which include their photos, a short bio, and a list of their interests. It extends its features by implementing its own custom matching algorithm. However, this clone does not include a messaging feature, *It leverages some other anonymous chat service.


# Development Usage (Installing)

- Clone `.env.local` into a `.env` file
    - Replace the variables accordingly to match your system
- Copy the migration in `internal/core/migration/que.sql` and run it on your db. (We can find a better way later)
- Run `go run cmd/main.go` in your terminal

# API Endpoints

##  Health check

**Request**

+ `GET /healthCheck`

**Response**

+ Returns a `StatusOk(200)` with an `ok` response if the server is up and running.

## Create User Account

**Sample Request**

+ `POST /createAccount`

+ `Content-Type`: application/json

+ `JSON body:` 
```
{
   "phone_number" : "1234567"
}
```

**Sample Response**
```
{
   "data": {
            "id": 2,
           "phone_number": "1234567"
           },
   "message": "phone number succesfully added"
}
```
