# order-listing-challenge

This project allows for the creation and listing of orders, offering multiple services through gRPC, GraphQL, and a Web Server. The orders are stored in a database, with the final price being automatically calculated based on the provided price and tax.

## Features

- **Create Order:** You can create a new order by providing `id`, `Price`, and `Tax`. The `FinalPrice` is automatically calculated, and the order is saved in the database.
- **List Orders:** You can list all created orders, displaying the `id`, `Price`, `Tax`, and `FinalPrice`.

## Available Services

- **gRPC** - Port: `50051`
- **GraphQL** - Port: `8080`
- **Web Server** - Port: `8000`

## Usage

### Docker

1. To start all services and the database (with the orders table created automatically by migrations), run the following command:

   ```bash
   docker-compose up --build
### Web Server

- The REST API is available on port `8000`.
- To facilitate usage, there is an `api.http` file in the `api` folder containing ready-made requests to create and list orders.
- To use the `api.http` file, you need to install the **REST Client** extension in VSCode.

### GraphQL

- The GraphQL interface is available on port `8080`.
- To access it, open a browser and go to `http://localhost:8080`.

#### Example Mutation (Create Order)

```graphql
mutation createOrder {
  CreateOrder(input: { id: "a", Price: 1, Tax: 0.5 }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

#### Example Query (List Orders)

```graphql
query listOrders {
  ListOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

### gRPC

- To interact with the gRPC service, it is recommended to use the [Evans](https://github.com/ktr0731/evans) tool.
- After installing Evans, start a terminal and run the following command:

   ```bash
   evans -r repl
   ```

#### Evans Commands

1. Select the `pb` package:

   ```
   package pb
   ```

2. Select the `ListOrders` service:

   ```
   service ListOrders
   ```

3. To create an order, use the command:

   ```
   call CreateOrder
   ```

4. To list all created orders, use the command:

   ```
   call ListOrders
   ```

## Requirements

- Docker and Docker Compose installed
- **REST Client** extension for VSCode (optional, to use the `api.http` file)
- [Evans](https://github.com/ktr0731/evans) installed to interact with the gRPC service
