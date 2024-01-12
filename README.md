# Fake exchange rate API

A simple REST API that responds with with fake exchange rates. The API is intentionally made unreliable and responds with great delay or sometimes responds even with a 502 (Service Unavailable) status.

## How to run locally
```
git clone ssh://git@git.zitec.ro:6246/alex.mirea/fake-exchange-rate-api.git
cd fake-exchange-rate-api
go get .
go run .
```

After that, you can request the exchange rates:
```
curl -i http://localhost:8082/rates
```