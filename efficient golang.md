# Hammer out rules for efficient golang

## Prefer using go install for Dependencies

If you want to pre-install dependencies to speed up builds, you can use go install to install the dependencies separately. This is similar to what the -i flag used to do.

- Example:
  ```bash
  go install ./...
  go build -o bin/nginx-operator main.go
  ```

  t nginx-operator
  ```bash
  go install ./...
  GO111MODULE=on CGO_ENABLED=0 go build -o bin/nginx-operator main.go
  ```

