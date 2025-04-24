import { Tag, Banner, Card, Space, Empty } from "@douyinfe/semi-ui";
import { Descriptions, Col, Row, Typography, Button } from '@douyinfe/semi-ui';
import { IconRating } from "@douyinfe/semi-icons-lab";
import { IllustrationNoContent, IllustrationNoContentDark } from '@douyinfe/semi-illustrations';
import ViewSideJava from './side_java';
import useConfigStore from '@/store/store_config_java';
import { splitIp } from '@/utils/text';
import { model } from '@wails/go/models';

function ViewJava() {
  const { Meta } = Card;
  const { Text } = Typography;

  const { dataList, update, resetForm, reset, show, form, Delete, AddOrUpdate } = useConfigStore();

  const open = (row: model.TableJava | null) => {
    update({ show: true, form: row || undefined });
  }

  const style = {
    backgroundColor: 'var(--semi-color-bg-2)',
    borderRadius: '10px',
  };

  return (
    <div style={{ marginTop: '16px' }}>
      <Banner
        type="info"
        description="若需要为不同厂商或者IPMI指定启动控制台的Java版本, 请点击[新增配置]按钮, 若无配置则使用系统默认的Java版本"
      />
      <ViewSideJava />
      <div style={{ marginTop: '16px' }}>
        <Row gutter={[10, 10]}>
          <Col span={8}>
            <div >
              <Card
                style={style}
                bordered={false}
                headerLine={true}
                shadows='hover'
              >
                <Button theme='borderless' type='primary' style={{ height: '100%', width: '100%' }} onClick={() => open(null)}>
                  <Empty
                    image={<IllustrationNoContent />}
                    darkModeImage={<IllustrationNoContentDark style={{ height: 107 }} />}
                    title={'新增配置'}
                    description="点击添加一个新的配置"
                  />
                </Button>
              </Card>
            </div>
          </Col>
          {dataList.map((item) => (
            <Col span={8}>
              <div >
                <Card
                  style={style}
                  bordered={false}
                  headerLine={true}
                  title={
                    <Meta
                      title={item.title}
                      description={item.created_at}
                      avatar={<IconRating size="extra-large" />}
                    />
                  }
                  headerExtraContent={
                    <div>
                      <Space wrap>
                        <Text link onClick={() => open(item)}>
                          编辑
                        </Text>
                        <Text link type="danger" onClick={() => Delete(item.id)}>
                          删除
                        </Text>
                      </Space>
                    </div>
                  }
                >
                  <Descriptions>
                    <Descriptions.Item itemKey="Java路径">{item.path}</Descriptions.Item>
                    <Descriptions.Item itemKey="应用型号">{item.allot}</Descriptions.Item>
                    <Descriptions.Item itemKey="应用IPMI">
                      {item.ips && splitIp(item.ips).length != 0 ?
                        <div>
                          <Tag shape='circle' color='violet'>
                            {splitIp(item.ips)[0]}
                          </Tag>
                          <Tag shape='circle' style={{ marginLeft: '8px' }}>
                            {"+" + splitIp(item.ips).length + "台"}
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
        </Row>
      </div>
    </div >
  )
}

export default ViewJava
