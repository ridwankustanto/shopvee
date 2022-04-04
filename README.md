# Shopvee
## _Simple E-Commerce With Microservice_


## Description
Shopvee is simple e-commerce that implemented microservice architecture. It used gRPC as main communication between services and serve graphql for the client to consume its function.

## Keypoint

- Mutate & Query services over GraphQL
- Microservices
- Use gRPC to communicate between services

## Architecture Design

![image](https://user-images.githubusercontent.com/19152005/161570629-8e2ac565-0205-4d78-93a0-658885f0f76c.png)


## System Design

![image](https://user-images.githubusercontent.com/19152005/161570681-25213180-23c1-4745-bd0a-a850d26fb35c.png)

## Installation

To run Shopvee, you need to clone this repo and run this command

```sh
docker-compose up -d
```

Make sure you already start the Docker engine.

## Testing

Since the client depends on their service and DB, so the test will run while running `docker-compose up --build -d unit_test` as a container. It will die after the test finished.


## Example

After containers up, open browser and go to http://localhost:8080

##### Mutation

- Create Account

Request

```
mutation {
    john: createAccount(account: {name: "John"}) {
        id
        name
    }
}
```
Response

```
{
  "data": {
    "createAccount": {
      "id": "281cf5f47d0942f199c82f0615cc2650",
      "name": "John"
    }
  }
}
```
- Create Product

Request

```
mutation {
    a : createProduct(product: {name: "Risol", description: "Risol pedas", price: 2500}) {
        id
        name
        description
        price
        timestamp
    }
    b : createProduct(product: {name: "Tahu Isi", description: "Tahu isi mercon", price: 2000}) {
        id
        name
        description
        price
        timestamp
    }
    c : createProduct(product: {name: "Tempe Kriuk", description: "Tempe kriuk banget", price: 1500}) {
        id
        name
        description
        price
        timestamp
    }
}
```    

Response

```
{
  "data": {
    "a": {
      "id": "1177101d45d7433ba05490a7f0158a25",
      "name": "Risol",
      "description": "Risol pedas",
      "price": 2500,
      "timestamp": "2022-04-04 17:51:28+0000"
    },
    "b": {
      "id": "ef09e0bb095f46a98213863f3acd355a",
      "name": "Tahu Isi",
      "description": "Tahu isi mercon",
      "price": 2000,
      "timestamp": "2022-04-04 17:51:28+0000"
    },
    "c": {
      "id": "4db01d4f07114328a5ffb9a0066059dc",
      "name": "Tempe Kriuk",
      "description": "Tempe kriuk banget",
      "price": 1500,
      "timestamp": "2022-04-04 17:51:28+0000"
    }
  }
}

```    
- Place Order

Request

```
mutation {
    createOrder(order: {accountId: "281cf5f47d0942f199c82f0615cc2650", products: [
        {id: "1177101d45d7433ba05490a7f0158a25", quantity: 1},
        {id: "ef09e0bb095f46a98213863f3acd355a", quantity: 2},
        {id: "4db01d4f07114328a5ffb9a0066059dc", quantity: 3},
        ]}) {
            id
            totalPrice
            timestamp
    }
}
``` 

Response

```
{
  "data": {
    "createOrder": {
      "id": "7d46e528105d4056a8b7ccf41599a0f4",
      "totalPrice": 11000,
      "timestamp": "2022-04-04 17:53:51+0000"
    }
  }
}
``` 
    
##### Query

- Get Accounts

Request

```
// Get single item
query {
  accounts(id: "281cf5f47d0942f199c82f0615cc2650") {
    id
    name
  }
}

// OR get list
query {
  accounts(pagination: {skip: 0, take: 10}) {
    id
    name
  }
}
```

Response
```
// Single item
{
  "data": {
    "accounts": [
      {
        "id": "281cf5f47d0942f199c82f0615cc2650",
        "name": "John"
      }
    ]
  }
}

// List item
{
  "data": {
    "accounts": [
      {
        "id": "ed20d2b8b68c4ba190b802122bee0082",
        "name": "Jonny"
      },
      {
        "id": "58c26beea9954ad782580ccf7fcabdbb",
        "name": "Jonny"
      },
      {
        "id": "341abf9f889048c682bdc7b6cf946426",
        "name": "Jonny"
      },
      {
        "id": "3287884891774b1788a705dc7bd587fe",
        "name": "Jonny"
      },
      {
        "id": "2f22955e92df4fdcb6f148500b50bc1e",
        "name": "Jonny"
      },
      {
        "id": "281cf5f47d0942f199c82f0615cc2650",
        "name": "John"
      }
    ]
  }
}
```
    
- Get Products

Request

```
// Get single item
query {
  products(id: "1177101d45d7433ba05490a7f0158a25") {
    id
    name
    description
    price
    timestamp
  }
}
    
// OR get list
query {
  products(pagination: {skip: 0, take: 10}) {
    id
    name
    description
    price
    timestamp
  }
}
```

Response
```
// Single item
{
  "data": {
    "products": [
      {
        "id": "1177101d45d7433ba05490a7f0158a25",
        "name": "Risol",
        "description": "Risol pedas",
        "price": 2500,
        "timestamp": "2022-04-04T17:51:28Z"
      }
    ]
  }
}
// List item
{
  "data": {
    "products": [
      {
        "id": "ef09e0bb095f46a98213863f3acd355a",
        "name": "Tahu Isi",
        "description": "Tahu isi mercon",
        "price": 2000,
        "timestamp": "2022-04-04T17:51:28Z"
      },
      {
        "id": "5745026f33744415a32f9d53320d467a",
        "name": "Pisang Goreng",
        "description": "Pisang goreng enak",
        "price": 2000,
        "timestamp": "2022-04-04T17:41:36Z"
      },
      {
        "id": "4db01d4f07114328a5ffb9a0066059dc",
        "name": "Tempe Kriuk",
        "description": "Tempe kriuk banget",
        "price": 1500,
        "timestamp": "2022-04-04T17:51:28Z"
      },
    ]
  }
}
```

- Get Order By Account ID

Request

```
query {
    orders(account_id: "281cf5f47d0942f199c82f0615cc2650") {
        id
        totalPrice
        timestamp
        products {
          id
          name
          description
          price
          quantity
        }
    }
}
```

Response
```
{
  "data": {
    "orders": [
      {
        "id": "7d46e528105d4056a8b7ccf41599a0f4",
        "totalPrice": 11000,
        "timestamp": "2022-04-04T17:53:51Z",
        "products": [
          {
            "id": "1177101d45d7433ba05490a7f0158a25",
            "name": "Risol",
            "description": "Risol pedas",
            "price": 2500,
            "quantity": 1
          },
          {
            "id": "4db01d4f07114328a5ffb9a0066059dc",
            "name": "Tempe Kriuk",
            "description": "Tempe kriuk banget",
            "price": 1500,
            "quantity": 3
          },
          {
            "id": "ef09e0bb095f46a98213863f3acd355a",
            "name": "Tahu Isi",
            "description": "Tahu isi mercon",
            "price": 2000,
            "quantity": 2
          }
        ]
      }
    ]
  }
}
```
