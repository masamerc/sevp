#!/bin/bash

set -e

# set up
mkdir -p /root/.aws
mkdir -p /root/.config

# Test case 1
cp ./tests/aws_config.test /root/.aws/config
cp ./tests/sevp_config.test.toml /root/.config/sevp.toml
go test -v -count=1 -run IntegrationCase1 ./...

# Test case 2
rm -rf /root/.aws/config
go test -v -count=1 -run IntegrationCase2 ./...

# Test case 3
rm -rf /root/.aws/config
rm -rf /root/.config/sevp.toml
go test -v -count=1 -run IntegrationCase3 ./...
