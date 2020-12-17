#!/bin/bash


docker images | grep fabric | grep -v amd64-0.4.15 | xargs docker rmi -f
docker images | grep none  | awk '{print $3}' | xargs docker rmi -f 

set -ex
startTime=$(date +%s)
pwd=$PWD
fabricPath=/home/services/go/src/github.com/hyperledger/fabric

cd $fabricPath
rm -rf release/linux-amd64/bin/ .build 
make release
mkdir -p .build/docker/gotools/
tar zxvf ~/go/bin.tar.gz -C .build/docker/gotools/
mkdir .build/bin
mv .build/docker/gotools/bin/chaintool .build/bin/
flag=1
while [[ $flag -ne 0 ]]; do
  set +e
  make docker
  flag=$?
  set -e
done

cd $pwd

docker tag hyperledger/fabric-baseos:amd64-0.4.15 swr.cn-east-3.myhuaweicloud.com/ebidsun/fabric-baseos:amd64-0.4.15
docker tag hyperledger/fabric-baseimage:amd64-0.4.15 swr.cn-east-3.myhuaweicloud.com/ebidsun/fabric-baseimage:amd64-0.4.15
docker tag hyperledger/fabric-tools:amd64-latest swr.cn-east-3.myhuaweicloud.com/ebidsun/fabric-tools:debug
docker tag hyperledger/fabric-orderer:amd64-latest swr.cn-east-3.myhuaweicloud.com/ebidsun/fabric-orderer:debug
docker tag hyperledger/fabric-peer:amd64-latest swr.cn-east-3.myhuaweicloud.com/ebidsun/fabric-peer:debug
docker tag hyperledger/fabric-buildenv:amd64-latest swr.cn-east-3.myhuaweicloud.com/ebidsun/fabric-buildenv:debug
docker tag hyperledger/fabric-ccenv:amd64-latest swr.cn-east-3.myhuaweicloud.com/ebidsun/fabric-ccenv:debug

delta=$(($(date +%s) - $startTime))
echo "--------------------------------------------------------------fabirc docker images build success, timeUsage=${delta} seconds-----------------------------------"


