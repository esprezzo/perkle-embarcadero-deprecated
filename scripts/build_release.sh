#!/bin/bash

rm -rf ~/_build_tmp;
mkdir ~/_build_tmp;

sed -i '\|PATH="$PATH:/usr/lib/go-1.10/bin"|d' ~/.profile;
echo 'PATH="$PATH:/usr/lib/go-1.10/bin"' >> ~/.profile;
source ~/.profile

cd ~/_build_tmp && \
git clone https://github.com/alanwilhelm/go-ethereum.git \
&& mv go-ethereum esprezzo-chain && \
cd ./esprezzo-chain && \
git checkout ezp && make clean && \
make all;

rm -rf ~/esprezzo-chain;
mv ~/_build_tmp/esprezzo-chain/ ~/;
rm -rf ~/_build_tmp;