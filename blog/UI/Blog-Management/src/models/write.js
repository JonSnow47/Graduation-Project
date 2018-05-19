/**
 * 2018-1-11 Jifeng Cheng
 */

import { uploadArticle } from "../services/write";
import { message } from 'antd';

export default {
  namespace: 'write',
  
  state: {
  },

  effects: { 
    *articleSubmit({ payload }, { call, put }) {
      console.log('article.......', payload)
      const params = {
        title: payload.data.title,
        // tagsid: payload.data.tagsid,
        author: payload.data.author,
        // label2: payload.data.label2Input,
        brief: payload.data.brief,
        content: payload.data.content,
      }

      const response = yield call(uploadArticle, params);
      
      console.log('1111111', response);
      if (response.status === 0) {
        message.success('成功！')
      } else if (response.state === 401) {
        message.error('没有登录')
      } else {
        message.error('失败')
      }
    }
  },

  reducers: {
  },
};
