# jsonrpc remote cli - template

Project template that creates rpc server and cli client.

# How to

* ```make setup``` prepares the project
* Implement your logic in ```usecases``` - some demo examples are already there
* Add adaption for new usecases in ```usescases.go```
* review and adjust setting values in ```cmd/server/main.go```
* review and adjust setting value in ```cmd/generator/main.go```
* adjust ```BINARY_NAME``` in ```Makefile```
* ```make build-all``` to create server and client executable

# start server

execute the binary

```shell
rpc-cmd-server

# Output
2021/11/14 17:48:10 starting JSON-RPC Example (v0.1.0) on :8080
```

# use cli client

sample usage according to the demos provided

```shell
rpc-cmd add -a=40 -b=40

# Output
42
```

```shell
rpc-cmd split -data=a,b,c -delim=,

# Output
a
b
c
```

### help

each command supports **--help** or **-h**

```shell
rpc-cmd add --help

# Output
Usage of add:
  -a float
        -a=4.2
  -b float
        -b=4.2
```  

### client exit codes

| code  | description  |  
|---|---|
| 0 | ok
| 1 | transport error
| 2 | rpc error
| 3 | input error  