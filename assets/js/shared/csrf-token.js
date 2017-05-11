import React from 'react';

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

const CsrfToken = () => (
  <input type="hidden" name="CSRF_TOKEN" value={csrfToken()}/>
)

export default CsrfToken;
