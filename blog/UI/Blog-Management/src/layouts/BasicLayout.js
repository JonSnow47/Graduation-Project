/**
 * 2018-1-09 Jifeng Cheng
 * the main page of the blog backstage
 */

import React from 'react';
import { Layout, Menu, Icon } from 'antd';
import { Switch, Route, Link } from 'dva/router';

import MyEditor from '../routes/Article/MyEditor';
import ArticleManagement from '../routes/Management/ArticleManagement';
import ModifyArticle from '../routes/Modify/ModifyArticle';

const { SubMenu } = Menu;
const { Header, Content, Sider } = Layout;

class BasicLayout extends React.Component {
  render() {
    return (
      <Layout style={{minHeight: '100vh'}}>
        <Header className="header">
          <h2 style={{color: '#fff'}}>个人博客后台管理界面</h2>
        </Header>
        <Layout>
          <Sider width={200} style={{ background: '#fff' }}>
            <Menu
              mode="inline"
              defaultSelectedKeys={['1']}
              defaultOpenKeys={['sub1']}
              style={{ height: '100%', borderRight: 0 }}
            >
              <SubMenu key="sub1" title={<span><Icon type="edit" />写文章</span>}>
                <Menu.Item key="1">
                  <Link to="/main/editor">新建文章</Link>
                </Menu.Item>
                <Menu.Item key="2">
                  <Link to="/main/management">管理文章</Link>
                </Menu.Item>
                <Menu.Item>
                  修改文章
                </Menu.Item>
              </SubMenu>
              <SubMenu key="sub2" title={<span><Icon type="laptop" />主界面信息修改</span>}>
                <Menu.Item key="1">头部信息</Menu.Item>
                <Menu.Item key="2">个人简介</Menu.Item>
                <Menu.Item key="3">底部链接</Menu.Item>
                <Menu.Item key="4">标签</Menu.Item>
              </SubMenu>
            </Menu>
          </Sider>
          <Layout style={{ padding: '0 24px 24px' }}>
            <Content style={{ background: '#f0f2f5', padding: 24, margin: 10, minHeight: 280 }}>
              <Switch>
                <Route path="/main/editor" render={() => <MyEditor />} />
                <Route path="/main/management" render={() => <ArticleManagement />} />
                <Route path="/main/modify" render={() => <ModifyArticle />} />
              </Switch>
            </Content>
          </Layout>
        </Layout>
      </Layout>
    )
  }
}

export default BasicLayout;