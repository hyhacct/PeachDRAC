import { Tag, Banner, Card, Space, Empty } from "@douyinfe/semi-ui";
import { Descriptions, Col, Row, Typography, OverflowList } from '@douyinfe/semi-ui';
import { IconRating } from "@douyinfe/semi-icons-lab";
import { IllustrationNoContent, IllustrationNoContentDark } from '@douyinfe/semi-illustrations';


function ViewJava() {
  const { Meta } = Card;
  const { Text } = Typography;

  const data = [
    {
      title: 'Java8 - 戴尔',
      path: "/usr/bin/java",
      version: "1.8.0_282",
      ipmi: ['11.98.11.1', '11.98.11.2', '11.98.11.3', '11.98.11.4', '11.98.11.5', '11.98.11.6', '11.98.11.7', '11.98.11.8', '11.98.11.9', '11.98.11.10'],
      model: "Dell R740",
      createAt: "2025-04-23 16:43:14"
    },
    {
      title: 'Java8 - 戴尔',
      path: "/usr/bin/java",
      version: "1.8.0_282",
      ipmi: [],
      model: "Dell R740",
      createAt: "2025-04-23 16:43:14"
    },
    {
      title: 'Java8 - 戴尔',
      path: "/usr/bin/java",
      version: "1.8.0_282",
      ipmi: [],
      model: "Dell R740",
      createAt: "2025-04-23 16:43:14"
    },
    {
      title: 'Java8 - 戴尔',
      path: "/usr/bin/java",
      version: "1.8.0_282",
      ipmi: [],
      model: "Dell R740",
      createAt: "2025-04-23 16:43:14"
    },
    {
      title: 'Java8 - 戴尔',
      path: "/usr/bin/java",
      version: "1.8.0_282",
      ipmi: [],
      model: "Dell R740",
      createAt: "2025-04-23 16:43:14"
    },
    {
      title: 'Java8 - 戴尔',
      path: "/usr/bin/java",
      version: "1.8.0_282",
      ipmi: [],
      model: "Dell R740",
      createAt: "2025-04-23 16:43:14"
    },
    {
      title: 'Java8 - 戴尔',
      path: "/usr/bin/java",
      version: "1.8.0_282",
      ipmi: [],
      model: "Dell R740",
      createAt: "2025-04-23 16:43:14"
    },
    {
      title: 'Java8 - 戴尔',
      path: "/usr/bin/java",
      version: "1.8.0_282",
      ipmi: [],
      model: "Dell R740",
      createAt: "2025-04-23 16:43:14"
    },
    {
      title: 'Java8 - 戴尔',
      path: "/usr/bin/java",
      version: "1.8.0_282",
      ipmi: [],
      model: "Dell R740",
      createAt: "2025-04-23 16:43:14"
    },
    {
      title: 'Java8 - 戴尔',
      path: "/usr/bin/java",
      version: "1.8.0_282",
      ipmi: [],
      model: "Dell R740",
      createAt: "2025-04-23 16:43:14"
    },
    {
      title: 'Java8 - 戴尔',
      path: "/usr/bin/java",
      version: "1.8.0_282",
      ipmi: [],
      model: "Dell R740",
      createAt: "2025-04-23 16:43:14"
    },
  ];


  const style = {
    border: '1px solid var(--semi-color-border)',
    backgroundColor: 'var(--semi-color-bg-2)',
    borderRadius: '3px',
    paddingLeft: '20px',
  };

  return (
    <div style={{ marginTop: '16px' }}>
      <Banner
        type="info"
        description="若需要为不同厂商或者IPMI指定启动控制台的Java版本, 请点击[新增配置]按钮, 若无配置则使用系统默认的Java版本"
      />
      <div style={{ marginTop: '16px' }}>
        <Row gutter={[10, 10]}>
          {data.map((item) => (
            <Col span={8}>
              <div >
                <Card
                  style={style}
                  bordered={false}
                  headerLine={true}
                  title={
                    <Meta
                      title={item.title}
                      description={item.createAt}
                      avatar={<IconRating size="extra-large" />}
                    />
                  }
                  headerExtraContent={
                    <div>
                      <Space wrap>
                        <Text link>
                          编辑
                        </Text>
                        <Text link type="danger">
                          删除
                        </Text>
                      </Space>
                    </div>
                  }
                >
                  <Descriptions>
                    <Descriptions.Item itemKey="Java版本">{item.version}</Descriptions.Item>
                    <Descriptions.Item itemKey="Java路径">{item.path}</Descriptions.Item>
                    <Descriptions.Item itemKey="应用型号">{item.model}</Descriptions.Item>
                    <Descriptions.Item itemKey="应用IPMI">
                      {item.ipmi.length != 0 ?
                        <div>
                          <Tag shape='circle' color='violet'>
                            {item.ipmi[0]}
                          </Tag>
                          <Tag shape='circle' style={{ marginLeft: '8px' }}>
                            {"+" + item.ipmi.length + "台"}
                          </Tag>
                        </div>
                        :
                        <Tag shape='circle' color='violet'>
                          暂无
                        </Tag>
                      }
                    </Descriptions.Item>
                  </Descriptions>
                </Card>
              </div>
            </Col>
          ))}

          <Col span={8}>
            <div >
              <Card
                style={style}
                bordered={false}
                headerLine={true}
                shadows='hover'
              >
                <Empty
                  image={<IllustrationNoContent style={{ width: 150, height: 150 }} />}
                  darkModeImage={<IllustrationNoContentDark style={{ width: 150, height: 133 }} />}
                  title={'新增配置'}
                  description="点击添加一个新的配置"
                />
              </Card>
            </div>
          </Col>
        </Row>
      </div>
    </div >
  )
}

export default ViewJava
