#!/bin/bash
cd ./lambda
go build .
cd ..
rm ./lambda/lambdaapi.zip
zip ./lambda/lambdaapi.zip ./lambda/lambdaapi
terraform destroy -auto-approve 
terraform apply -auto-approve