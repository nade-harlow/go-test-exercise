# Web Service
The web service creates a virtual card using flutterwave and blockchain wallet using stellar blockchain

### Setup Locally
- Run `go mod tidy` on your terminal to download service dependencies
- Run `go run main.go` to start up app.

## API ENDPOINTS DOCUMENTATION

### Create Card (POST)
creates a card using flutterwave

Endpoint
```
127.0.0.1:8088/api/v1/card
```
Payload
```
{
    "first_name": "Example",
    "last_name": "User",
    "date_of_birth": "1996/12/30",
    "email": "userg@example.com",
    "phone": "07030000000",
    "currency": "USD",
    "amount":5,
    "title": "MR",
    "gender": "M",
}
```
Response
```
{
    "id": "199a344f-1dbe-4b00-ba4d-beb014345fae",
    "account_id": 2061620,
    "amount": "5.00",
    "currency": "USD",
    "card_pan": "5319938155020288",
    "masked_pan": "531993*******0288",
    "city": "San Francisco",
    "state": "CA",
    "address_1": "333 Fremont Street",
    "zip_code": "94105",
    "cvv": "905",
    "expiration": "2025-09",
    "card_type": "mastercard",
    "name_on_card": "Tobi Olusola",
    "created_at": "2022-09-21T16:54:53.3851427+00:00",
    "is_active": true,
}
```

### Create Wallet (POST)
create wallet using stellar blockchain

Endpoint
```
127.0.0.1:8088/api/v1/wallet
```
Response
```
{
    "address": "GBJZO2QHGF5PM7S75H3LUODWUQ6NLKE5KTQNF26XVKNODSY52DTMPLHQ",
    "password": "SCJUKDQYOBCT2E6FJ2PHRD6B44DGCWOVHTCAL7NN43EYUPQJVCV2VVND"
}
```