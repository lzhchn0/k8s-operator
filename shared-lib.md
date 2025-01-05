## **Verify Your Project Structure**

```bash
mkdir -p  /root/shared-lib-demo/pkg/mathops
```

- The directory structure:
  ```
  /root/shared-lib-demo/
  ├── go.mod
  ├── pkg/
  │   └── mathops/
  │       └── mathops.go
  └── main.go
  ```




## Create the `pkg/mathops/mathops.go` file with the following content:

```go
// pkg/mathops/mathops.go
package mathops

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}
```


## Create the `main.go` file in the root directory

```go
// main.go
package main

import (
    "fmt"
    "shared-lib-demo/pkg/mathops"
)

func main() {
    sum := mathops.Add(10, 5)
    diff := mathops.Subtract(10, 5)
    fmt.Println("Sum:", sum)
    fmt.Println("Difference:", diff)
}
```

```bash
cd /root/shared-lib-demo
go install -buildmode=shared std
```
## Otherwise Errors and Explaination:
- The error message:
```
cannot use packages shared-lib-demo/pkg/mathops and runtime/cgo from different roots and /usr/local/go/pkg/linux_amd64_dynlink
```
- The error occurs because Go cannot mix packages from different roots when building a shared library. To resolve this, either avoid using shared libraries or rebuild the standard library as a shared library. Shared libraries in Go are not commonly used, so carefully consider whether they are necessary for your use case.


## **Initialize the Go Module**
```bash
cd /root/shared-lib-demo
go mod init shared-lib-demo
```

     
## **Rebuild the Standard Library as a Shared Library**     
```bash
go install -x -buildmode=shared -linkshared ./pkg/mathops
```
---
     
## **Clean the Build Cache**:
   - Clear the Go build cache to ensure there are no corrupted artifacts:
     ```bash
     go clean -cache -modcache -i -r
     ```
