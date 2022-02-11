# API Endpoints

##  Health check

**Request**

+ `GET /healthCheck`

**Response**

+ Returns a `StatusOk(200)` with an `ok` response if the server is up and running

## Create User Account

**Request**

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

**Request**

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
"message": "otp code verified"
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
This error is returned when the service is unable to unmarshal JSON POST requests. Take a look at the ideal JSON POST formats [here](#create-user-account)

```
{
"message": "user already exists"
}
```
In `create_account.go`, if a user attempts to create an account with a phone number that already exists in the database, this error is thrown. Make the request with a unique phone number or login instead

## 500

+ `StatusInternalServerError`
```
{
"message": "no user found with that id"
}
```
a 500 error will be thrown when a database query fails, typically when queries to the database return empty. errors related to this will take the format of the reponse above