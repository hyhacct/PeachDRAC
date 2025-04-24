import { IconTick, IconClose, IconLoading } from '@douyinfe/semi-icons';
import { Tag, Tooltip } from '@douyinfe/semi-ui';
import { cutOff } from '@/utils/text';


const columns = [
  {
    title: 'IPMI',
    dataIndex: 'ipmi',
  },
  {
    title: '状态',
    dataIndex: 'status',
    render: (text: string) => {
      const dataMap: any = {
        online: {
          text: '在线',
          color: 'cyan',
          icon: <IconTick />,
        },
        offline: {
          text: '离线',
          color: 'violet',
          icon: <IconClose />,
        },
        ready: {
          text: '就绪',
          color: 'amber',
          icon: <IconLoading />,
        },
      }
      return <Tag type='ghost' size='large' prefixIcon={dataMap[text].icon} shape='circle' color={dataMap[text].color}>{dataMap[text].text}</Tag>;
    },
  },
  {
    title: '型号',
    dataIndex: 'model',
  },
  {
    title: '序列号',
    dataIndex: 'sn',
  },
  {
    title: '消息',
    dataIndex: 'msg',
    render: (text: string) => {
      return (
        <Tooltip content={text}
          arrowPointAtCenter={false}
          position='top'
        >
          {cutOff(text, 20)}
        </Tooltip>
      )

    },
  },
];

export default columns;