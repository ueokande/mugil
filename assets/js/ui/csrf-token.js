import React from 'react';
import { getCsrfToken } from '../shared/csrf'

const CsrfToken = () => (
  <input type="hidden" name="CSRF_TOKEN" value={getCsrfToken()}/>
)

export default CsrfToken;
