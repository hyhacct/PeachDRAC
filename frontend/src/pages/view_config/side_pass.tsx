import React, { forwardRef, useImperativeHandle, useState } from 'react';
import { SideSheet, Button, Notification } from '@douyinfe/semi-ui';
import { Form, Tooltip } from '@douyinfe/semi-ui';
import { FormState } from '@douyinfe/semi-ui/lib/es/form';
import {
  ConfigPassAddOrUpdate,
} from '@/wailsjs/go/apps/App';
import WailsResp from '@/types/wails_resp';


const ViewSidePass = forwardRef((props, ref) => {
  useImperativeHandle(ref, () => ({
    change,
  }));


  // 控制侧边栏的显示
  const [visible, setVisible] = useState(false);
  const change = (row: any) => {
    if (row) {
      setFormData(row);
    }
    setVisible(!visible);
  }

  const submit = async () => {
    try {
      const resp: WailsResp = await ConfigPassAddOrUpdate(formData.id, formData.username, formData.password, formData.port);
      if (!resp.Status) {
        throw new Error(resp.Msg);
      }
      change(null);
      Notification.success({
        title: '成功',
        content: resp.Msg,
      });
    } catch (error: any) {
      Notification.error({
        title: '错误',
        content: error?.message,
      });
    }
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
    id: 0,
    username: '',
    password: '',
    port: '',
    enable: true,
  });

  const handleFormChange = (formState: FormState) => {
    setFormData(formState.values);
  };

  return (
    <div>
      <SideSheet title="密码组" visible={visible} onCancel={change} footer={footer()}>
        <div>
          <Form initValues={formData} onChange={handleFormChange}>
            <Form.Input field='username' label='用户名' placeholder="登录IPMI的用户名" />
            <Form.Input field='password' label='密码' placeholder="登录IPMI的密码" />
            <Form.Input field='port' label='端口' placeholder="一般默认就行" />
          </Form>
        </div>
      </SideSheet>
    </div>
  )
})

export default ViewSidePass
