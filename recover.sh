ssh root@r41 "rm -rf /opt/starrocks && rm -rf /data/starrocks && cp -r /opt/starrocks.bak /opt/starrocks && cp -r /data/starrocks.bak /data/starrocks && killall -9 starrocks_be" &
ssh root@r42 "rm -rf /opt/starrocks && rm -rf /data/starrocks && cp -r /opt/starrocks.bak /opt/starrocks && cp -r /data/starrocks.bak /data/starrocks && killall -9 starrocks_be" &
ssh root@r43 "rm -rf /opt/starrocks && rm -rf /data/starrocks && cp -r /opt/starrocks.bak /opt/starrocks && cp -r /data/starrocks.bak /data/starrocks && killall -9 starrocks_be"
