import React from 'react';
import { Card, Select, Input, Space, Table, Button, Notification } from '@douyinfe/semi-ui';
import { IconDesktop, IconDescend, IconFilter, IconCopy, IconRefresh2 } from '@douyinfe/semi-icons';
import { IconGettingStarted } from '@douyinfe/semi-icons-lab';
import { options, optionsRules } from './fixed';
import columns from './columns';
import useSurveyStore from '@/store/store_survey';
import { verify_ip, generate_ip_list } from '@/utils/address';
import { isEmpty } from '@/utils/text';

// const data = [
//   {
//     ipmi: '192.168.1.1',
//     model: 'H3C S12500',
//     sn: '1234567890',
//     status: 'online',
//   },
//   {
//     ipmi: '192.168.1.2',
//     model: 'H3C S12500',
//     sn: '1234567890',
//     status: 'offline',
//   },
// ];
// // 循环加入data*10
// for (let i = 0; i < 3; i++) {
//   data.push(...data);
// }

const pagination = { pageSize: 10 };

const copyOnline = () => {
  Notification.success({
    title: '成功',
    content: '复制成功',
  });
}


const ViewSurvey = () => {
  const { ipmi, paragraph, filter, data, update, reset } = useSurveyStore();

  // 开始探测
  const startSurvey = () => {
    if (!verify_ip(ipmi) || isEmpty(paragraph)) {
      Notification.error({
        title: '错误',
        content: 'IPMI格式不合法或网段为空',
      });
      return;
    }
    try {
      const ip_list = generate_ip_list(ipmi, paragraph);
      let data_body = [];
      for (const ip of ip_list) {
        data_body.push({
          ipmi: ip,
          status: 'ready',
        })
      }
      update({ data: data_body });
    } catch (error: any) {
      Notification.error({
        title: '错误',
        content: error.message,
      });
    }
  }

  return (
    <div>
      <div>
        <Card>
          <Space>
            <Input prefix={<IconDesktop />} showClear style={{ width: 180 }} placeholder="输入IPMI" value={ipmi} onChange={(value) => update({ ipmi: value })} />
            <Select prefix={<IconDescend />} filter style={{ width: 150 }} placeholder="网段" optionList={options} value={paragraph} onChange={(value) => update({ paragraph: value as string })} />
            <Select prefix={<IconFilter />} filter style={{ width: 150 }} placeholder="过滤规则" optionList={optionsRules} value={filter} onChange={(value) => update({ filter: value as string })} />
            <Button theme='light' type='tertiary' style={{ marginRight: 8 }} icon={<IconRefresh2 />} onClick={reset}>重置</Button>
            <Button theme='light' type='primary' style={{ marginRight: 8 }} icon={<IconGettingStarted />} onClick={startSurvey}>开始探测</Button>
            <Button theme='light' type='secondary' style={{ marginRight: 8 }} icon={<IconCopy />} onClick={copyOnline}>复制在线</Button>
          </Space>
        </Card>
      </div>
      <div style={{ marginTop: 10 }}>
        <Table columns={columns} dataSource={data} pagination={pagination} />
      </div>
    </div>
  );
};

export default ViewSurvey;