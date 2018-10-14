#!/bin/bash
kill $(ps aux | grep './perkle/build/bin/geth' | awk '{print $2}');
ps -aux | grep geth;
