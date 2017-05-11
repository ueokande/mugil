function csrfToken() {
  var fields = document.cookie.split(';');
  for (var i = 0; i < fields.length; ++i) {
    var kv = fields[i].split('=');
    if (kv[0] === "_csrf") {
      return kv[1];
    }
  }
  return null;
}

window.onload = function() {
  var input = document.querySelector("[name=CSRF_TOKEN]");
  if (input === null) {
    return;
  }
  input.value = csrfToken();
};
