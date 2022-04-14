# Third-Party Token Server

A proxy server to cache access tokens and which can be retreived by Golui for silent login.


## Endpoints

To cache the token;
```
POST /token
{"id":"1234", "token":"56789101112131415"}
```

To retrieve the token;
```
GET /token?id=1234

```
<br>

## Usage

Before redirecting the user to Golui the third-party app will Post the access token to cache and pass the link to retrieve the token to the token redirect URL

```
https://localhost:8444/v1/customer/third-party?token_url=www.https://<host>/token?id=1234
```
 

