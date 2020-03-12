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
