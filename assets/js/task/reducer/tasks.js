const initialState = {
  entries: []
}

export default function tasks(state = initialState, action) {
  switch (action.type) {
    case "TASKS_UPDATE":
      return Object.assign({}, state, {
        entries: action.entries
      });
    default:
      return state;
  }
}

