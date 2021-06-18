# Подключение сертификата
### !!! Способ подключения устарел, использовать только для ознакомления !!!

Создаем директорию **ssl** и переходим в нее  
Установка на локалхост:
```
SERVER_CN=localhost
```

1. generate cetificate authority + Trust cert (ca.crt)
```
openssl genrsa -passout pass:1111 -des3 -out ca.key 4096
openssl req -passin pass:1111 -new -x509 -days 365 -key ca.key -out ca.crt -subj "/CN=${SERVER_CN}"
```
2. generate the server private key (server.key)
```
openssl genrsa -passout pass:1111 -des3 -out server.key 4096
   ```
3. get a cert signing request from the CA (server.csr)
```
openssl req -passin pass:1111 -new -key server.key -out server.csr -subj "/CN=${SERVER_CN}"
```
4. sign the cert with the CA we created (self signing) - server.crt
```
openssl x509 -req -passin pass:1111 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt
```
5. convert the server cert to .pem format (server.pem) - for gRPC
```
openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem
```
Запуск сервера и клиента (см. пример в **03-greet**):
```
make s                          // server
GODEBUG=x509ignoreCN=0 make c   // client 
```