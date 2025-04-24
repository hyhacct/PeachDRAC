import React, { useEffect } from 'react';
import { Table, Button, Tag, Switch, Input, Space } from '@douyinfe/semi-ui';
import { IconClock, IconSearch } from '@douyinfe/semi-icons';
import ViewSidePass from './side_pass';
import useConfigStore from '@/store/store_config';


function ViewPass() {
  const { Delete, GetList, dataList, update } = useConfigStore();

  const columns = [
    {
      title: '优先级',
      dataIndex: 'priority',
    },
    {
      title: '账号',
      dataIndex: 'username',
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
      dataIndex: 'status',
      render: (text: string) => {
        return <Switch onChange={(v, e) => console.log(v)} />
      }
    },
    {
      title: '创建时间',
      dataIndex: 'created_at',
      render: (text: string) => {
        return <Tag prefixIcon={<IconClock />} shape='circle' color='blue' style={{ userSelect: 'text' }}>{text}</Tag>;
      }
    },
    {
      title: '操作',
      dataIndex: 'action',
      render: (text: string, record: any) => {
        return (
          <div>
            <Button theme='outline' type='primary' style={{ marginRight: 8 }} onClick={() => addOrUpdatePass(record)}>编辑</Button>
            <Button theme='outline' type='danger' style={{ marginRight: 8 }} onClick={() => Delete(record.id)}>删除</Button>
          </div>
        );
      }
    },
  ];

  useEffect(() => {
    GetList();
  }, []); // 空依赖数组确保只在挂载时调用一次


  // 新增或编辑密码组
  const addOrUpdatePass = (row: any | null) => {
    update({ show: true, form: row });
  }

  return (
    <div style={{ marginTop: '16px' }}>
      <div>
        <ViewSidePass />
      </div>
      <div>
        <Space>
          <Input prefix={<IconSearch />} showClear placeholder="搜索内容"></Input>
          <Button theme='outline' type='primary' style={{ marginRight: 8 }} onClick={() => addOrUpdatePass(null)}>新增</Button>
        </Space>
      </div>
      <div style={{ marginTop: 16 }}>
        <Table size='small' columns={columns} dataSource={dataList} pagination={{ pageSize: 10 }} />
      </div>
    </div>
  )
}

export default ViewPass
