
# 有用的脚本

- 通过名称删除无用的镜像
    - `docker rmi -f  `docker images | grep '<none>' | awk '{print $3}'` `