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

import React from 'react';
import marked from 'marked';
import highlight from 'highlight.js';
import {
  Header,
  Segment,
  Divider,
  Label,
} from 'semantic-ui-react';
import ScrollReveal from 'scrollreveal';
import { connect } from 'dva';
import '../utils/atom.less';
import '../../node_modules/highlight.js/styles/atom-one-dark.css';
import styles from './ArticleShow.less';

class ArticleShow extends React.Component {
  componentWillMount() {
    marked.setOptions({
      highlight: code => highlight.highlightAuto(code).value,
    });
  }
  componentDidMount() {
    this.node.scrollIntoView();
    const config = {
      reset: false, // 滚动鼠标时，动画开关
      origin: 'left', // 动画开始的方向
      duration: 1500, // 动画持续时间
      delay: 0, // 延迟
      rotate: { x: 0, y: 0, z: 0 }, // 过度到0的初始角度
      opacity: 0, // 初始透明度
      scale: 0.2, //缩放
      viewFactor: 0.2,
      easing: 'cubic-bezier(0.6, 0.2, 0.1, 1)', // 缓动'ease', 'ease-in-out'，'linear'
    }
    ScrollReveal().reveal(this.refs.box1, config)
  }
  render() {
    const { Article, keys } = this.props;
    return (
      <div ref={node => (this.node = node)} className={styles.bgImage}>
        {
          Article.data === undefined ? null :
          <div ref='box1'>
            <Segment style={{ width: '80%', minHeight: '100vh', margin: 'auto', opacity: 0.9, backgroundColor: '#FBFBEA' }}>
              <Header>{Article.data[keys].Title}</Header>
              <div>
                <Label as="a" color="blue" tag>{Article.data[keys].Label1}</Label>
                <Label as="a" color="violet" tag>{Article.data[keys].Label2}</Label>
              </div>
              <Divider horizontal>2018-3-08</Divider>
              <div dangerouslySetInnerHTML={{ __html: marked(Article.data[keys].Content) }} />
            </Segment>
          </div>
        }
      </div>
    );
  }
}

export default connect(({ article }) => ({
  ...article,
}))(ArticleShow);
