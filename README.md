# Serverless with Golang 

## Prerequisites

- Go 1.1x
- Serverless Framework
- AWS SAM CLI
- Docker

## Database

I use postgres, Here is instalation postgres using docker

```bash
docker pull postgres

# alpine is lite version
docker run -d -p 0.0.0.0:5432:5432 postgres:alpine
```

## Installing

Installation just type 
```bash
make build
```

## Run in local with SAM (AWS-SAM-CLI)

Run with SAM
```bash
# http://localhost:8000/
sam local start-api -p 8000 --env-vars local.env.json
```

## Deployment

```bash
# set your serverless config
serverless config credentials --provider aws --key <ACCESS KEY ID> --secret <SECRET KEY>

# deploy
serverless deploy
```

## Get All Users

> live demo [https://0d8n48wogc.execute-api.us-east-1.amazonaws.com/dev/user](https://0d8n48wogc.execute-api.us-east-1.amazonaws.com/dev/user)

Request
```bash
curl --request GET --url http://localhost:8000/user | json_pp 
```

Response
```json
{
  "data": {
    "Users": [
      ...
      {
        "user_id": 3,
        "email": "user_3@gmail.com"
      },
      {
        "user_id": 4,
        "email": "user_4@gmail.com"
      },
      {
        "user_id": 5,
        "email": "user_5@yahoo.com"
      },
      ...
    ]
  }
}
```

## Get User By Id

> live demo [https://0d8n48wogc.execute-api.us-east-1.amazonaws.com/dev/user/2](https://0d8n48wogc.execute-api.us-east-1.amazonaws.com/dev/user/2)

Request
```bash
curl --request GET --url http://localhost:8000/user/4 | json_pp
```

Response
```json
{
  "data": {
    "User": {
      "user_id": 4,
      "email": "user_4@gmail.com"
    }
  }
}
```

### References
- Create serverless template with golang, https://serverless.com/blog/framework-example-golang-lambda-support/
- Deploy Serverless, https://medium.com/devopslinks/aws-lambda-serverless-framework-python-part-1-a-step-by-step-hello-world-4182202aba4a
- Run local aws-sam-cli, https://github.com/awslabs/aws-sam-cli
- Generating fake data using SQL, https://vnegrisolo.github.io/postgresql/generate-fake-data-using-sql

