todo-app-go
===
Sample of todo app written in Go.


# Usage

1. clone this repository

2. start server

3. request


# Development

## Generate server code

```bash
cd api \
&& oapi-codegen --package=api --generate types,chi-server,spec -o todoapp.gen.go  ../spec.yaml
```
