import { getCsrfToken } from '../../shared/csrf'

function checkStatus(response) {
  if (response.status >= 200 && response.status < 300) {
    return response
  } else {
    var error = new Error(response.statusText)
    error.response = response
    throw error
  }
}

export function fetchTasks(date) {
  let params = { date: date};
  let query = Object.keys(params).reduce((a, k) => {
    a.push(k + '=' + window.encodeURIComponent(params[k]));
    return a
  }, []).join('&');
  let url = '/tasks?' + query;

  return (dispatch) => {
    fetch(url, {
      method: 'GET',
      credentials: 'same-origin',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    .then(checkStatus)
    .then((response) => {
      return response.json()
    })
    .then((entries) => {
      dispatch(updateTasks(entries))
    })
  }
}

export function postTask(priority, estimatedTime, description, successed) {
  return (dispatch) => {
    fetch('/tasks', {
      method: 'POST',
      credentials: 'same-origin',
      headers: {
        'Content-Type': 'application/json',
        'X-CSRF-TOKEN': getCsrfToken()
      },
      body: JSON.stringify({
        priority: priority,
        estimated_time: estimatedTime,
        description: description
      })
    })
    .then(checkStatus)
    .then((response) => {
      return response.json()
    })
    .then((task) => {
      if (typeof successed !== 'undefined') {
        successed(task);
      }
    })
  };
}

export function addTask(task) {
  return {
    type: "TASKS_ADD",
    task: task
  };
}

export function updateTasks(entries) {
  return {
    type: "TASKS_UPDATE",
    entries: entries
  };
}
