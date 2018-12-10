#!/bin/bash
echo "********************************************************"
echo "Running user-cli"
echo "********************************************************"

while ! `nc -z user-service1 8080`; do sleep 3; done
while ! `nc -z user-service2 8080`; do sleep 3; done

./user-cli
