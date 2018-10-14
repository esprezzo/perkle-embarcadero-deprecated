#!/bin/bash
ps aux | grep './perkle/build/bin/geth' > /dev/null
if [ $? -eq 0 ]; then
  echo "Process is running."
else
  echo "Process is not running."
fi
