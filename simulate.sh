#!/bin/bash

# just to test nodes discovery system

mkdir sim
go build main.go

for i in {5000..5003}
do
  mkdir sim/p$i
  sleep 2
  exec ./main -bind 127.0.0.1 -port $i -data sim/p$i -seed http://127.0.0.1:1312 &

  mypid=$!
  bash -c "sleep 30 && kill -9 $mypid" &
done