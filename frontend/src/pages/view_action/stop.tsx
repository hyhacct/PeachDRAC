import { Banner, Button, Space, Row, Col } from "@douyinfe/semi-ui";
import { IconHelpCircle } from "@douyinfe/semi-icons";
import { getRandomText } from "@/utils/doyouknow";
import useActionsStore from "@/store/store_actions";

function Stop() {
  const { isExiting, Stop } = useActionsStore();

  return (
    <div>
      <Space vertical align="center" style={{ width: "100%" }}>
        <Banner
          fullMode={false}
          type="info"
          bordered
          icon={<IconHelpCircle />}
          closeIcon={null}
          title={
            <div
              style={{
                fontWeight: 600,
                fontSize: "14px",
                lineHeight: "20px",
              }}
            >
              你知道吗？
            </div>
          }
          description={<div>{getRandomText()}</div>}
        />
        <Button
          type="primary"
          loading={isExiting}
          onClick={() => {
            Stop();
          }}
        >
          手动结束任务
        </Button>
      </Space>
    </div>
  );
}

export default Stop;
