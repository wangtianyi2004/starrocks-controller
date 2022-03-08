ssh root@r31 "rm -rf /opt/starrocks && rm -rf /data/starrocks && cp -r /opt/starrocks.bak /opt/starrocks && cp -r /data/starrocks.bak /data/starrocks && killall -9 starrocks_be" &
ssh root@r32 "rm -rf /opt/starrocks && rm -rf /data/starrocks && cp -r /opt/starrocks.bak /opt/starrocks && cp -r /data/starrocks.bak /data/starrocks && killall -9 starrocks_be" &
ssh root@r33 "rm -rf /opt/starrocks && rm -rf /data/starrocks && cp -r /opt/starrocks.bak /opt/starrocks && cp -r /data/starrocks.bak /data/starrocks && killall -9 starrocks_be"
