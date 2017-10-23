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
ssh root@IP

apt update
apt install git

mkdir app
cd app

git clone https://github.com/hunterhug/AmazonBigSpider
git clone https://github.com/hunterhug/AmazonBigSpiderWeb

apt install docker.io
apt install docker-compose

cd AmazonBigSpider
cd sh/docker
chmod ./build.sh

./build

docker ps
docker exec -it GoSpider-redis redis-cli -a GoSpider
redis> keys *  (Ctrl+C)

docker exec -it GoSpider-mysqldb mysql -uroot -p459527502
mysql> show databases;
mysql> exit

```
