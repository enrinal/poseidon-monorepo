# Poseidon - Auth-SVC
Auth-SVC API

## Version: 1.0.0

### /api/v1/claims

#### GET
##### Description

API for get claims

##### Headers
```json
"Authorization" : "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZW5yaW5hbGwiLCJwaG9uZSI6IjYyODEyODE3OTQ4NDAiLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE2NTQ5NjM5MTR9.MYo7sqb-E3h0F9HsoQ3Fnm2KOSguGCZ2wQlJFKuDlIY"
```

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Example response |

```json
{
    "message": "success",
    "data": {
        "name": "enrinall",
        "phone": "6281281794840",
        "role": "admin",
        "exp": 1654963914
    }
}
```

##### Security

| Security Schema | Scopes |
| --- | --- |
| bearerAuth | |

### /api/v1/users

#### POST
##### Description

API for create users

##### Request
```json
{
    "name":"john doe",
    "phone":"6281281794899",
    "role": "admin"
}
```

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Example response |

```json
{
    "message": "success",
    "data": {
        "name": "enrinall",
        "phone": "6281281794840",
        "role": "admin",
        "password": "1vbV"
    }
}
```

### /api/v1/users/login

#### POST
##### Description

API login return token 

##### Request
```json
{
    "phone":"6281281794849",
    "password": "1SD7"
}
```

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Example response |

```json
{
    "message": "success",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZW5yaW5hbCIsInBob25lIjoiNjI4MTI4MTc5NDg0OSIsInJvbGUiOiJhZG1pbiIsImV4cCI6MTY1NDk2Mzc4MX0.Ek8E8uMmbIqlNhJ6Q5G-9xN8vI8kRMhwXno89CeuCh8"
    }
}
```


