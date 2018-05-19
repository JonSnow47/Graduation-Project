/**
 * 2018-1-11 Jifeng Cheng
 */

// import axios from 'axios';

// export async function uploadArticle(payload) {
//   const { inputTitle, inputLabel, inputContent } = payload;
//   const response = await axios.post('http://10.0.0.48:8080/insert', {
//     title: payload.inputTitle,
//     label: inputLabel,
//     content: inputContent,
//   });

//   return response.data;
// }

import request from '../utils/request';

export async function uploadArticle(params) {
  console.log('wwwwww', params)
  return request({
    url: '/blog/article/new',
    method: 'post',
    data: params,
  });
}