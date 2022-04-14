ssh starrocks@nd3 "rm -rf /opt/starrocks/* && rm -rf /data/starrocks/* && killall -9 starrocks_be"
ssh starrocks@nd4 "rm -rf /opt/starrocks/* && rm -rf /data/starrocks/* && killall -9 starrocks_be"
ssh starrocks@nd5 "rm -rf /opt/starrocks/* && rm -rf /data/starrocks/* && killall -9 starrocks_be"
