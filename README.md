# starrocks-controller
In currently, I just build the playground functionality. 

Next step, I will build setup/startup/stop functionality

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




-------------------------------------------------------------------
目前，我只开发了 playground 功能，接下来，我将开发 setup/startup/stop 功能

如何使用
1. yum install -y mysql
2. mkdir -p /root/.starrocks-controller/download
3. 从官网下载 jdk-8u301-linux-x64.tar.gz & StarRocks-2.0.0-GA.tar.gz 

-- 其实这里是可以自动下载的，可以查看一下 sr-controller/playground/prepare/preparePkg.go

-- 里面是有下载链接的，pkgUrl := "http://10.10.10.20:9000/starrocks/StarRocks-2.0.0-GA.tar.gz"

-- 这里本来是我的云服务器，但是太穷了，只有 1MB 的带宽，下载要两个多小时，所以我改成了内网 ip，关闭了自动下载的入口

-- 官网的链接里面都设置了下载过期

4. 直接运行 starrocks-controller-main
-- chmod 751 starrocks-controller-main && ./starrocks-controller-main


