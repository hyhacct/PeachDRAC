import { SideSheet, Button } from '@douyinfe/semi-ui';
import { Form } from '@douyinfe/semi-ui';
import useConfigStore from '@/store/store_config_java';


const ViewSideJava = () => {
  const { form, show, AddOrUpdate, update } = useConfigStore();

  // 处理表单值变化
  const handleFormChange = (values: any) => {
    update({ form: { ...form, ...values } });
  };

  const footer = () => {
    return (
      <div>
        <div style={{ display: 'flex', justifyContent: 'flex-end' }}>
          <Button style={{ marginRight: 8 }} onClick={() => update({ show: false })}>取消</Button>
          <Button theme="solid" type='primary' onClick={AddOrUpdate}>提交</Button>
        </div>
      </div>
    )
  }

  return (
    <div>
      <SideSheet title="Java配置" visible={show} onCancel={() => update({ show: false })} footer={footer()}>
        <div>
          <Form initValues={form} onValueChange={handleFormChange}>
            <Form.Input field='title' label='标题' placeholder="为这份配置起个名字吧" />
            <Form.Input field='path' label='Java路径(需要完整路径,如下)' placeholder="C:/Windows/Java/jdk8/javaws.exe" />
            <Form.Input field='allot' label='厂商(可空)' placeholder="Dell" />
            <Form.Input field='ips' label='适用IPMI(可空,多个IPMI用逗号隔开)' placeholder="192.168.1.1,192.168.1.2" />
          </Form>
        </div>
      </SideSheet>
    </div>
  )
}

export default ViewSideJava
