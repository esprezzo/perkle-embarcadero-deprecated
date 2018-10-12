#!/bin/bash
kill $(ps aux | grep '[./esprezzo]-chain/build/bin/geth' | awk '{print $2}');
ps -aux | grep geth;
