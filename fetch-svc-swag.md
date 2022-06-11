# Poseidon - Fetch-SVC
Fetch-SVC API

## Version: 0.1

### /api/v1/claims

#### GET
##### Description

API for get claims for token

##### Headers
```json
"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZW5yaW5hbGwiLCJwaG9uZSI6IjYyODEyODE3OTQ4NDAiLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE2NTQ5NjM5MTR9.MYo7sqb-E3h0F9HsoQ3Fnm2KOSguGCZ2wQlJFKuDlIY"
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

### /api/v1/commodity/aggregated

#### GET
##### Description

API for get commodity aggregated

##### Headers
```json
"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZW5yaW5hbGwiLCJwaG9uZSI6IjYyODEyODE3OTQ4NDAiLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE2NTQ5NjM5MTR9.MYo7sqb-E3h0F9HsoQ3Fnm2KOSguGCZ2wQlJFKuDlIY"
```

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Example response |

```json
{
    "message": "success",
    "data": [
      {
            "year": 2022,
            "month": 5,
            "week": 21,
            "province_area": "JAWA TENGAH",
            "size_aggregate": {
                "min": 70,
                "max": 70,
                "median": 70,
                "avg": 70
            },
            "price_aggregate": {
                "min": 5000,
                "max": 5000,
                "median": 5000,
                "avg": 5000
            }
        },
        {
            "year": 2022,
            "month": 5,
            "week": 21,
            "province_area": "JAWA BARAT",
            "size_aggregate": {
                "min": 120,
                "max": 120,
                "median": 120,
                "avg": 120
            },
            "price_aggregate": {
                "min": 80000,
                "max": 80000,
                "median": 80000,
                "avg": 80000
            }
        },
        {
            "year": 2022,
            "month": 5,
            "week": 21,
            "province_area": "ACEH",
            "size_aggregate": {
                "min": 80,
                "max": 80,
                "median": 80,
                "avg": 80
            },
            "price_aggregate": {
                "min": 50000,
                "max": 50000,
                "median": 50000,
                "avg": 50000
            }
        },
     ]
```

### /api/v1/commodity

#### GET
##### Description

Auto generated using Swagger Inspector

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Example response |

```json
{
    "message": "success",
    "data": [
        {
            "uuid": "1653364757694",
            "komoditas": "dede kendor",
            "area_provinsi": "SUMATERA BARAT",
            "area_kota": "PADANG PARIAMAN",
            "size": "40",
            "price": "222222",
            "tgl_parsed": "2022/05/30 11:1:13",
            "timestamp": "1653364757694",
            "price_usd": 15.205115310007901
        },
        {
            "uuid": "2507dbb0-4144-4507-83ab-9a602dae8f4a",
            "komoditas": "Supplier Lele",
            "area_provinsi": "JAWA BARAT",
            "area_kota": "BANDUNG",
            "size": "60",
            "price": "200000",
            "tgl_parsed": "2022/05/30 11:0:47",
            "timestamp": null,
            "price_usd": 13.684617463624576
        },
        {
            "uuid": "955ca8fc-05aa-44c8-bee8-8a5213d710c8",
            "komoditas": "flea",
            "area_provinsi": "BALI",
            "area_kota": "PANDEGLANG",
            "size": "80",
            "price": "12312",
            "tgl_parsed": "31/5/2022",
            "timestamp": "1654011903546",
            "price_usd": 0.8424250510607288
        },
        {
            "uuid": "56ecf018-74f0-4036-ba4b-fe57d726dede",
            "komoditas": "test",
            "area_provinsi": "BALI",
            "area_kota": "BULELENG",
            "size": "50",
            "price": "123",
            "tgl_parsed": "31/5/2022",
            "timestamp": "1654012271921",
            "price_usd": 0.008416039740129114
        },
        {
            "uuid": "a715fd24-5b52-4148-8193-ab9fd449215d",
            "komoditas": "awfeaw",
            "area_provinsi": "BALI",
            "area_kota": "BULELENG",
            "size": "40",
            "price": "444444",
            "tgl_parsed": "31/5/2022",
            "timestamp": "1654012824963",
            "price_usd": 30.410230620015803
        },
        {
            "uuid": "b7ab27b5-830a-4d6b-bc05-6d692be5b0ae",
            "komoditas": "fawewa",
            "area_provinsi": "DKI JAKARTA",
            "area_kota": "KOTA TUA",
            "size": "40",
            "price": "44444",
            "tgl_parsed": "31/5/2022",
            "timestamp": "1654012941633",
            "price_usd": 3.040995692766653
        },
    ]       
```

### Models

#### KomoditasAggregated

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| year | string |  | Yes |
| month | string |  | Yes |
| week | string |  | Yes |
| province_area | string |  | Yes |
| size_aggregate | object |  | Yes |
| price_aggregate | object |  | Yes |

#### Komoditas

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| uuid | string |  | Yes |
| komoditas | string |  | Yes |
| area_provinsi | string |  | Yes |
| area_kota | string |  | Yes |
| size | string |  | Yes |
| price | string |  | Yes |
| price_usd | string |  | Yes |
| tgl_parsed | string |  | Yes |
| timestamp | string |  | Yes |
