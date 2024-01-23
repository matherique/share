# 📤 Share 

Share any string (code, snippet, notes, etc) and get a link to share it

## Usage 

Start project
```bash
make run
```

Send hole file as `binary` to `localhost:8080`
```bash
curl localhost:8080 --data-binary "teste 123"
```
should return `htttp://localhost:8080/<id>`


If u curl the response url you should get the same sended data
```bash
curl htttp://localhost:8080/k12j41a
```


