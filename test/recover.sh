cp -r ~/.starrocks-controller/cluster/sr-c1.bak ~/.starrocks-controller/cluster/sr-c1

ssh starrocks@nd3 "rm -rf /opt/starrocks/* && rm -rf /data/starrocks/* && cp -r /opt/starrocks.bak/* /opt/starrocks/ && cp -r /data/starrocks.bak/* /data/starrocks/ && killall -9 starrocks_be" &
ssh starrocks@nd4 "rm -rf /opt/starrocks/* && rm -rf /data/starrocks/* && cp -r /opt/starrocks.bak/* /opt/starrocks/ && cp -r /data/starrocks.bak/* /data/starrocks/ && killall -9 starrocks_be" &
ssh starrocks@nd5 "rm -rf /opt/starrocks/* && rm -rf /data/starrocks/* && cp -r /opt/starrocks.bak/* /opt/starrocks/ && cp -r /data/starrocks.bak/* /data/starrocks/ && killall -9 starrocks_be"
