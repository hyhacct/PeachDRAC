#!/usr/bin/env bash

# 获取临时Token
curl -k -i "http://11.28.22.1/redfish/v1/SessionService/Sessions" \
  -H "Content-Type: application/json" \
  -d '{"UserName": "root", "Password": "Abcd001002~!"}'

# 挂载NFS
curl -k -X POST https://11.28.22.1/redfish/v1/Managers -u root:abcd001002
