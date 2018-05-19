import dynamic from 'dva/dynamic';

// wrapper of dynamic(动态导入)
const dynamicWrapper = (app, models, component) => dynamic({
  app,
  moddels: () => models.map(m => import(`../models/${m}.js`)),
  component: () => component,
});

// data
export const getNavData = app => [
  {
    component: dynamicWrapper(app, [''], import('../layout/BasicLayout.js')),
    layout: 'BasicLayout',
    name: '我的博客首页',
    path: '/',
  },
];
