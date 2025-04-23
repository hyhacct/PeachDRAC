
import { Tabs, TabPane, Banner } from "@douyinfe/semi-ui";
import ViewJava from "./java";
import ViewPass from "./pass";

function ViewConfig() {

  const data = [
    {
      id: 1,
      tab: "爪哇配置",
      component: <ViewJava />,
    },
    {
      id: 2,
      tab: "密码组配置",
      component: <ViewPass />,
    },
  ]

  return (
    <div>
      <Tabs type="line">
        {data.map((item) => (
          <TabPane tab={item.tab} itemKey={item.id.toString()}>
            {item.component}
          </TabPane>
        ))}
      </Tabs>
    </div>
  )
}


export default ViewConfig
