/**
 * 2018-1-09 Jifeng Cheng
 * login page
 */

import React from 'react';
import { connect } from 'dva';
import { Layout, Form, Input, Icon, Button, Checkbox } from 'antd';
import styles from './LoginLayout.less';

const { Content } = Layout;
const FormItem = Form.Item;

class LoginLayout extends React.Component {
  handleSubmit = () => {
    const Username = document.getElementById('username').value;
    const Passward = document.getElementById('passward').value;
    const params = {
      name: Username,
      password: Passward,
    }
    this.props.dispatch({
      type: 'login/adminLogin',
      payload: params
    })
  }
  render() {
    // const { dispatch } = this.props
    // console.log('passward.....', Input);
    // 1.完整写法
    // (async ()=>{
    //   try {
    //     let res = await createAdmin();//等待fetch被resolve()后才能继续执行
    //     console.log("result:", res);
    //   } catch(e) {
    //     console.error(e);
    //   }
    // })();

    // 简写
    // AdminLogin({ name: "shjad", password: "123456" }).then(data => console.log("result: ", data))

    return (
      <Layout className="layout">
        <Content className={styles.bgimg}>
          <Form className={styles.login_form}>
            <FormItem>
              <Input
                prefix={<Icon type="user"
                style={{ color: 'rgba(0,0,0,.25)' }} />}
                placeholder="Username"
                id="username"
                />
            </FormItem>
            <FormItem>
              <Input
                prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />}
                type="password"
                placeholder="Password"
                id="passward"
              />
            </FormItem>
            <FormItem>
              <Checkbox>Remember me</Checkbox>
              <Button
                type="primary"
                htmlType="submit"
                className={styles.login_form_button}
                onClick={this.handleSubmit}
              >
                Log in
              </Button>
            </FormItem>
          </Form>
        </Content>
      </Layout>
    )
  }
}

export default connect(({login}) => ({login}))(LoginLayout);
