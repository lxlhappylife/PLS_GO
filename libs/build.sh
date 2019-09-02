#!/bin/bash
target=$1

echo $target
go build -buildmode=c-shared -o ../python/${target}.so ${target}.go
#cp ${target}.so ../python/
#cp ${target}.h ../python/

