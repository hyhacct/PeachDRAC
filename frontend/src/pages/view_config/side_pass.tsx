import { SideSheet, Button } from '@douyinfe/semi-ui';
import { Form } from '@douyinfe/semi-ui';
import useConfigStore from '@/store/store_config_pass';


const ViewSidePass = () => {
  const { form, show, AddOrUpdate, update } = useConfigStore();

  // 处理表单值变化
  const handleFormChange = (values: any) => {
    console.log('Form values changed:', values);
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
      <SideSheet title="密码组" visible={show} onCancel={() => update({ show: false })} footer={footer()}>
        <div>
          <Form initValues={form} onValueChange={handleFormChange}>
            <Form.InputNumber field='priority' label='优先级' placeholder="越高越优先" min={1} max={1000} style={{ width: '100%' }} />
            <Form.Input field='username' label='用户名' placeholder="登录IPMI的用户名" />
            <Form.Input field='password' label='密码' placeholder="登录IPMI的密码" />
            <Form.InputNumber field='port' label='端口' placeholder="一般默认就行" min={1} max={65535} style={{ width: '100%' }} />
          </Form>
        </div>
      </SideSheet>
    </div>
  )
}

export default ViewSidePass
