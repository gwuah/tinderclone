# API Endpoints

##  Health check

**Request**
`   
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
"country_code": "+234",
"phone_number": "090677114589"
}
```


|     **Key**    | **Type** | **Remarks** |
|:--------------:|:--------:|:-----------:|
| `country_code` | `string` | `required`  |
| `phone_number` | `string` | `required`  |


**Sample Response**
```
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

## Verify User OTP

**Request**

+ `POST /verifyOTP`

+ `Content-type:` application/json

+ `JSON body:`
```
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
```
{
"message": "otp code verified",
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUzNTAxNTAsInN1YiI6IjM1YjA4NDA1LWVmZTMtNGM4Zi1iNjlhLTg3YTk2NWI3MjgxMSJ9.CT-i6XpqlTd_MlwoutLXMA55nYYWQx0x-P0zWUgvKf8",
"data": {
    "id": "35b08405-efe3-4c8f-b69a-87a965b72811",
    "otp": "79298"
}
}
```
**Request**

+ `POST /updateAccount`

+ `Content-type:` application/json

+ `JSON body:`
```
{
"id": "2fe743f9-9d1d-45ce-bb4f-cb253e248099",
"first_name": "kwaku",
"dob" : "26/04/2000",
"location": {"longitude":164.7541232424 , "latitude":-77.508333},
}

```

|     **Key**    | **Type** | **Remarks** |
|:--------------:|:--------:|:-----------:|
| `id` | `uuid` | required  |
| `first_name` | `string` | required  |
| `dob` | `string` | required  |
| `location` | `json object` | required  |


**Sample Response**
```
{
"data": {
"id": "2fe743f9-9d1d-45ce-bb4f-cb253e248099",
"country_code": "",
"phone_number": "",
"otp": "",
"raw_otp": "",
"created_at": "0001-01-01T00:00:00Z",
"otp_created_at": "0001-01-01T00:00:00Z",
"first_name": "kwaku",
"dob": "2000-04-26T00:00:00Z",
"location": ""
},
"message": "user successfully updated"
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
"message": "user already exists, otp sent to user"
}
```
On the `create_account/` endpoint, if a user attempts to create an account with a phone number that already exists in the database, this error is thrown. Make the request with a unique phone number or login instead

## 500

+ `StatusInternalServerError`
```
{
"message": "no user found with that id"
}
```
a 500 error will be thrown when a database query fails, typically when queries to the database return empty. errors related to this will take the format of the reponse above

special 500 errors could be sent on the `verifyOTP/` endpoint if SMS service fails. The response takes this format:

+ `StatusInternalServerError`
```
{
"message": "failed to send OTP using SMS"
}
```

a 500 will be thrown if user update fails on an `updateAccount/`. The response takes this format:
+ `StatusInternalServerError`
```
{
"message": "failed to update user"
}
```

a 500 will be thrown if user location update fails. The response takes this format:
+ `StatusInternalServerError`
```
{
"message": "failed to update user location"
}
```
