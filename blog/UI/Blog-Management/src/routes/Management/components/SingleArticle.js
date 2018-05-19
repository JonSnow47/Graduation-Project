/**
 * 2018-1-10 Jifeng Cheng
 * single article info
 */

import React from 'react';
import { Tag, Button, Card } from 'antd';

import '../../../../node_modules/react-loading-bar/dist/index.css';
import styles from './SingleArticle.less';
import '../../../../node_modules/highlight.js/styles/atom-one-dark.css';

export default class SingleArticle extends React.Component {
  submitDeleteId = (id) => {
    this.props.deleteItem(id)
    console.log('调用了submitDeleteId!', id)
  }

  modifyId = (index) => {
    this.props.modifyItem(index)
    console.log('调用了modifyId！', index)
  }

  render() {
    return(
      <div key={this.props.Id} style={{ margin: '10px'}}>
        <Card>
          <div>
            <div>
              <h2>{this.props.Title}</h2>
              <div className={styles.tag}>
                <Tag color="#2db7f5">{this.props.Tags}</Tag>
                {/* <Tag color="#2db7f5">{this.props.Label2}</Tag> */}
              </div>
            </div>
            <hr />
            <div>{this.props.Brief}</div>
          </div>
          <div className={styles.tag}>
            <div>
              <Button
                type="danger"
                size="small"
                className={styles.button}
                onClick={() => this.submitDeleteId(this.props.Id)}
              >删除</Button>
            </div>
            <div>
              <Button
                size="small"
                className={styles.button}
                onClick={() => this.modifyId(this.props.index)}
              >修改</Button>
            </div>
          </div>
        </Card>
      </div>
    )
  }
}
