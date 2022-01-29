# Tinder clone

This clone mirrors Tinder's geosocial networking application. It allows users anonymously like or dislike other users' profiles, which include their photos, a short bio, and a list of their interests. It extends its features by implementing its own custom matching algorithm. However, this clone does not include a messaging feature, *It leverages some other anonymous chat service.


# Development Usage (Installing)

- Clone `.env.local` into a `.env` file
    - Replace the variables accordingly to match your system
- Copy the migration in `internal/migration/que.sql` and run it on your db. (We can find a better way later)
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

+ `Content-Type:` application/json

+ `JSON body:` 
```
    {
    "phone_number": "1234567"
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

## Verify User OTP

**Sample Request**

+ `POST /verifyOTP`

+ `Content-type:` application/json

+ `JSON body:`
```
    {
    "id": 1,
    "otp": "12647"
    }
```

**Sample Response**
```
    {
    "message": "OTP code verified."
    }
```
# Errors

## 400

+ `StatusBadRequest`
```
    {
    "message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD"
    }
```
This error is returned when the service is unable to unmarshall JSON POST requests. Take a look at the ideal JSON POST formats [here]()

```
    {
    "message": "user already exists."
    }
```
In `create_account.go`, if a user attempts to create an account with a phone number that already exists in the database, this error is thrown. Make the request with a unique phone number or login instead.

## 500

+ `StatusInternalServerError`
```
    {
    "message": "no user found with that id."
    }
```
a 500 error will be thrown when a database query fails, typically when queries to the database return empty. errors related to this will take the format of the reponse above.