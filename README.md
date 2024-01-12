# Fake exchange rate API

A simple REST API that responds with fake exchange rates. The API is intentionally made unreliable and responds with great delay or sometimes responds even with a 503 (Service Unavailable) status.

## How to run locally
Clone the repository and run the following commands:
```
cd fake-exchange-rate-api
go get .
go run .
```

After that, you can request the exchange rates:
```
curl -i http://localhost:8082/rates
```
