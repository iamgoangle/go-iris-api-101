#!/bin/sh

echo "Glide Create"
glide create

echo "Get vendor"
glide get github.com/kataras/iris

echo "Install vendor packages"
glide install

echo "Build"
go build

echo "Update vendor dependency"
glide up