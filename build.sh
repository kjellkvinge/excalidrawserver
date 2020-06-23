#!/bin/bash

# clone, patch and build frontend
git clone https://github.com/excalidraw/excalidraw.git
cd excalidraw
git apply ../mypatch.patch
npm install
npm run build:app:docker
cd ..

# build binary
go run generator.go 
go build

#for windows
GOOS=windows GOARCH=amd64 go build -o static.exe

#for mac
GOOS=darwin GOARCH=amd64 go build -o static_darwin

#for arm
env GOOS=linux GOARCH=arm GOARM=5 go build static_arm
