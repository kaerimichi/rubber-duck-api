# Rubber Duck API

This project is a simple REST API built using the [Echo](https://echo.labstack.com/) framework in Go. It provides CRUD operations to manage a collection of rubber ducks. The API is designed as an example for a presentation about Terraform providers.

## Features

- **List Rubber Ducks**: Retrieve a list of all rubber ducks.
- **Create Rubber Duck**: Add a new rubber duck with attributes like color, material, and size.
- **Update Rubber Duck**: Modify the details of an existing rubber duck.
- **Delete Rubber Duck**: Remove a rubber duck from the collection.

## API Endpoints

### 1. List Rubber Ducks
**GET** `/rubberducks`

- **Response**: Returns a JSON array of all rubber ducks.

### 2. Create Rubber Duck
**POST** `/rubberducks`

- **Request Body**:
  ```json
  {
    "color": "yellow",
    "material": "plastic",
    "size": "medium"
  }
