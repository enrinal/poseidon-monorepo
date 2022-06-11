# Poseidon - Fetch-SVC
Fetch-SVC API

## Version: 0.1

### /api/v1/claims

#### GET
##### Description

Auto generated using Swagger Inspector

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Example response |

### /api/v1/commodity/aggregated

#### GET
##### Description

Auto generated using Swagger Inspector

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Example response |

### /api/v1/commodity

#### GET
##### Description

Auto generated using Swagger Inspector

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Example response |

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
