import { Nav, Layout } from '@douyinfe/semi-ui';
import { IconList, IconConfig, IconDivider, IconCodeHighlight } from '@douyinfe/semi-icons-lab';
import { Route, Routes, Link } from 'react-router-dom';
import ViewAction from '@/pages/view_action';
import ViewSurvey from '@/pages/view_survey';
import ViewSetting from '@/pages/view_seting';
import ViewConfig from '@/pages/view_config';

export const ViewLayout = () => {
  const { Sider, Content } = Layout;

  const LeftNav = () => (
    <Nav
      header={{
        // logo: <IconSemiLogo style={{ height: '36px', fontSize: 36 }} />,
        text: 'PeachDRAC'
      }}
      style={{ maxWidth: 150, height: '100vh' }}
      defaultSelectedKeys={['item-1']}
      renderWrapper={({ itemElement, isSubNav, isInSubNav, props }) => {
        const routerMap = {
          "item-1": "/action",
          "item-2": "/survey",
          "item-3": "/config",
          "item-4": "/setting",
        };
        return (
          <Link
            style={{ textDecoration: "none" }}
            to={routerMap[props.itemKey as keyof typeof routerMap]}
          >
            {itemElement}
          </Link>
        );
      }}
      items={[
        {
          itemKey: 'item-1',
          text: '批量动作',
          icon: <IconList />,
        },
        {
          itemKey: 'item-2',
          text: '探测扫描',
          icon: <IconDivider />,
        },
        {
          itemKey: 'item-3',
          text: '配置管理',
          icon: <IconCodeHighlight />,
        },
        {
          itemKey: 'item-4',
          text: '系统设置',
          icon: <IconConfig />,
        },
      ]}
      footer={{
        collapseButton: true,
      }}
    />
  );

  return (
    <Layout style={{ border: '1px solid var(--semi-color-border)' }}>
      <Layout>
        <Sider style={{ backgroundColor: 'var(--semi-color-bg-1)' }}>
          <LeftNav />
        </Sider>
        <Content
          style={{
            padding: '10px',
            backgroundColor: 'var(--semi-color-bg-0)',
          }}
        >
          <div
            style={{
              borderRadius: '10px',
              border: '1px solid var(--semi-color-border)',
              padding: '32px',
            }}
          >
            <Routes>
              <Route path="/action" element={<ViewAction />} />
              <Route path="/survey" element={<ViewSurvey />} />
              <Route path="/config" element={<ViewConfig />} />
              <Route path="/setting" element={<ViewSetting />} />
            </Routes>
          </div>
        </Content>
      </Layout>
    </Layout>
  );
};
