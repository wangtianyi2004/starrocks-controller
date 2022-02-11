ssh root@nd3 "rm -rf /opt/starrocks && rm -rf /data/starrocks && cp -r /opt/starrocks.bak /opt/starrocks && cp -r /data/starrocks.bak /data/starrocks && killall -9 starrocks_be" &
ssh root@nd4 "rm -rf /opt/starrocks && rm -rf /data/starrocks && cp -r /opt/starrocks.bak /opt/starrocks && cp -r /data/starrocks.bak /data/starrocks && killall -9 starrocks_be" &
ssh root@nd5 "rm -rf /opt/starrocks && rm -rf /data/starrocks && cp -r /opt/starrocks.bak /opt/starrocks && cp -r /data/starrocks.bak /data/starrocks && killall -9 starrocks_be"
