#!/bin/bash

# Forward local port to a port on the Pod
# Use 127.0.0.1:FIRST-PORT-BEFORE-COLON to connect
kubectl -n kafka      port-forward service/kafka-headless           29092:29092 2>&1 > /tmp/kafka.log &
kubectl -n minio      port-forward service/minio                    50000:9000  2>&1 > /tmp/minio.log &
kubectl -n clickhouse port-forward service/clickhouse-journal       8123:8123   2>&1 > /tmp/clickhouse-http.log &
kubectl -n clickhouse port-forward service/clickhouse-journal       9000:9000   2>&1 > /tmp/clickhouse-bin.log &
kubectl -n mysql      port-forward service/minimal-cluster-haproxy  3306:3306   2>&1 > /tmp/mysql.log &
