# laundry-app-api

# Final Project API Documentation

## Auth

### POST Login
- Endpoint: `http://localhost:8080/login`
- Body (raw JSON):
  ```json
  {
      "username": "owner",
      "password": "rahasia"
  }
  
## Register Customer

### POST
- **Endpoint:** `http://localhost:8080/register/customer`
- **Body (raw JSON):**
  ```json
  {
      "username": "customer",
      "password": "rahasia",
      "email": "customer@mail.cm",
      "name": "customer",
      "address": "123 Main St",
      "phone_number": "023223424"
  }

## Register Admin

### POST
- **Endpoint:** `http://localhost:8080/register/admin`
- **Body (raw JSON):**
  ```json
  {
      "username": "employee",
      "password": "rahasia",
      "email": "employee@mail.cm",
      "name": "Employee",
      "address": "123 Main St",
      "phone_number": "0887654212"
  }

## Durations

### GET All Durations
- **Endpoint:** `http://localhost:8080/durations`
- **Authorization:** Bearer Token
- **Token:** `<token>`

### POST Create Duration
- **Endpoint:** `http://localhost:8080/duration`
- **Authorization:** Bearer Token
- **Token:** `<token>`
- **Body (raw JSON):**
  ```json
  {
    "duration_name": "Instant",
    "duration_days": 1
  }
  
### PUT Update Duration
- **Endpoint:** `http://localhost:8080/duration/2`
- **Authorization:** Bearer Token
- **Token:** `<token>`
- **Body (raw JSON):**
  ```json
  {
    "duration_days": 2
  }
  
### DELETE Duration
- **Endpoint:** `http://localhost:8080/duration/1`
- **Authorization:** Bearer Token
- **Token:** `<token>`

## Laundry Services

### GET All Laundry Services
- **Endpoint:** `http://localhost:8080/services`
- **Authorization:** Bearer Token
- **Token:** `<token>`

### POST Create Service
- **Endpoint:** `http://localhost:8080/service`
- **Authorization:** Bearer Token
- **Token:** `<token>`
- **Body (raw JSON):**
  ```json
  {
    "service_name": "Dry Cleaning",
    "service_desc": "Professional dry cleaning service",
    "service_price": 5000
  }
  
### PUT Update Service
- **Endpoint:** `http://localhost:8080/service/1`
- **Authorization:** Bearer Token
- **Token:** `<token>`
- **Body (raw JSON):**
  ```json
  {
    "laundry_name": "wet Cleaning",
    "laundry_desc": "amature dry cleaning service",
    "laundry_price": 7000
  }
  
### DELETE Service
- **Endpoint:** `http://localhost:8080/service/2`
- **Authorization:** Bearer Token
- **Token:** `<token>`

## Orders

### GET All Orders
- **Endpoint:** `http://localhost:8080/orders`
- **Authorization:** Bearer Token
- **Token:** `<token>`

### GET Order By User
- **Endpoint:** `http://localhost:8080/orders/1`
- **Authorization:** Bearer Token
- **Token:** `<token>`

### POST Create Order
- **Endpoint:** `http://localhost:8080/orders`
- **Authorization:** Bearer Token
- **Token:** `<token>`
- **Body (raw JSON):**
  ```json
  {
    "user_id": 2,
    "quantity": 2,
    "service_id": 1,
    "duration_id": 3
  }
