# Lucky Dine Session Store

First set an environment variable `LDT_SECRET` Make it large. 256 chars
`LOGIN_REDIRECT` the redirect for echo middleware to redirect to login page.
`SESSION_COOKIE` the name of the cookie that will store the token string.

1. Get a token and set your info like so.
```golang
  tokenString := New(map[string]string{"username": "bsmith", "admin": "false"})
```
2. Get the tokens info by decrypting it.
```golang
  // request is *http.Request
  // set to expire in 15 mins.
  isExpired, token, err := GetLdtTokenFromRequest(request)
  isExpired, token, err := GetLdtToken(tokenString)
```

3. Renew the token
```golang
  // request is *http.Request
  // will renew expired tokens. Leaving up to application with isExpired.
  isExpired, tokenString, err := RenewLdtTokenFromRequest(request)
  isExpired, tokenString, err := RenewLdtToken(tokenString)
  ```
### Echo Middleware

* ValidateSession - validates whether the token or cookie is expired and sets `custom_token` object in context to be retrieved by context.Get("custom_token") (if valid).  Will return 401 Unauthorized (not valid)

* ValidateSessionOrRedirectToLogin - validates like function above but instead of returning 401 if not validated. Redirects to a login page with a query string paramter of `forward` and the path the user was trying to reach.

* RenewSession - can only be used after one of the above middleware functions and only with a cookie version of the session.
