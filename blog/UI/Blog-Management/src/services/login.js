/**
 * 2018-1-10 Jifeng Cheng
 */
import request from '../utils/request';

export async function AdminLogin(params) {
  return request({
    url: '/admin/login',
    method: 'POST',
    data: params,
    // body: {
    //   name: params.name,
    //   pwd: params.pwd
    // }
  });
}

// import axios from 'axios';

// export async function AdminLogin(payload) {
//   console.log('qqqqqqq', payload)
//   const response = await axios.post('http://10.0.0.48:8080/adminLogin', {
//     name: payload.name,
//     password: payload.password,
//   })

//   return response.data
// }

