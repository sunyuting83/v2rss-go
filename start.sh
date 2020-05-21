#!/bin/bash
echo "start..."
export PATH=/sbin:/bin:/usr/bin:/usr/local/bin
basepath=$(cd `dirname $0`; pwd)
pidpath="$basepath/v2rss.pid"
if [ -f $pidpath ]; then
  pid=`cat $pidpath`
  kill -HUP $pid
  sleep 1
  cd $basepath
  nohup ./v2rss -p 5500 > run.log 2 >run.log 2>&1 &
  echo $! > $pidpath
  echo "The process $! is running..."
else
  cd $basepath
  nohup ./v2rss -p 5500 > run.log 2 >run.log 2>&1 &
  echo $! > $pidpath
  echo "The process $! is running..."
fi