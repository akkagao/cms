#!/bin/sh

tarfile="cms-v1.0.tar.gz"

echo "开始打包$tarfile..."

export GOARCH=amd64
export GOOS=linux

bee pack -exs="pack.sh:pack:bat:nginx.conf" -exr=data

mv cms.tar.gz $tarfile
