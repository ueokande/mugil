import React from 'react';
import ReactDOM from 'react-dom';
import {Provider} from "react-redux"

import TaskIndex from './component/task-index';
import createStoreWithMiddleware from './store/configure';

const store = createStoreWithMiddleware();

window.addEventListener('load', () => {
  var parent = document.getElementById('mugil-task-index');
  if (parent !== null) {
    ReactDOM.render(
      <Provider store={store}>
        <TaskIndex />
      </Provider>,
      parent
    );
  }
});
