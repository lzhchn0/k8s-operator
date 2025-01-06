## Member
- ianlancetaylor
- [cmd/go: remove -buildmode=shared (not c-shared)](https://github.com/golang/go/issues/47788)
- rasky

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
```
# build shared std lib
%if %{shared}
GOROOT=$(pwd) PATH=$(pwd)/bin:$PATH go install -buildmode=shared -v -x std
%endif
```

```bash
cd /root/shared-lib-demo
go install -v -x -buildmode=shared std
```
## Otherwise Errors and Explaination:
- The error message:
```
cannot use packages shared-lib-demo/pkg/mathops and runtime/cgo from different roots and /usr/local/go/pkg/linux_amd64_dynlink
```
- The error occurs because Go cannot mix packages from different roots when building a shared library. To resolve this, either avoid using shared libraries or rebuild the standard library as a shared library. Shared libraries in Go are not commonly used, so carefully consider whether they are necessary for your use case.

- makeslice: cap out of range
- This error typically occurs when attempting to create a slice with a length greater than its capacity

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
$ go env
```
GO111MODULE=''
GOARCH='amd64'
GOBIN=''
GOCACHE='/root/.cache/go-build'
GOENV='/root/.config/go/env'
GOEXE=''
GOEXPERIMENT=''
GOFLAGS=''
GOHOSTARCH='amd64'
GOHOSTOS='linux'
GOINSECURE=''
GOMODCACHE='/root/go/pkg/mod'
GONOPROXY=''
GONOSUMDB=''
GOOS='linux'
GOPATH='/root/go'
GOPRIVATE=''
GOPROXY='https://proxy.golang.org,direct'
GOROOT='/usr/local/go'
GOSUMDB='sum.golang.org'
GOTMPDIR=''
GOTOOLCHAIN='auto'
GOTOOLDIR='/usr/local/go/pkg/tool/linux_amd64'
GOVCS=''
GOVERSION='go1.22.3'
GCCGO='gccgo'
GOAMD64='v1'
AR='ar'
CC='gcc'
CXX='g++'
CGO_ENABLED='1'
**GOMOD='/root/shared-lib-demo/go.mod'**
GOWORK=''
CGO_CFLAGS='-O2 -g'
CGO_CPPFLAGS=''
CGO_CXXFLAGS='-O2 -g'
CGO_FFLAGS='-O2 -g'
CGO_LDFLAGS='-O2 -g'
PKG_CONFIG='pkg-config'
GOGCCFLAGS='-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build46567694=/tmp/go-build -gno-record-gcc-switches'

```

![alt text](./go_env.jpeg)

### go version go1.19.4 linux/amd64

```bash
go install -v -x -buildmode=shared std
```



```bash
go build -linkshared  main.go 
```

```
controlplane $ ls -ltr
-rwxr-xr-x 1 root root 19664 Jan  6 02:00 shared-lib-demo
-rwxr-xr-x 1 root root 19664 Jan  6 02:04 main
```

```
$ ldd  ./main
        linux-vdso.so.1 (0x00007ffcf9b9a000)
        libstd.so => /usr/local/go/pkg/linux_amd64_dynlink/libstd.so (0x00007fb83fd5b000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fb83fb59000)
        libdl.so.2 => /lib/x86_64-linux-gnu/libdl.so.2 (0x00007fb83fb53000)
        libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007fb83fb30000)
        /lib64/ld-linux-x86-64.so.2 (0x00007fb842955000)
```



```bash
export LD_LIBRARY_PATH=/path/to/your/shared/libraries:$LD_LIBRARY_PATH
```


main program works without `LD_LIBRARY_PATH` because the dynamic linker is finding the `.so` files through one of the following mechanisms:
- The `.so` files are located in a default system path (e.g., `/usr/local/lib`).
- The executable has an `RPATH` or `RUNPATH` embedded in it.
- The `.so` files are registered in the system cache (`/etc/ld.so.cache`).

Use tools like `ldd`, `readelf`, and `strace` to debug and understand how the linker is locating your shared libraries.

## Workflow to Debug:
1. **Check Shared Library Dependencies**:
   ```bash
   ldd ./myprogram
   ```

2. **Check for RPATH or RUNPATH**:
   ```bash
   readelf -d ./myprogram | grep RPATH
   readelf -d ./myprogram | grep RUNPATH
   ```

3. **Check System Cache**:
   ```bash
   ldconfig -p | grep libmathops
   ```

4. **Trace Library Loading**:
   ```bash
   strace ./myprogram 2>&1 | grep openat
   ```


## Tools
### 1. **`ldd`**:
   - Run `ldd` on your executable to see which shared libraries it depends on and where they are located:
     ```bash
     ldd ./myprogram
     ```
   - Example output:
     ```
     linux-vdso.so.1 (0x00007ffd45df0000)
     libmathops.so => /usr/local/lib/libmathops.so (0x00007f8c12345000)
     libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f8c12123000)
     /lib64/ld-linux-x86-64.so.2 (0x00007f8c12367000)
     ```
   - In this example, `libmathops.so` is located in `/usr/local/lib`.

### 2. **`readelf`**:
   - Use `readelf` to check for `RPATH` or `RUNPATH` in your executable:
     ```bash
     readelf -d ./myprogram | grep RPATH
     readelf -d ./myprogram | grep RUNPATH
     ```

### 3. **`strace`**:
   - Use `strace` to trace the system calls made by your program and see where it is looking for shared libraries:
     ```bash
     strace ./myprogram 2>&1 | grep openat
     ```
   - Look for lines where the program opens `.so` files.
