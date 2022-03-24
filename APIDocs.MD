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
    "raw_otp": "82743",
    "created_at": "2022-02-13",
    "otp_created_at": "2022-02-13", 
    "first_name": "Lee Nover",
    "dob": "2022-02-13",
    "location": "",
    "profile_photo": "https://linktophotos.com/129405"
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
## Errors

<details>
<summary> 400 StatusBadRequest </summary>

```
{
"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD"
}
```
This error is returned when the service is unable to unmarshal JSON requests. Depending on where this error is encountered, take a look at the expected request format of the handler in the docs. [start here](#create-user-account)

```
{
"message": "user already exists, otp sent to user"
}
```
At the `/createAccount` endpoint, if a user attempts to create an account with a phone number that already exists in the database, this error is thrown. To resolve this, make the request with a unique phone number or try to login instead
```
{
"message": "must provide a phone number. fields cannot be left empty"
}
```
This error is returned when the client does not provide a `phone_number` when trying to create an account. To resolve this, provide a valid a `phone_number` and `country_code`
```
{
"message": "must provide an OTP and an ID. fields cannot be left empty"
}
```
This error message is received after a `POST` request is made to the `verifyOTP` endpoint without an `ID` or `OTP`. Make sure these params are provided in the request and make the call again
</details>

<details>
<summary> 401 StatusUnauthorized </summary>

```
{
"message": "otp has expired. generate a new one"
}
```
This error is returned when a user tries to verify an account with an expired `OTP`(OTPs are valid for only 3 minutes). To fix this, try to login again to generate a new `OTP`
```
{
"message": "failed to verify otp"
}
```
When this message occurs, it signals that user access has been restricted because a wrong `OTP` was submitted. Repeat the process with the correct code to fix this
```
{
"message": "failed to generate jwt"
}
```
Service unable to generate jwt

</details>


<details>
<summary> 500 StatusInternalServerError </summary>

```
{
"message": "no user found with that phone number"
}
```
Issue while trying to make a query to the database. repeating the process should fix this, contact someone on the team if problem persists
```
{
"message": "failed to create otp"
}
```
Service encountered a problem trying to generate and assign an otp to the user. repeating the process should fix this, contact someone on the team if problem persists
```
{
"message": "failed to hash otp"
}
```
Service unable to hash the user's assigned otp. repeating the process should fix this, contact someone on the team if problem persists
```
{
"message": "failed to update otp for existing user"
}
```
unable to overwrite the user's old verification code. repeating the process should fix this, contact someone on the team if problem persists
```
{
"message": "failed to create user"
}
```
unable to create and store user instance. repeating the process should fix this, contact someone on the team if problem persists
```
{
"message": "no user found with that id"
}
```
a 500 error will be thrown when a database query fails, typically when queries to the database return empty. errors related to this will take the format of the reponse above
```
{
"message": "failed to send OTP using SMS"
}
```
special 500 errors could be sent on the `verifyOTP` endpoint if SMS service fails. repeating the process should fix this, contact someone on the team if problem persists
</details>