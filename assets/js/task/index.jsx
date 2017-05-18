import React from 'react';
import ReactDOM from 'react-dom';
import TaskIndex from './component/task-index';

window.addEventListener('load', () => {
  var parent = document.getElementById('mugil-task-index');
  if (parent !== null) {
    ReactDOM.render(
      <TaskIndex />,
      parent
    );
  }
});
