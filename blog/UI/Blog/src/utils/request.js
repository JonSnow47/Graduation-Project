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

import axios from 'axios';
import qs from 'qs';
import HttpStatus from 'http-status-codes';
// import {
//   baseURL,
//   requestTimeOut
// }                 from './config'

axios.defaults.baseURL = 'http://192.168.0.147:8080';
axios.defaults.timeout = 10000;
axios.defaults.withCredentials = true;

const fetch = (options) => {
  const {
    method,
    data,
    url,
  } = options;

  switch (method.toLowerCase()) {
    case 'get':
      return axios.get(`${url}${data ? `?${qs.stringify(data)}` : ''}`);
    case 'delete':
      return axios.delete(url, { data });
    case 'head':
      return axios.head(url, data);
    case 'post':
      return axios.post(url, data);
    case 'put':
      return axios.put(url, data);
    case 'patch':
      return axios.patch(url, data);
    default:
      return axios(options);
  }
};

export default function request(options) {
  return fetch(options).then((response) => {
    console.log('options: ', options, 'response: ', response);
    if (response.status === HttpStatus.OK) {
      return response.data;
    }
    throw { response } // eslint-disable-line
  }).catch((error) => {
    const { response } = error;
    console.log('request error: ', error);
    let message, status // eslint-disable-line
    if (response) {
      status = response.status;
      const { data, statusText } = response;
      message = data.message || statusText || HttpStatus.getStatusText(status);
    } else {
      status = 600;
      message = 'Network Error';
    }
    throw { status, message } // eslint-disable-line
  });
}

export const setToken = function (authToken) {
  axios.defaults.headers.common.Authorization = `Bearer ${authToken}`;
};
