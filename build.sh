#! /bin/bash

mkdir -p /osv-scanner/reports
cd osv-scanner/cmd/osv-scanner/
go build
mv ./osv-scanner /usr/local/bin
cd ../..
mv python-script /osv-scanner/

export PATH=$PATH:$(pwd)