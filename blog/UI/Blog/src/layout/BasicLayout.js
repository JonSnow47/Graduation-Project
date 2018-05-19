/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Co., Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2018/03/31        Cheng Jifeng
 */

import React, { Component } from 'react';
import {
  Button,
  Container,
  Grid,
  Header,
  List,
  Segment,
  Visibility,
  Menu,
  Input,
} from 'semantic-ui-react';
import { connect } from 'dva';
import 'semantic-ui-css/semantic.min.css';
import styles from './BasicLayout.less';
import Introduce from '../components/Introduce';
import ArticleInfo from '../components/Article/ArticleInfo';
import Pigeonhole from '../components/Pigeonhole';
import Tag from '../components/Tag';

const FixedMenu = () => (
  <Menu fixed="top" size="large" inverted pointing>
    <Container>
      <Menu.Item as="a" active>首页</Menu.Item>
      <Menu.Item as="a">留言板</Menu.Item>
      <Menu.Item as="a">Github</Menu.Item>
      <Menu.Menu position="right">
        <Input icon="search" placeholder="Search..." />
      </Menu.Menu>
    </Container>
  </Menu>
);

class BasicLayout extends Component {
  onShow = () => { // eslint-disable-line
    this.props.dispatch({
      type: 'menu/showFixedMenu',
    });
  }

  onHide = () => { // eslint-disable-line
    this.props.dispatch({
      type: 'menu/hideFixedMenu',
    });
  }

  render() {
    const { visible } = this.props;
    return (
      <div>
        { visible ? <FixedMenu /> : null }
        <Visibility
          onBottomPassed={this.onShow}
          onBottomVisible={this.onHide}
          once={false}
        >
          <div
            className={styles.bgImage}
          >
            <Header
              as="h1"
              textAlign="center"
              content="Cease to struggle and you cease to live"
              inverted
              style={{ fontSize: '4em', fontWeight: 'normal', marginBottom: 0, marginTop: '3em' }}
            />
            <Header
              as="h2"
              textAlign="center"
              content="Welcome to my Blog!"
              inverted
            />
            <Container
              textAlign="center"
            >
              <Button color="black" target="_blank" href="https://github.com/JonSnow47">Github</Button>
              <Button color="black" target="_blank" href="mailto:c745539141@gmail.com">e-mail</Button>
            </Container>
          </div>
        </Visibility>
        <h4 className="ui horizontal divider header"><i className="tag icon" /> Article </h4>

        <Segment vertical>
          <Grid container stackable>
            <Grid.Row>
              <Grid.Column width={10}>
                <ArticleInfo />
              </Grid.Column>
              <Grid.Column floated="right" width={4}>
                <Introduce />
                <Pigeonhole />
                <Tag />
              </Grid.Column>
            </Grid.Row>
          </Grid>
        </Segment>
        <Segment inverted vertical style={{ padding: '5em 0em' }}>
          <Container>
            <Grid divided inverted stackable>
              <Grid.Row>
                <Grid.Column width={3}>
                  <Header inverted as="h4" content="不错的网站" />
                  <List link inverted>
                    <List.Item as="a">Sitemap</List.Item>
                    <List.Item as="a">Contact Us</List.Item>
                    <List.Item as="a">Religious Ceremonies</List.Item>
                    <List.Item as="a">Gazebo Plans</List.Item>
                  </List>
                </Grid.Column>
                <Grid.Column width={3}>
                  <Header inverted as="h4" content="友情链接" />
                  <List link inverted>
                    <List.Item as="a">Banana Pre-Order</List.Item>
                    <List.Item as="a">DNA FAQ</List.Item>
                    <List.Item as="a">How To Access</List.Item>
                    <List.Item as="a">Favorite X-Men</List.Item>
                  </List>
                </Grid.Column>
                <Grid.Column width={7}>
                  <Header as="h4" inverted>Footer Header</Header>
                  <p>Extra space for a call to action inside the foote</p>
                </Grid.Column>
              </Grid.Row>
            </Grid>
          </Container>
        </Segment>
      </div>
    );
  }
}

export default connect(state => ({
  visible: state.menu.visible,
}))(BasicLayout);

