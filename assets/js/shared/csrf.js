function getCsrfToken() {
  var fields = document.cookie.split(';');
  for (var i = 0; i < fields.length; ++i) {
    var kv = fields[i].split('=');
    if (kv[0] === "_csrf") {
      return kv[1];
    }
  }
  return null;
}

export { getCsrfToken };
