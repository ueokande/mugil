import { expect } from "chai";
import { getCsrfToken } from './csrf'

describe("#getCsrfToken", () => {
  it("return CSRF token from cookie", () => {
    document.cookie = "apple=red";
    document.cookie = "_csrf=my_csrf_token";
    document.cookie = "banana=yellow";

    expect(getCsrfToken()).to.equal("my_csrf_token");
  });

  it("return null if _csrf is not in cookie", () => {
    document.cookie = "apple=red";
    document.cookie = '_csrf=dummy; Expires=Thu, 01 Jan 1970 00:00:01 GMT;' // remove _csrf
    document.cookie = "banana=yellow";

    expect(getCsrfToken()).to.be.null;
  });
})
