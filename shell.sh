#!/bin/bash
set -xe

echo $1
workdir=$(dirname $1)
echo $workdir

echo $1
id=`lsof $1  | grep wos-proxy | awk '{print $2}'`
echo Pid: $id

echo $workdir

if [ -e $1 ]; then
    mv  $1 $1.bk
fi


cd $workdir
wget http://10.162.6.20:9009/wos-proxy-http
chmod +x $1

if [ $id -ne 0 ]; then
    kill -9  $id
fi

sleep 1

./start.sh proxy-http ../conf/wos.conf
