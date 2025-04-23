import React, { useRef } from 'react';
import { Table, Button, Tag, Switch, Input, Space } from '@douyinfe/semi-ui';
import { IconClock, IconSearch } from '@douyinfe/semi-icons';
import ViewSidePass from './side_pass';


const columns = [
  {
    title: '优先级',
    dataIndex: 'priority',
  },
  {
    title: '账号',
    dataIndex: 'account',
  },
  {
    title: '密码',
    dataIndex: 'password',
  },
  {
    title: '端口',
    dataIndex: 'port',
  },
  {
    title: '启用',
    dataIndex: 'port',
    render: (text: string) => {
      return <Switch onChange={(v, e) => console.log(v)} />
    }
  },
  {
    title: '创建时间',
    dataIndex: 'createAt',
    render: (text: string) => {
      return <Tag prefixIcon={<IconClock />} shape='circle' color='blue' style={{ userSelect: 'text' }}>{text}</Tag>;
    }
  },
  {
    title: '操作',
    dataIndex: 'action',
    render: (text: string) => {
      return (
        <div>
          <Button theme='outline' type='primary' style={{ marginRight: 8 }}>编辑</Button>
          <Button theme='outline' type='danger' style={{ marginRight: 8 }}>删除</Button>
        </div>
      );
    }
  },
];

const data = [
  {
    priority: 1,
    account: 'admin',
    password: '123456',
    port: 1234,
    createAt: '2021-01-01 12:00',
  },
  {
    priority: 2,
    account: 'root',
    password: 'abcd001002',
    port: 438,
    createAt: '2021-01-01 12:00',
  },
]



function ViewPass() {
  const sidePassRef = useRef<any>(null);

  // 新增或编辑密码组
  const addOrUpdatePass = () => {
    if (sidePassRef.current) {
      sidePassRef.current.change();
    }
  }

  return (
    <div style={{ marginTop: '16px' }}>
      <div>
        <ViewSidePass ref={sidePassRef} />
      </div>
      <div>
        <Space>
          <Input prefix={<IconSearch />} showClear placeholder="搜索内容"></Input>
          <Button theme='outline' type='primary' style={{ marginRight: 8 }} onClick={addOrUpdatePass}>新增</Button>
        </Space>
      </div>
      <div style={{ marginTop: 16 }}>
        <Table size='small' columns={columns} dataSource={data} pagination={{ pageSize: 3 }} />
      </div>
    </div>
  )
}

export default ViewPass
