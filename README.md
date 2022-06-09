# Third-Party Token Server

Prototype for a server to cache access tokens 


## Endpoints

To cache the token;
```
POST /token
{"id":"1234", "token":"56789101112131415"}
```

To retrieve the token;
```
GET /token?id=1234
response; 
   200 {"token":"56789101112131415" }
```
<br>

