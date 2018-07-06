#!/bin/bash 

go test

go build -o food-api

echo "Your Food API is running..."
./food-api
