#!/bin/sh

set -e

nodes="$(hostname -i):7001 redis7002:7002 redis7003:7003 redis7004:7004 redis7005:7005 redis7006:7006"
(if [ "$1" == "create-cluster" ]; then
    for node in $nodes; do
        hostAndPort=$(echo $node | tr ":" " ")
        until [ $(nc -vz $hostAndPort) ]
        do
            echo "$hostAndPort open"
            break
        done
    done

    echo "yes" | eval /usr/local/bin/redis-cli -a "${REDIS_PASSWORD}" --cluster create $nodes --cluster-replicas 1
fi)&

redis-server /etc/redis/redis.conf
