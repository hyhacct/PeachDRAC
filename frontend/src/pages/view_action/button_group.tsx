import { Tabs, TabPane, Button, Slider, Input, Space } from "@douyinfe/semi-ui";
import {
  IconSun,
  IconWifi,
  IconDesktop,
  IconBolt,
  IconSearch,
} from "@douyinfe/semi-icons";
import useActionsStore from "@/store/store_actions";

function ButtonGroup() {
  const { formNfs, formFan, update, Start, loading } = useActionsStore();

  const handleStart = (actions: string) => {
    Start(actions);
  };

  return (
    <div>
      <Tabs tabPosition="left" type="line">
        <TabPane
          tab={
            <span>
              <IconBolt />
              电源控制
            </span>
          }
          itemKey="1"
        >
          <div>
            <Button
              theme="solid"
              type="warning"
              style={{ margin: 5 }}
              onClick={() => handleStart("PowerOn")}
            >
              开机
            </Button>
            <Button
              theme="solid"
              type="warning"
              style={{ margin: 5 }}
              onClick={() => handleStart("PowerOff")}
            >
              优雅关机
            </Button>
            <Button
              theme="solid"
              type="warning"
              style={{ margin: 5 }}
              onClick={() => handleStart("PowerOff")}
            >
              强制关机
            </Button>
            <Button
              theme="solid"
              type="warning"
              style={{ margin: 5 }}
              onClick={() => handleStart("Restart")}
            >
              强制重启
            </Button>
          </div>
        </TabPane>
        <TabPane
          tab={
            <span>
              <IconDesktop />
              控制台
            </span>
          }
          itemKey="2"
        >
          <div>
            <Button
              theme="solid"
              type="primary"
              style={{ margin: 5 }}
              onClick={() => handleStart("启动Java控制台[踢掉其他会话]")}
            >
              启动Java控制台[踢掉其他会话]
            </Button>
            <Button
              theme="solid"
              type="primary"
              style={{ margin: 5 }}
              onClick={() => handleStart("StartJavaConsole")}
            >
              启动Java控制台[保留其他会话]
            </Button>
          </div>
        </TabPane>
        <TabPane
          tab={
            <span>
              <IconSun />
              温度控制
            </span>
          }
          itemKey="3"
        >
          <div>
            <Slider
              step={10}
              marks={{
                10: "10%",
                20: "20%",
                30: "30%",
                40: "40%",
                50: "50%",
                60: "60%",
                70: "70%",
                80: "80%",
                90: "90%",
                100: "100%",
              }}
              defaultValue={formFan}
              onChange={(value) => {
                update({ formFan: value as number });
              }}
            ></Slider>
            <div style={{ margin: "5%", width: "100%" }}>
              <Button
                theme="solid"
                style={{ margin: 5 }}
                block
                onClick={() => handleStart("手动调整")}
              >
                手动调整
              </Button>
              <Button
                theme="solid"
                style={{ margin: 5 }}
                block
                onClick={() => handleStart("恢复自动")}
              >
                恢复自动
              </Button>
            </div>
          </div>
        </TabPane>
        <TabPane
          tab={
            <span>
              <IconWifi />
              NFS控制
            </span>
          }
          itemKey="4"
        >
          <div>
            <Space vertical style={{ width: "100%" }}>
              <Input
                prefix={<IconSearch />}
                showClear
                placeholder="NFS服务器IP"
                value={formNfs}
                onChange={(value) => {
                  update({ formNfs: value });
                }}
              ></Input>
              <Button
                theme="solid"
                type="danger"
                block
                onClick={() => handleStart("确认挂载")}
              >
                确认挂载
              </Button>
              <Button
                theme="solid"
                type="danger"
                block
                onClick={() => handleStart("取消挂载")}
              >
                取消挂载
              </Button>
            </Space>
          </div>
        </TabPane>
      </Tabs>
    </div>
  );
}

export default ButtonGroup;
