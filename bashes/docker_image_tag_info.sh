#!/usr/bin/env bash
# 查看镜像有哪些tag


repo_url=https://registry.hub.docker.com/v1/repositories
image_name=$1

curl -s ${repo_url}/${image_name}/tags | json_reformat | grep name | awk '{print $2}' | sed -e 's/"//g'