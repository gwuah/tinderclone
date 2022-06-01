# API Endpoints and Error Status Codes

Table of Contents:

- [API Endpoints and Error Status Codes](#api-endpoints-and-error-status-codes)
  - [API Endpoints](#api-endpoints)
    - [Health check](#health-check)
    - [Create User Account](#create-user-account)
    - [Verify User OTP](#verify-user-otp)
    - [Update User Account](#update-user-account)
    - [Get User Location](#get-user-location)
    - [Get User Data](#get-user-data)
  - [Error HTTP Status Codes](#error-http-status-codes)
    - [400](#400)
    - [401](#401)
    - [500](#500)


## API Endpoints

### Health check

**Request**

- `GET /healthCheck`

**Response**

- Returns a `StatusOk(200)` with an `ok` response if the server is up and running

### Create User Account

**Request**

- `POST /createAccount`
  
- `Content-Type:` application/json
  
- `JSON body:` 


```json
{
"country_code": "+234",
"phone_number": "090677114589"
}
```

|     **Key**    | **Type** | **Remarks** |
|:--------------:|:--------:|:-----------:|
| `country_code` | `string` | required  |
| `phone_number` | `string` | required  |


**Sample Response**

```json
{
"data": {
    "id": "ec13c712-e796-463a-9fe1-c823db96d8ab",
    "phone_number": "23733839347",
    "otp": "$2a$10$HG2Ln6EvaK7BQoC3Kpdppeqhrq0xuXVIiawr8fif.Gvedi7fX4lXS",
    "created_at": "2022-02-13T10:30:09.9675082+01:00",
    "otp_created_at": "2022-02-13T10:33:09.9675082+01:00"
},
"message": "user succesfully created"
}
```

### Verify User OTP

**Request**

- `POST /verifyOTP`

- `Content-type:` application/json

- `JSON body:`

```json
{
"id": "ec13c712-e796-463a-9fe1-c823db96d8ab",
"otp": "84165"
}
```

|     **Key**    | **Type** | **Remarks** |
|:--------------:|:--------:|:-----------:|
| `id` | `uuid` | required  |
| `otp` | `string` | required  |


**Sample Response**

```json
{
"message": "otp code verified",
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUzNTAxNTAsInN1YiI6IjM1YjA4NDA1LWVmZTMtNGM4Zi1iNjlhLTg3YTk2NWI3MjgxMSJ9.CT-i6XpqlTd_MlwoutLXMA55nYYWQx0x-P0zWUgvKf8",
"data": {
    "id": "35b08405-efe3-4c8f-b69a-87a965b72811",
    "otp": "79298"
}
}
```

### Update User Account

**Request**

- `POST /auth/updateAccount`

- `Content-type:` application/json

- `JSON body:`

```json
{
"id": "a704ad75-149c-4ea5-b799-b1d0f1d4a296",
"first_name": "kwaku",
"dob" : "26/04/2000",
"location": {"longitude":164.7541232424 , "latitude":-77.508333},
"interests": ["football", "basketball", "biking"]
}

```

|     **Key**    | **Type** |
|:--------------:|:--------:|
| `id` | `uuid` | required  |
| `first_name` | `string` |  
| `dob` | `string` |  
| `gender` | `string` |  
| `bio` | `string` |  
| `location` | `json object` |  
| `last_name` | `string` |  
| `interests` | `[]string` |  
| `profile_photo` | `string` |  


**Sample Response**

```json
{
"data": {
"id": "9763e6a7-a8ae-438c-868d-16252bb7c8d9",
"country_code": "",
"phone_number": "",
"otp": "",
"raw_otp": "",
"created_at": "0001-01-01T00:00:00Z",
"otp_created_at": "0001-01-01T00:00:00Z",
"first_name": "kwaku",
"last_name": "",
"bio": "",
"location": {
"longitude": 164.7541232424,
"latitude": -77.508333
},
"dob": "2000-04-26T00:00:00Z",
"gender": "",
"interests": "football,basketball,biking",
"scores": {
"firstname": 0,
"lastname": 0,
"location": 0,
"bio": 0,
"gender": 0,
"dob": 0,
"interests": 0,
"profile_photo": 0
}
},
"message": "user successfully updated"
}
```

### Get User Location

**Request**

- `POST /getLocation`

- `Content-type:` application/json

- `JSON body:`
  
```json
{
"latitude": "1.2223424",
"longitude": "-1.232427",
}
```

**Sample Response**

```json
{
"data":
}
```


### Get User Data

**Request**

- `GET /getUser/<userID>`

**Sample Response**

```json
{
"message": "user successfully retrieved",
"user": {
  "id": "9763e6a7-a8ae-438c-868d-16252bb7c8d9",
  "country_code": "",
  "phone_number": "+233550156396",
  "otp": "",
  "raw_otp": "99832",
  "created_at": "2022-05-01T10:16:47.582649Z",
  "otp_created_at": "0001-01-01T00:00:00Z",
  "first_name": "kwaku",
  "last_name": "og",
  "bio": "cool kid",
  "location": {
    "longitude": 164.75412324,
    "latitude": -77.508333
  },
  "dob": "2000-04-26T00:00:00Z",
  "gender": "male",
  "interests": "['football', 'basketball'. 'biking']",
  "score": {
    "firstname": 5,
    "lastname": 5,
    "location": 15,
    "bio": 5,
    "gender": 20,
    "dob": 15,
    "interests": 10,
    "profile_photo": 25
  },
  "profile_photo": "http://someurl"
}
}
```

***

## Error HTTP Status Codes

 **Error Codes**                              | **Meaning**
----------------------------------------------|------------------------------------------------------------------------------
 **[400 or StatusBadRequest](#400)**          | Malformed parameters or other bad requests to the server.
 **[401 or StatusUnauthorized](#401)**        | Requested resource denied due to a lack of valid authentication credentials. 
 **[500 or StatusInternalServerError](#500)** | Processing errors, probably relating to backing services.

### 400

`StatusBadRequest`

```json
{
"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD"
}
```

This error is returned when the service is unable to unmarshal JSON POST requests. Take a look at the ideal JSON POST formats [here](#api-endpoints). Some other error messages attached to this code are:

```json
{
"message": "must provide a phone number. field cannot be left empty",
}
```

```json
{
"message": "must provide an OTP and an ID. fields cannot be left empty",
}
```

### 401

`StatusUnauthorized`

```json
{
  "message": "otp has expired. regenerate a new one"
}
```

The 401 error is issued when a client without accurate authentication credentials tries to request a private service/resource. It will typically take the the form of the error message above. Additional error messages of this type include:

```json
{
"message": "failed to verify otp"
}
```

```json
{
"message": "not authorized"
}
```

### 500

`StatusInternalServerError`

```json
{
"message": "no user found with that id"
}
```

500 errors will be returned when the service is unable to fulfill a client's request, this can be related to faulty (or empty) database requests, processing errors, or faults from third party services being used -- We employ third party services to handle sending OTP codes to users over SMS and to cache our user's feeds. Errors related to the database will typically take the format of the response above. Some examples of error messages related to processing and thrid party issues are:

```json
{
"message": "failed to hash otp"
}
```

```json
{
"message": "failed to queue sms otp"
}
```

```json
{
"message": "failed to send OTP using SMS"
}
```
