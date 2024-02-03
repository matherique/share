# 📤 Share 

Create a shareable link from any string (code, snippet, notes, etc)

## Usage 

To send a file just send it as binary to the endpoint `/`
```bash
curl -X POST localhost:8080 --data-binary "teste 123"
htttp://localhost:8080/<id>
```

To send an encrypted file use `/secure`
```bash
curl -X POST localhost:8080/secure --data-binary "teste 123"
link: htttp://localhost:8080/<id>
key: random_key==
```

To return the file use the endpoint returned by the create endpoint
```bash
curl -X GET http://localhost:8080/<id>
```
To retrive the secure data use `/secure/<id>` endpoint and add key in `Authorization` header
