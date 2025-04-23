import React, { forwardRef, useImperativeHandle, useState } from 'react';
import { SideSheet, Button } from '@douyinfe/semi-ui';
import { Form, Tooltip } from '@douyinfe/semi-ui';
import { FormState } from '@douyinfe/semi-ui/lib/es/form';
import {
  ConfigPassCreate,
  ConfigPassDelete,
  ConfigPassUpdate
} from '@/wailsjs/go/apps/App';


const ViewSidePass = forwardRef((props, ref) => {
  useImperativeHandle(ref, () => ({
    change,
  }));


  // 控制侧边栏的显示
  const [visible, setVisible] = useState(false);
  const change = () => {
    setVisible(!visible);
  }

  const submit = () => {
    ConfigPassCreate(formData.Username, formData.Password, formData.Port);
  }

  const footer = () => {
    return (
      <div>
        <div style={{ display: 'flex', justifyContent: 'flex-end' }}>
          <Button style={{ marginRight: 8 }} onClick={change}>取消</Button>
          <Button theme="solid" type='primary' onClick={submit}>提交</Button>
        </div>
      </div>
    )
  }

  const [formData, setFormData] = useState({
    Username: '',
    Password: '',
    Port: '',
  });

  const handleFormChange = (formState: FormState) => {
    setFormData(formState.values);
  };

  return (
    <div>
      <SideSheet title="密码组" visible={visible} onCancel={change} footer={footer()}>
        <div>
          <Form initValues={formData} onChange={handleFormChange}>
            <Form.Input field='Username' label='用户名' placeholder="登录IPMI的用户名" />
            <Form.Input field='Password' label='密码' placeholder="登录IPMI的密码" />
            <Form.Input field='Port' label='端口' placeholder="一般默认就行" />
          </Form>
        </div>
      </SideSheet>
    </div>
  )
})

export default ViewSidePass
