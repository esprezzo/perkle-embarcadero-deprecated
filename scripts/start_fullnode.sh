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
    --nodekey boot.key \
    # --bootnodes "enode://2d158bff0060f2cae3d2d01268cac4e652ecd7caa67e020e1b42c5ecef9c6c3e87e19198d968d2a4faa1d6fd09c6688dde595cf887a1223170129fe9930bf9f5@40.121.105.44:30310" \
    # --bootnodes "enode://4d29c257069826099e1b37b2bb70c8c5ba4dc6a897bf03ae4e078be8cb3ab3a6b3d4d7425ba1ad189e5c652955f061bb5aed36324fbd94a6b4a331133f5ae6e3@52.136.225.99:30311" \
    # --bootnodes "enode://d5e4956e254c067893394a9b89924275c3696c2e254b8cc74e0694718f4b657bf64427e98776c1675fb40188ebe2426170ba758dd367f87e5a46fc8cdbc2d783@104.42.157.173:30311" \
    --verbosity 3 \
    # --mine \
    # --minerthreads 6 \
   
