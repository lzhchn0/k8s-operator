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
     go install -buildmode=shared -linkshared ./pkg/mathops
     ```

     
     ```bash
     go mod init shared-lib-demo
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
