#! /usr/bin/env sh

wget https://github.com/antirez/redis/archive/unstable.tar.gz
tar -xzvf unstable.tar.gz
cd redis-unstable
make
./src/redis-server &
