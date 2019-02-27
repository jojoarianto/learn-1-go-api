# Serverless with Golang 

## Prerequisites

- Go 1.1x
- Serverless Framework
- AWS SAM CLI
- Docker

## Goal
- [X] API to [get all user](#get-all-users)
- [X] API to [get user by id](#get-user-by-id)
- [x] API to [create user](#create-user) 
- [ ] API to delete user
- [ ] API to update user

## Run in local with SAM (AWS-SAM-CLI)

Compile go project
```bash
# see Makefile
make build
```

Run with SAM, (required : docker, postgres)
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
curl --request GET \
  --url https://0d8n48wogc.execute-api.us-east-1.amazonaws.com/dev/users
```

Response
```json

  "data": [
    ...
    {
      "user_id": 6,
      "email": "user_6@gmail.com"
    },
    {
      "user_id": 7,
      "email": "user_7@hotmail.com"
    },
    ...
  ]
}
```

## Get User By Id

> live demo [https://0d8n48wogc.execute-api.us-east-1.amazonaws.com/dev/user/2](https://0d8n48wogc.execute-api.us-east-1.amazonaws.com/dev/user/2)

Request
```bash
curl --request GET \
  --url https://0d8n48wogc.execute-api.us-east-1.amazonaws.com/dev/user/14
```

Response
```json
{
  "data": {
    "user_id": 14,
    "email": "create-something-cool@gmail.com"
  }
}
```

## Create User
Request
```bash
curl --request POST \
  --url https://0d8n48wogc.execute-api.us-east-1.amazonaws.com/dev/user \
  --header 'content-type: application/json' \
  --data '{
	      "email": "anything@gmail.com"
    }'
```

### References
- Create serverless template with golang, https://serverless.com/blog/framework-example-golang-lambda-support/
- Deploy Serverless, https://medium.com/devopslinks/aws-lambda-serverless-framework-python-part-1-a-step-by-step-hello-world-4182202aba4a
- Run local aws-sam-cli, https://github.com/awslabs/aws-sam-cli
- Generating fake data using SQL, https://vnegrisolo.github.io/postgresql/generate-fake-data-using-sql

## Additional note

### Database

I use postgres, Here is instalation postgres using docker

```bash
docker pull postgres

# alpine is lite version
docker run -d -p 0.0.0.0:5432:5432 postgres:alpine
```
