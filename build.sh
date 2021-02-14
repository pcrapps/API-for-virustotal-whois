#!/bin/bash
cd ./lambda
go build .
rm  lambdaapi.zip
zip lambdaapi.zip lambdaapi
cd ..
terraform destroy -auto-approve 
terraform apply -auto-approve