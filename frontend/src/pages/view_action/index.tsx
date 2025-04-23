import { Card, Button, Space, TextArea, Col, Row, Typography, Tag, Descriptions, Spin } from '@douyinfe/semi-ui';


const { Text } = Typography;


const response = [
  {
    "ipmi": "192.168.1.1",
    "status": "success",
    "model": "Dell R740",
    "sn": "1234567890",
    "action": "Java - 启动控制台"
  },
  {
    "ipmi": "192.168.1.2",
    "status": "error",
    "model": "Dell R740",
    "sn": "1234567890",
    "action": "Java - 启动控制台"
  },
  {
    "ipmi": "192.168.1.1",
    "status": "success",
    "model": "Dell R740",
    "sn": "1234567890",
    "action": "Java - 启动控制台"
  },
  {
    "ipmi": "192.168.1.2",
    "status": "error",
    "model": "Dell R740",
    "sn": "1234567890",
    "action": "Java - 启动控制台"
  },
  {
    "ipmi": "192.168.1.1",
    "status": "success",
    "model": "Dell R740",
    "sn": "1234567890",
    "action": "Java - 启动控制台"
  },
  {
    "ipmi": "192.168.1.2",
    "status": "error",
    "model": "Dell R740",
    "sn": "1234567890",
    "action": "Java - 启动控制台"
  },
  {
    "ipmi": "192.168.1.2",
    "status": "error",
    "model": "Dell R740",
    "sn": "1234567890",
    "action": "Java - 启动控制台"
  },
]
function ViewAction() {
  return (
    <div>
      <div>
        <Row gutter={16}>
          <Col span={12}>
            <Card>
              <Space wrap={true}>
                <Button>Java - 启动控制台</Button>
                <Button type="warning">电源 - 开机</Button>
                <Button type="warning">电源 - 关机</Button>
                <Button type="warning">电源 - 强制关机</Button>
                <Button type="warning">电源 - 强制重启</Button>
                <Button type="tertiary">风扇 - 设置转速</Button>
                <Button type="tertiary">风扇 - 自动转速</Button>
                <Button type="danger">NFS - 挂载</Button>
                <Button type="danger">NFS - 卸载</Button>
              </Space>
            </Card>
          </Col>

          <Col span={12}>
            <Card style={{ maxWidth: "100%" }}>
              <TextArea maxCount={100} showClear placeholder="请在此输入IPMI, 每行一个即可" />
            </Card>
          </Col>
        </Row>
      </div>

      <div style={{ marginTop: "16px" }}>
        <Row gutter={[16, 16]}>
          {response.map((item, index) => (
            <Col span={6}>
              <div>
                <Spin size="large" >
                  <Card
                    title={item.ipmi}
                    style={{ maxWidth: 300 }}
                    headerExtraContent={
                      <Text link color="warning">
                        重试
                      </Text>
                    }
                  >
                    <Descriptions>
                      <Descriptions.Item itemKey="状态">
                        {item.status == "success" ? <Tag color="green">成功</Tag> : <Tag color="pink">失败</Tag>}
                      </Descriptions.Item>
                      <Descriptions.Item itemKey="型号">{item.model}</Descriptions.Item>
                      <Descriptions.Item itemKey="序列号">{item.sn}</Descriptions.Item>
                      <Descriptions.Item itemKey="操作">{item.action}</Descriptions.Item>
                    </Descriptions>
                  </Card>
                </Spin>
              </div>
            </Col>
          ))}
        </Row>
      </div>
    </div>
  )
}

export default ViewAction;