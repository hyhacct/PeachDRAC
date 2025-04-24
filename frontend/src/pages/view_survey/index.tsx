import { Card, Select, Input, Space, Table, Button, Notification, Banner } from '@douyinfe/semi-ui';
import { IconDesktop, IconDescend, IconFilter, IconCopy, IconRefresh2, IconHelpCircle } from '@douyinfe/semi-icons';
import { IconGettingStarted } from '@douyinfe/semi-icons-lab';
import { options, optionsRules } from './fixed';
import columns from './columns';
import useSurveyStore from '@/store/store_survey';
import { verify_ip, generate_ip_list } from '@/utils/address';
import { isEmpty } from '@/utils/text';
import { getRandomText } from '@/utils/doyouknow';

const pagination = { pageSize: 10 };

const copyOnline = () => {
  Notification.success({
    title: '成功',
    content: '复制成功',
  });
}


const ViewSurvey = () => {
  const { ipmi, paragraph, filter, dataList, update, reset, Start, loading } = useSurveyStore();

  // 开始探测
  const startSurvey = async () => {
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
      update({ dataList: data_body as any });
      await Start(ip_list); // 开始探测
    } catch (error: any) {
      Notification.error({
        title: '错误',
        content: error.message,
      });
    }
  }

  return (
    <div>
      <div style={{ marginTop: 10 }}>
        <Card>
          <Space>
            <Input disabled={loading} prefix={<IconDesktop />} showClear style={{ width: 180 }} placeholder="输入IPMI" value={ipmi} onChange={(value) => update({ ipmi: value })} />
            <Select loading={loading} prefix={<IconDescend />} filter style={{ width: 150 }} placeholder="网段" optionList={options} value={paragraph} onChange={(value) => update({ paragraph: value as string })} />
            <Select loading={loading} prefix={<IconFilter />} filter style={{ width: 150 }} placeholder="过滤规则" optionList={optionsRules} value={filter} onChange={(value) => update({ filter: value as string })} />
            <Button loading={loading} theme='light' type='tertiary' style={{ marginRight: 8 }} icon={<IconRefresh2 />} onClick={reset}>重置</Button>
            <Button loading={loading} theme='light' type='primary' style={{ marginRight: 8 }} icon={<IconGettingStarted />} onClick={startSurvey}>开始探测</Button>
            <Button loading={loading} theme='light' type='secondary' style={{ marginRight: 8 }} icon={<IconCopy />} onClick={copyOnline}>复制在线</Button>
          </Space>
        </Card>
      </div>
      {/* loading && */}
      {<div style={{ marginTop: 10 }}>
        <Banner fullMode={false} type="info" bordered icon={<IconHelpCircle />} closeIcon={null}
          title={<div style={{ fontWeight: 600, fontSize: '14px', lineHeight: '20px' }}>你知道吗？</div>}
          description={<div>{getRandomText()}</div>}
        />
      </div>}
      <div style={{ marginTop: 10 }}>
        <Table columns={columns} dataSource={dataList} pagination={pagination} loading={loading} />
      </div>
    </div>
  );
};

export default ViewSurvey;