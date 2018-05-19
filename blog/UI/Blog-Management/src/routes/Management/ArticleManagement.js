/**
 * 2018-1-10 Jifeng Cheng
 * management the articles
 */

import React from 'react';
import { connect } from 'dva';
import QueueAnim from 'rc-queue-anim';
import Loading from 'react-loading-bar';
import SingleArticle from './components/SingleArticle';

class ArticleManagement extends React.Component {
  componentDidMount() {
    this.props.dispatch({
      type: 'article/showArticle',
    })
  }

  handleDelete = (Id) => {
    console.log('id111>>>>>>', Id);
    this.props.dispatch({
      type: 'article/deleteArticle',
      payload: Id,
    })
  }

  getModifyId = (index) => {
    this.props.dispatch({
      type: 'article/getModifyInfo',
      payload: index
    })
  }
  render() {
    const { Article, loading } = this.props;
    console.log('Article_+_+_+_+_', Article);
    return(
      <QueueAnim delay={300}>
        <Loading
          show={loading}
          color="#99BBFF"
        />
        {
          Article.data === undefined ? null : Article.data.map((item, index) => (
            <div key={index}>
              <SingleArticle
                index={index}
                Id={item.Id}
                Title={item.Title}
                Label1={item.Label1}
                Label2={item.Label2}
                Brief={item.Brief}
                Content={item.Content}
                deleteItem={(id) => this.handleDelete(id)}
                modifyItem={(index) => this.getModifyId(index)}
              />
            </div>
          ))
        }
      </QueueAnim>
    )
  }
}

export default connect(state => ({
  Article: state.article.Article,
  loading: state.loading.models.article,
}))(ArticleManagement);
