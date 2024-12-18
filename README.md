# gRPC API in Go

This repository contains a simple implementation of a gRPC API using **Go**. It provides CRUD operations for a product catalog system, demonstrating how to build and consume gRPC services in Go. The API allows clients to interact with the product data, including pagination, creating, updating, retrieving, and deleting products.

## Features

- **CRUD Operations**: Create, Read, Update, and Delete products.
- **Pagination**: Implemented for the list of products, returning paginated data.
- **Go & gRPC**: A demonstration of how to build and use gRPC services in Go.
- **Protocol Buffers**: Uses `.proto` files to define service methods and messages.

## Requirements

- **Go 1.18+**
- **gRPC** and **Protocol Buffers** installed
- **MySQL** or any other database (for a real production setup, if needed)

## Installation

### Clone the Repository

```bash
git clone https://github.com/yourusername/grpc-api-go.git
cd grpc-api-go
```
