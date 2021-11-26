#!/usr/bin/env sh
cd /tmp
git clone https://github.com/SpiderLabs/ModSecurity.git
cd ModSecurity
git checkout ${MODSEC_VERSION:-v3.0.4}
git submodule update --init
./build.sh
./configure
make
make install