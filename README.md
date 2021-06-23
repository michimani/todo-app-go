todo-app-go
===
Sample of todo app written in Go.


# Preparing

## clone this repository

```bash
git clone https://github.com/michimani/todo-app-go.git \
&& cd todo-app-go
```

## set up DynamoDB local

1. Download DynamoDB local from the link bellow.

    [Deploying DynamoDB Locally on Your Computer - Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.DownloadingAndRunning.html)

2. Start DynamoDB local.

    ```bash
    export DYNAMODB_LOCAL_DIR=your_dynamo_db_local_path \
    && sh ./scripts/start_local_db.sh
    ```

3. Create table.

    ```bash
    aws dynamodb create-table \
    --cli-input-json file://local/dynamodb_local_schema.json \
    --endpoint-url http://localhost:8001
    ```

## start server

TBD

## request

TBD

# Development

## Generate server code

```bash
cd api \
&& oapi-codegen --package=api --generate types,chi-server,spec -o todoapp.gen.go  ../spec.yaml
```
