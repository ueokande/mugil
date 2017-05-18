import {combineReducers} from "redux"

const initialState = {
  open: false
}

function form(state = initialState, action) {
  switch (action.type) {
    case "FORM_OPEN_NEW":
      return Object.assign({}, state, {
        open: true,
      });
    case "FORM_CLOSE":
      return Object.assign({}, state, {
        open: false,
      });
    default:
      return state;
  }
}

export default combineReducers({
  form
});
