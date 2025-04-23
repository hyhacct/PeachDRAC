import React, { useRef, useState, useEffect } from 'react';
import { Table, Button, Tag, Switch, Input, Space, Modal, Notification } from '@douyinfe/semi-ui';
import { IconClock, IconSearch } from '@douyinfe/semi-icons';
import ViewSidePass from './side_pass';
import { ConfigPassDelete, ConfigPassGetList } from '@/wailsjs/go/apps/App';
import WailsResp from '@/types/wails_resp';

function ViewPass() {
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
      dataIndex: 'port',
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
            <Button theme='outline' type='danger' style={{ marginRight: 8 }} onClick={() => deletePass(record.id)}>删除</Button>
          </div>
        );
      }
    },
  ];

  // const data = [
  //   {
  //     priority: 1,
  //     username: 'admin',
  //     password: '123456',
  //     port: 1234,
  //     createAt: '2021-01-01 12:00',
  //   },
  //   {
  //     priority: 2,
  //     username: 'root',
  //     password: 'abcd001002',
  //     port: 438,
  //     createAt: '2021-01-01 12:00',
  //   },
  // ]

  const [data, setData] = useState<any[]>([]);

  useEffect(() => {
    getAllList();
  }, []); // 空依赖数组确保只在挂载时调用一次

  const getAllList = async () => {
    try {
      const resp: WailsResp = await ConfigPassGetList();
      if (!resp.Status) {
        throw new Error(resp.Msg);
      }
      setData(resp.Data);
      console.log("data ========= ", resp.Data);
    } catch (error: any) {
      Notification.error({
        title: '错误',
        content: error?.message,
      });
    }
  }

  // 删除密码组
  const deletePass = async (id: number) => {
    try {
      Modal.error({
        title: '危险', content: '确定要删除密码组吗？删除后将无法恢复此记录, 请谨慎操作！',
        okButtonProps: {
          onClick: async () => {
            const resp: WailsResp = await ConfigPassDelete(id);
            if (!resp.Status) {
              throw new Error(resp.Msg);
            }
            Notification.success({
              title: '成功',
              content: resp.Msg,
            });
          }
        }
      });
    } catch (error: any) {
      Notification.error({
        title: '错误',
        content: error?.message,
      });
    }
  }


  const sidePassRef = useRef<any>(null);

  // 新增或编辑密码组
  const addOrUpdatePass = (row: any) => {
    if (sidePassRef?.current) {
      sidePassRef.current.change(row);
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
