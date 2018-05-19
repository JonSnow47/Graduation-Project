/**
 * 2017-1-10 Jifeng CHeng
 * the markdown editor by marked && simplemde && highlight.js
 */

import React from 'react';
import { connect } from 'dva';
import SimpleMDE from 'simplemde'
import marked from 'marked';
import highlight from 'highlight.js';
import { Button } from 'antd';
import styles from './MyEdtor.less';
// import '/Users/a8/github/React/Myblog-Backstage-management/Blog-Management/node_modules/highlight.js/styles/atom-one-dark.css';

class MyEditor extends React.Component {
  handleSubmit = () => {
    const tags = []
    tags.push(document.getElementById('tags').value)
    const inputArticle = {
      title: document.getElementById('title').value,
      author: document.getElementById('author').value,
      tags: tags,
      brief: document.getElementById('brief').value,
      content: this.smde.value(),
    }
    console.log('inputContent......', inputArticle.contentInput)

    const params = {
      data: inputArticle,
    }

    this.props.dispatch({
      type: 'write/articleSubmit',
      payload: params,
    })
  }
  
  componentDidMount() {
    this.smde = new SimpleMDE({
      element: document.getElementById('editor'), 
      indentWithTabs: false,
      tabSize: 2,
      status: ["autosave", "lines", "words", "cursor"],
      previewRender: function(plainText) {
        return marked(plainText,{
          // autofocus: true,
          renderer: new marked.Renderer(),
          gfm: true,
          pedantic: false,
          sanitize: false,
          tables: true,
          breaks: true,
          smartLists: true,
          smartypants: true,
          highlight: function (code) {
            return highlight.highlightAuto(code).value;
          }
        });
      },
    })
  }
  render() {
    return(
      <div>
        <div className={styles.head}>
          <input
            style={{ marginBottom: 30 }}
            placeholder="标题"
            ref={(input) => {this.input = input}}
            id="title"
          />
        </div>

        <div className={styles.head}>
          <input
            style={{ marginBottom: 30 }}
            placeholder="作者"
            id="author"
          />
          <input
            style={{ marginBottom: 30 }}
            placeholder="标签"
            id="tags"
          />
          <input
            style={{ marginBottom: 30 }}
            placeholder="简介"
            id="brief"
          />
        </div>

        <div>
          <textarea
            id="editor"
            placeholder="文章内容"
          />
        </div>
        <div>
          <Button
            type="primary"
            size="small"
            onClick={this.handleSubmit}
          >发布</Button>
        </div>
      </div>
    )
  }
}

export default connect(({ write }) => ({ write }))(MyEditor);