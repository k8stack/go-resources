#### installing protobuf on ubuntu

download protobuf-all from https://github.com/protocolbuffers/protobuf/releases

```
./configure
make
make check
sudo make install
sudo ldconfig # refresh shared library cache.
```

#### generating proto files in python

```
protoc -I src/user --python_out=src/user proto/*.proto
```

#### generating proto files in go

unlike python, go expects to generate all file in one package, so this 
has to be done one by one

```
protoc -I src/user/ --go_out=src/user/ src/user/eyecolor.proto 
protoc -I src/user/ --go_out=src/user/ src/user/address.proto 
protoc -I src/user/ --go_out=src/user/ src/user/user.proto 
```

#### generating proto using GOFAST plugin

```
protoc -I src/user/ --gofast_out=src/user/ src/user/user.proto
```

to get this working you need to install gofast `go get github.com/gogo/protobuf/protoc-gen-gofast`
referece: https://github.com/gogo/protobuf/blob/master/Readme.md
