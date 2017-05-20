const initialState = {
  loading: false,
  entries: [],
}

export default function tasks(state = initialState, action) {
  switch (action.type) {
    case "TASKS_ADD":
      return Object.assign({}, state, {
        entries: state.entries.concat(action.task)
      });
    case "TASKS_FETCH":
      return Object.assign({}, state, {
        loading: true
      });
    case "TASKS_UPDATE":
      return Object.assign({}, state, {
        entries: action.entries,
        loading: false,
      });
    default:
      return state;
  }
}

