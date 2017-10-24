机器配置

```
root@iZwz9j36enzsqt8ab9tvkxZ:~# lsb_release -a
LSB Version:	core-9.20160110ubuntu0.2-amd64:core-9.20160110ubuntu0.2-noarch:security-9.20160110ubuntu0.2-amd64:security-9.20160110ubuntu0.2-noarch
Distributor ID:	Ubuntu
Description:	Ubuntu 16.04.2 LTS
Release:	16.04
Codename:	xenial

root@iZwz9j36enzsqt8ab9tvkxZ:~# free -h
              total        used        free      shared  buff/cache   available
Mem:           992M        280M         83M        3.1M        628M        534M
Swap:            0B          0B          0B
root@iZwz9j36enzsqt8ab9tvkxZ:~# df -h
Filesystem      Size  Used Avail Use% Mounted on
udev            479M     0  479M   0% /dev
tmpfs           100M  3.1M   97M   4% /run
/dev/vda1        40G  2.7G   35G   8% /
tmpfs           497M  124K  497M   1% /dev/shm
tmpfs           5.0M     0  5.0M   0% /run/lock
tmpfs           497M     0  497M   0% /sys/fs/cgroup
tmpfs           100M     0  100M   0% /run/user/0

```
以下是执行过程(全过程, 小白操作):

```
# 登录
ssh root@IP

# 更新安装git
apt update
apt install git

# 拉代码
mkdir -p ~/gocode/src/github.com/hunterhug
cd ~/gocode/src/github.com/hunterhug
git clone https://github.com/hunterhug/AmazonBigSpider
git clone https://github.com/hunterhug/AmazonBigSpiderWeb

# 安装必要软件
apt install docker.io
apt install docker-compose

# 启动MYSQL和Redis
cd AmazonBigSpider
cd sh/docker
chmod 777 ./build.sh
./build

#  检测是否安装成功
docker ps
docker exec -it GoSpider-redis redis-cli -a GoSpider
redis> keys *  (Ctrl+C)

docker exec -it GoSpider-mysqldb mysql -uroot -p459527502
mysql> show databases;
mysql> exit

# scp go1.8压缩包到远程机器
# scp xxxx.tar.gz  ssh@IP:
# 安装golang1.8
tar -zxvf xxxx.tar.gz
vim /etc/profile.d/myenv.sh

>>>>
export GOROOT=/root/go
export GOPATH=/root/gocode
export GOBIN=$GOPATH/bin
export PATH=.:$PATH:$GOROOT/bin:$GOBIN
:wq
>>>>

source /etc/profile.d/myenv.sh
go env

# 解决依赖
cd /root/gocode/src/github.com/hunterhug/AmazonBigSpider
go get -v github.com/tools/godep

# 编译爬虫端二进制, 并且初始化数据库(包括获取类目URL)
#

```
