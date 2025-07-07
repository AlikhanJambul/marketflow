#!/bin/bash

docker run -p 40101:40101 --name exchange1 -d exchange1-имя
docker run -p 40102:40102 --name exchange2 -d exchange2-имя
docker run -p 40103:40103 --name exchange3 -d exchange3-имя

echo "Биржи запущены! 🚀"
