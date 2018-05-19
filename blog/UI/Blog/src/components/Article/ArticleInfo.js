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
import { connect } from 'dva';
import {
  Card,
  Header,
  Image,
  Label,
  Button,
  Divider,
} from 'semantic-ui-react';
import ScrollReveal from 'scrollreveal';

// import sr from '../../components/ScrollReveal';
import styles from './ArticleInfo.less';
import '../../../node_modules/highlight.js/styles/atom-one-dark.css';
import images from '../../utils/images';

class ArticleInfo extends React.Component {
  componentWillMount() {
    this.props.dispatch({
      type: 'article/showArticle',
    });
    marked.setOptions({
      highlight: code => highlight.highlightAuto(code).value,
    });
  }
  componentDidMount() {
    const config = {
      reset: false, // 滚动鼠标时，动画开关
      origin: 'left', // 动画开始的方向
      duration: 1000, // 动画持续时间
      delay: 0, // 延迟
      rotate: { x: 0, y: 0, z: 0 }, // 过度到0的初始角度
      opacity: 0, // 初始透明度
      scale: 0.2, //缩放
      easing: 'cubic-bezier(0.6, 0.2, 0.1, 1)', // 缓动'ease', 'ease-in-out'，'linear'
    }
    ScrollReveal().reveal(this.refs.box1, config)
  }
  readArticles = (id) => { // eslint-disable-line
    this.props.dispatch({
      type: 'article/readMore',
      payload: id,
    });
  }
  render() {
    const { Article, loading } = this.props;
    console.log('loading>>>>>', loading);
    return (
      <div ref='box1'>
        {
          Article.data === undefined ? null : Article.data.map((item, index) => (
            <Card
              fluid
              key={index}
              // ref='box1'
            >
              <Image
                style={{ height: 250 }}
                src={images.headImage[Math.floor(Math.random() * images.headImage.length)]}
              />
              <Card.Content>
                <Header>{item.Title}</Header>
                <div className={styles.type}>
                  <div>
                    <Label as="a" color="blue" tag>{item.Tags[0]}</Label>
                    <Label as="a" color="violet" tag>{item.Tags[1]}</Label>
                  </div>
                </div>
                <Divider horizontal>2018-3-08</Divider>
                <div className={styles.briefInfo}>{item.Brief}</div>
                <Button
                  content="阅读全文"
                  size="small"
                  color="black"
                  style={{
                    float: 'right',
                    padding: 10,
                    marginTop: 15,
                  }}
                  onClick={() => this.readArticles(index)}
                />
              </Card.Content>
            </Card>
          ))
        }
      </div>
    );
  }
}

export default connect(state => ({
  Article: state.article.Article,
  loading: state.loading.models.article,
}))(ArticleInfo);
