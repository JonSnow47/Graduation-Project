import dva from 'dva';
import createLoading from 'dva-loading';
// import { browserHistory } from 'dva/router';
import createHistory from 'history/createBrowserHistory';
import './index.css';
import article from './models/article';
import menu from './models/menu';
import router from './router';

// 1. Initialize
const app = dva({
  history: createHistory(),
});

// const loading = {
//   global: false,
//   models: {
//     article: false,
//   },
// };

// 2. Plugins
// app.use({});
app.use(createLoading());

// 3. Model
// app.model(require('./models/article'));
// app.model(require('./models/menu'));
app.model(article);
app.model(menu);

// 4. Router
app.router(router);

// 5. Start
app.start('#root');
