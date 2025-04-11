#!/usr/bin/env bash

# https://www.inspur.com/eportal/fileDir/active_download/platformBookZh/6747/%E6%B5%AA%E6%BD%AE%E8%8B%B1%E4%BF%A1%E6%9C%8D%E5%8A%A1%E5%99%A8%20Redfish%E7%94%A8%E6%88%B7%E6%89%8B%E5%86%8C%20V1.1.pdf

# https://github.com/Mongo-Hao/redfish

# 11.58.56.1

# 获取临时Token
curl -k -i -X POST "https://11.59.23.1/redfish/v1/SessionService/Sessions" \
  -H "Content-Type: application/json" \
  -d '{"UserName": "root", "Password": "Abcd001002~!"}'

# 挂载NFS

curl -k -X POST https://11.52.15.1/redfish/v1/Managers -u root:abcd001002
