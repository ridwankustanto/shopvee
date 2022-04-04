# Shopvee
## _Simple E-Commerce With Microservice_


## Description
Shopvee is simple e-commerce that implemented microservice architecture. It used gRPC as main communication between services and serve graphql for the client to consume it's function.

## Keypoint

- Mutate & Query services over GraphQL
- Use gRPC to communicate between services

## Architecture Design

<!--Upload Image-->

## System Design

<!--Upload Image-->

## Installation

To run Shopvee, you need to clone this repo and run this command

```sh
docker-compose up -d
```

Make sure you already start the Docker engine.

## Example

After containers up, open browser and go to http://localhost:8080

##### Mutation

- Create Account


    mutation {
        john: createAccount(account: {name: "John"}) {
            id
            name
        }
    }

- Create Product


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
    
    
- Place Order


    mutation {
        createOrder(order: {accountId: "account_id", products: [
            {id: "product_id", quantity: 1},
            {id: "product_id", quantity: 2},
            {id: "product_id", quantity: 3},
            ]}) {
                id
                totalPrice
                timestamp
        }
    }
    
    
    
##### Query

- Get Accounts


    // Get single item
    query {
      accounts(id: "account_id") {
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
    
    
- Get Products


    // Get single item
    query {
      products(id: "product_id") {
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
    
- Get Order By Account ID


    query {
        orders(account_id: "account_id") {
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