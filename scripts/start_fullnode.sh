#!/bin/bash
~/perkle/build/bin/geth \
    --datadir ~/.perkle \
    --syncmode 'full' \
    --port 30311 \
    --rpc \
    --rpcaddr '0.0.0.0' \
    --rpcport 8501 \
    --rpcapi 'db,eth,net,web3,txpool,miner,personal' \
    # --gasprice '1' \
    --etherbase '0x8C12E659ACE965C4A17F3664135E6BB0292789E0' \
    --nodekey ~/boot.key \
    --verbosity 3 \
    --mine \
    --minerthreads 6 \
   
