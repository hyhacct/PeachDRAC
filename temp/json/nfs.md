---
查看NFS
---

## 查看NFS状态

```bash
curl -k -u root:abcd001002 -X GET \
  "https://11.15.14.1/redfish/v1/Managers/iDRAC.Embedded.1/VirtualMedia/CD"
```

```json
{"@odata.context":"/redfish/v1/$metadata#VirtualMedia.VirtualMedia","@odata.id":"/redfish/v1/Managers/iDRAC.Embedded.1/VirtualMedia/CD","@odata.type":"#VirtualMedia.v1_2_0.VirtualMedia","Actions":{"#VirtualMedia.EjectMedia":{"target":"/redfish/v1/Managers/iDRAC.Embedded.1/VirtualMedia/CD/Actions/VirtualMedia.EjectMedia"},"#VirtualMedia.InsertMedia":{"target":"/redfish/v1/Managers/iDRAC.Embedded.1/VirtualMedia/CD/Actions/VirtualMedia.InsertMedia"}},"ConnectedVia":"NotConnected","Description":"iDRAC Virtual Media Services Settings","Id":"CD","Image":"http://<address>/centos_7_9_2009.iso","ImageName":"centos_7_9_2009.iso","Inserted":false,"MediaTypes":["CD","DVD"],"MediaTypes@odata.count":2,"Name":"Virtual CD","WriteProtected":true}
```

## 取消NFS挂载

```bash
curl -k -u root:abcd001002 -X POST \
  "https://11.15.14.1/redfish/v1/Managers/iDRAC.Embedded.1/VirtualMedia/CD/Actions/VirtualMedia.EjectMedia" \
  -H "Content-Type: application/json" \
  -d '{}'
```

## 挂载NFS

```bash
curl -k -u root:abcd001002 -X POST \
  "https://11.15.14.1/redfish/v1/Managers/iDRAC.Embedded.1/VirtualMedia/CD/Actions/VirtualMedia.InsertMedia" \
  -H "Content-Type: application/json" \
  -d '{ "Image": "http://10.37.50.3/isos/ubuntu.iso", "Inserted": true }'
```
