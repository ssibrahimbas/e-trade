<p align="center"><br><img src="https://avatars.githubusercontent.com/u/76786120?v=4" width="128" height="128" style="border-radius: 50px;" /></p>
<h3 align="center">Ssibrahimbas E-Trade (with Go)</h3>
<p align="center">
  An e-trade application with Go and MongoDB.
</p>

### API

<docgen-index>

| Module |
| ------ |
| **[Product](#Product)**|

</docgen-index>

<docgen-api>

## Product

Operations related to the product are listed here.

### Create Product

This endpoint should be used to create a product.

|Params|Type|Description|Required|
|------|----|-----------|--------|
|**name**|*string*|Product name|**true**|
|**price**|*float*|Product price|**true**|
|**stockCount**|*int*|Product stock count|**true**|

#### Simple Request

```http request
POST http://localhost:3000/v1/products/create
Content-Type: application/json

{
    "name": "Apple-pie",
    "price": 34.00,
    "stockCount": 12
}
```

#### Simple Response

```json
{
    "success": true,
    "message": "Product created successfully.",
    "data": "6280fcd46623395ac01e3f51"
}
```

### Get All Products

All products can be listed with this endpoint.

#### Simple Request

```http request
GET http://localhost:3000/v1/products
```

#### Simple Response

```json
{
    "success": true,
    "message": "Products successfully fetched.",
    "data": [
        {
            "id": "6280fcd46623395ac01e3f51",
            "name": "Apple-pie",
            "price": 34,
            "stockCount": 12,
            "date": "2022-05-15T13:15:00.33Z"
        }
    ]
}
```

### Get Product Detail

Product detail can be listed with this endpoint.

|Params|Type|Description|Required|
|------|----|-----------|--------|
|**id**|*ObjectId*|Product Id|**true**|

#### Simple Request

```http request
GET http://localhost:3000/v1/products/6280fcd46623395ac01e3f51
```

#### Simple Response

```json
{
    "success": true,
    "message": "Product successfully fetched.",
    "data": {
        "id": "6280fcd46623395ac01e3f51",
        "name": "Apple-pie",
        "price": 34,
        "stockCount": 12,
        "date": "2022-05-15T13:15:00.33Z"
    }
}
```


</docgen-api>