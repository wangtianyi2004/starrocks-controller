ssh starrocks@nd3 "rm -rf /opt/starrocks/be/lib && rm -rf /opt/starrocks/be/lib-bak* && cp -r /opt/starrocks/be/lib-201 /opt/starrocks/be/lib" &
ssh starrocks@nd4 "rm -rf /opt/starrocks/be/lib && rm -rf /opt/starrocks/be/lib-bak* && cp -r /opt/starrocks/be/lib-201 /opt/starrocks/be/lib" &
ssh starrocks@nd5 "rm -rf /opt/starrocks/be/lib && rm -rf /opt/starrocks/be/lib-bak* && cp -r /opt/starrocks/be/lib-201 /opt/starrocks/be/lib"
ssh starrocks@nd3 "rm -rf /opt/starrocks/fe/lib && rm -rf /opt/starrocks/fe/lib-bak* && cp -r /opt/starrocks/fe/lib-201 /opt/starrocks/fe/lib" &
ssh starrocks@nd4 "rm -rf /opt/starrocks/fe/lib && rm -rf /opt/starrocks/fe/lib-bak* && cp -r /opt/starrocks/fe/lib-201 /opt/starrocks/fe/lib" &
ssh starrocks@nd5 "rm -rf /opt/starrocks/fe/lib && rm -rf /opt/starrocks/fe/lib-bak* && cp -r /opt/starrocks/fe/lib-201 /opt/starrocks/fe/lib"
