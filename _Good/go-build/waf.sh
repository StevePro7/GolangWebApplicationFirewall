#!/usr/bin/env sh
echo "bootstrap start"

cd /tmp
git clone https://github.com/SpiderLabs/ModSecurity.git
cd ModSecurity
git submodule init
git submodule update
./build.sh
./configure
make
make install

echo "bootstrap end!!"
