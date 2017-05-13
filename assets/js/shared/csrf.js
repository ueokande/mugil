function getCsrfToken() {
  var fields = document.cookie.split(';');
  for (var i = 0; i < fields.length; ++i) {
    var kv = fields[i].split('=');
    if (kv[0].trim() === "_csrf") {
      return kv[1].trim();
    }
  }
  return null;
}

export { getCsrfToken };
