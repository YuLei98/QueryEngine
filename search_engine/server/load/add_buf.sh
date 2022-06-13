#!/bin/bash
mkdir /swap_for_biz
cd /swap_for_biz/
dd if=/dev/zero of=swap_memory bs=4k count=2M
ls -lh
free -m
chmod 600 swap_memor
mkswap /swap_for_biz/swap_memory
swapon /swap_for_biz/swap_memory