import { TextArea, Col, Row } from "@douyinfe/semi-ui";
import ButtonGroup from "./button_group";
import TaskList from "./list";
import useActionsStore from "@/store/store_actions";

function ViewAction() {
  const { formIpmiInput, update } = useActionsStore();

  return (
    <div>
      <div>
        <Row gutter={16}>
          <Col span={12}>
            <ButtonGroup />
          </Col>
          <Col span={12}>
            <TextArea
              maxCount={100}
              showClear
              placeholder="请在此输入IPMI, 每行一个即可"
              value={formIpmiInput}
              onChange={(value) => {
                update({ formIpmiInput: value });
              }}
            />
          </Col>
        </Row>
      </div>
      <div style={{ marginTop: "16px" }}>
        <TaskList />
      </div>
    </div>
  );
}

export default ViewAction;
