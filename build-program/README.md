# Building and installing go programs

This folder contains basic demonstration of build and install process of go programs

## Steps

1. Check if program runs

```shell
go run
```

2. Build binary

```shell
go build
```

3. Install go binaries/packages

```shell
go install 
```

4. Running go built programms

* go should install package in ~/HOME/go/bin/ folder (even in windows)
* to run it - in shell type folder/package name (in this case build-program)

5. * Uninstalling program

In Go, go install builds a single-file binary and "installs" it by copying it to the appropriate directory (*). To "uninstall" this binary, simply remove it with rm. [https://stackoverflow.com/questions/66663173/is-there-a-go-uninstall#:~:text=1%20Answer&text=Removing%20the%20installed%20executable%20with,simply%20remove%20it%20with%20rm%20.]

```shell
rm
```

or tidy the unused binaries/modules

```shell
go mod tidy
```