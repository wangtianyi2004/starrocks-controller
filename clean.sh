ssh root@r41 "rm -rf /opt/starrocks/* && rm -rf /data/starrocks/ && killall -9 starrocks_be"
ssh root@r42 "rm -rf /opt/starrocks/* && rm -rf /data/starrocks/ && killall -9 starrocks_be"
ssh root@r43 "rm -rf /opt/starrocks/* && rm -rf /data/starrocks/ && killall -9 starrocks_be"
