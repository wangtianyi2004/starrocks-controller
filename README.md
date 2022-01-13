# starrocks-controller

How to use it
1. yum install -y mysql
2. mkdir -p /root/.starrocks-controller/download
3. download jdk-8u301-linux-x64.tar.gz & StarRocks-2.0.0-GA.tar.gz 

--  In fact, the package can download these package automaticly. 

--  I upload these 2 package on my cloud server, but it too slow to download it. So I disable this link.

--  So pls download them from official website. 

--  [root@r22 download]# md5sum jdk-8u301-linux-x64.tar.gz

--  e77f9ea4c9ad849960ca295684ff9143  jdk-8u301-linux-x64.tar.gz

--  [root@r22 download]# md5sum StarRocks-2.0.0-GA.tar.gz

--  88a501c25e25a533759d6188cf12c2db  StarRocks-2.0.0-GA.tar.gz

4. run starrocks-controller as bellowing
[root@r22 ~]# ./sr-controller-main
