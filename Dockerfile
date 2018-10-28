# docker build -t creditcard .
# docker run -it --name creditcard -v "$(pwd)"/src/:/root/src/ creditcard
# docker exec -it creditcard bash
FROM golang:1.8

WORKDIR /root/src
