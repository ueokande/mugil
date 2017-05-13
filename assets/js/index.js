import injectTapEventPlugin from "react-tap-event-plugin";

require('./login/index');
require('./task/index');

window.addEventListener('load', () => {
  injectTapEventPlugin();
});
