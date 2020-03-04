#!/usr/bin/env sh




#public_ip=$(curl -s http://ifconfig.co/ip)
#
echo "ip地址: "$1



cpu_core=$(cat /proc/cpuinfo | grep "cpu cores" | head -1 | awk '{print $4}')


echo "cpu核数: "$cpu_core

memory_size=$(free -m | grep Mem | awk '{print $2}')
echo -n "内存: "
echo $memory_size | awk '{ byte=$1/1024; print byte "GB" }'




lvmsize=$(df -Th | grep ^[[:space:]] | awk '{sum += $2};END {print sum}')

if [ -z "$lvmsize" ]; then
    lvmsize=0
fi


standardsize=$(df -Th | grep ^/dev | awk '{sum += $3};END {print sum}')


disk_result=$(echo | awk "{print $lvmsize+$standardsize}")



echo "硬盘: "$disk_result"G"
