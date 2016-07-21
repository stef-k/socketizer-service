#!/usr/bin/env bash
# deploy to production server
server_ip="85.159.213.249"

mkdir -p  "/tmp/socketizer-service/static/js/service/wordpress"
mkdir "/tmp/socketizer-service/logs"

go clean

go build && \
cp "static/js/service/wordpress/socketizer.min.js" "/tmp/socketizer-service/static/js/service/wordpress" && \
cp "socketizer-service" "/tmp/socketizer-service" && \
cp "socketizer-service.conf" "/tmp/socketizer-service" && \
tar -C /tmp/ -czf socketizer-service.tar.gz socketizer-service && \
scp socketizer-service.tar.gz stef@"$server_ip":/home/stef

rm -r "/tmp/socketizer-service"

rm socketizer-service.tar.gz

