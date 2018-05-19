/**
 * 2018-1-14 Jifeng Cheng
 */

import { message } from 'antd';
import { routerRedux } from 'dva/router';

import { getArticle } from '../services/article';
import { deleteArticle } from '../services/article';
import { modifyArticle } from '../services/article';

export default {
  namespace: 'article',

  state: {
    Article: [],
    loading: false,
    keys: null,
    modifyResult: null,
  },

  effects: {
    *showArticle({ payload }, { call, put, select }) {
      // const params = getArticle.data
      console.log('adasdasdsada')
      const response = yield call(getArticle);
      console.log('response......', response);
      yield put({
        type: 'getArticles',
        payload: response,
      })
    },

    *deleteArticle({ payload }, { call, put, select }) {
      console.log('payload+++++', payload)
      const data = {
        id: payload
      };
      // id.push(payload);
      const response = yield call(deleteArticle, data)
      
      if(response.status === 0) {
        message.success('删除成功！')
        const res = yield call(getArticle)
        console.log('deleteRes>>>>>>', res)
        yield put({
          type: 'updateArticle',
          payload: res,
        })
      } else {
        message.error('删除失败!')
      }
    },

    *getModifyInfo({ payload }, { put }) {
      yield put({
        type: 'modifyId',
        payload: payload,
      });
      yield put(routerRedux.push('/main/modify'))
    },

    *modifyArticle({ payload }, { call, put, select }) {
      console.log('payload>>>>>>>', payload)
      const tags = []
      tags.push(payload.tags)
      const params = {
        id: payload.id,
        title: payload.title,
        author: payload.author,
        tags: tags,
        brief: payload.brief,
        content: payload.content,
      }
      console.log('params???????', params)
      console.log('id>>>>>>>', params.id)
      const response = yield call(modifyArticle, params);
      console.log('modify answer+++++++', response);
      if (response.status === 0) {
        message.success('asd')
      } else {
        message.error('xzc')
      }
    },
  },

  reducers: {
    getArticles(state, action) {
      return {
        ...state,
        Article: action.payload,
      }
    },

    updateArticle(state, { payload }) {
      return {
        ...state,
        Article: payload,
      }
    },

    deleteId(state, action) {
      return {
        ...state,
        DeleteId: action.payload,
      }
    },

    modifyId(state, action) {
      return {
        ...state,
        keys: action.payload,
      }
    },

    modifyInfo(state, action) {
      return {
        ...state,
        modifyResult: action.payload,
      }
    }
  }
}