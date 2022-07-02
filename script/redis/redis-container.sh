# set -x

# 拉取镜像
docker pull redis

basePath="/home/xxx/redis/"
dataDir=$basePath"data"
confPath=$basePath"conf"
imageVersion="redis"

# echo $dataDir
# echo $confPath
# echo $imageVersion

# 启动容器
docker run \
 --name redis \
 -p 6379:6379 \
 -v $confPath:/etc/redis \
 -v $dataDir:/data \
 -d $imageVersion \
 redis-server /etc/redis/redis.conf
