# Tinderclone API
Backend service for tinder clone

# How to run
- Clone `.env.local` into a `.env` file
- Copy the migration in `internal/core/migration/que.sql` and run it on your db. (We can find a better way later)
- Replace the variables accordingly to match your system
- Run `go run cmd/main.go` in your terminal

# API Endpoints

`GET /healthCheck`
+ returns a `StatusOk(200)` if the server is up and running.

`POST /createAccount`
+ returns data in this format if a `phonenumber` post request was succesful.
```
{   
    "message": "phone number succesfully added",
    "data": 
        {
        "id":     id,
        "digits": digits 
        }
}
```
