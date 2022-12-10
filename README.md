# Simple HTTP Auth

### Configuration
1. Set your _SIGNING_KEY_ in config file `./config/config.yml`
2. Change the ports of the apps if you need it

### Run project
```
git clone https://github.com/evilbebra/auth
cd auth
docker-compose up
```

## Auth

### GET localhost:8001/

GET jwt token

#### Example Response:
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV2aWxiZWJyYUBnbWFpbC5jb20iLCJpZCI6MSwidmFsaWRUaWxsIjoxNjk4MDA2ODg5fQ.L9VhKeBc4OxNq-t8i2V24WPb-b7J0_ARJKq2NjPrRuQ
```

## Api

### GET localhost:8000/hello

GET without jwt token

#### Example Response:
```
401 Unauthorized Error
```

### GET localhost:8000/hello

GET with jwt token 

#### Example Header:
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV2aWxiZWJyYUBnbWFpbC5jb20iLCJpZCI6MSwidmFsaWRUaWxsIjoxNjk4MDA3NDQ4fQ.jtiUC0KL3WdlxKEfXzhfCqAUt1I3DoliaIuQPw_a7QM
```

#### Example Response:
```
Hello, you are successfully authenticated!
```
