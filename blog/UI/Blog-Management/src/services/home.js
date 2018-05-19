/**
 * 2018-1-11 Jifeng Cheng
 */

import request from '../utils/request';

export async function createAdmin() {
  return request('http://192.168.0.222:8080/admin/new', {
    method: 'POST',
    mode: 'cors',
    body: {
      name: 'wangriyu',
      password: '123456'
    },
  });
}


