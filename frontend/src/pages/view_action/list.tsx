import { List, Divider, Descriptions, Tag, Space } from "@douyinfe/semi-ui";
import useActionsStore from "@/store/store_actions";
import { IconVerify, IconHourglass, IconClear } from "@douyinfe/semi-icons";

const TaskList = () => {
  const style = {
    border: "1px solid var(--semi-color-border)",
    backgroundColor: "var(--semi-color-bg-2)",
    borderRadius: "3px",
    paddingLeft: "20px",
    margin: "8px 2px",
  };

  const { dataList } = useActionsStore();

  const getStatusComponent = (status: string) => {
    const body = {
      success: (
        <Tag color="teal" prefixIcon={<IconVerify />}>
          成功
        </Tag>
      ),
      error: (
        <Tag color="pink" prefixIcon={<IconClear />}>
          失败
        </Tag>
      ),
      ready: (
        <Tag color="blue" prefixIcon={<IconHourglass />}>
          进行中
        </Tag>
      ),
    };
    return body[status as keyof typeof body];
  };

  return (
    <div>
      <div>
        <Divider margin="12px" />
      </div>
      <div>
        <List
          grid={{
            gutter: 12,
            span: 6,
          }}
        >
          {dataList.map((item) => (
            <List.Item style={style}>
              <div>
                <h3
                  style={{ color: "var(--semi-color-text-0)", fontWeight: 500 }}
                >
                  <Space>
                    {getStatusComponent(item.status)}
                    {item.ipmi}
                  </Space>
                </h3>
                <Descriptions
                  align="left"
                  size="small"
                  column={1}
                  data={[
                    {
                      key: "操作",
                      value: item.action,
                    },
                    { key: "厂商", value: item.manufacturer },
                    { key: "型号", value: item.model },
                    { key: "SN", value: item.sn },
                    { key: "消息", value: item.message },
                  ]}
                />
              </div>
            </List.Item>
          ))}
        </List>
      </div>
    </div>
  );
};

export default TaskList;
