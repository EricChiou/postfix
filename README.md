## Postfix and Golang
Send Email by Postfix, made with Golang.
<br><br/>

### Install Postfix
``` console
sudo apt-get install postfix
```

### Postfix config
``` console
sudo vim /etc/postfix/main.cf
```

### Check Postfix config as below
``` console
inet_interfaces = all
inet_protocols = all
```
<br><br/>

## API
```
// send text mail
POST http://127.0.0.1:26/postfix/send/text

// send html mail
POST http://127.0.0.1:26/postfix/send/html
```

### request body
``` js
{
  from: {
    name: "System",
    email: "system@calicomoomoo.com"
  },
  to: [
    {
      name: "User1",
      email: "user1@mail.com"
    },
    {
      name: "User2",
      email: "user2@mail.com"
    },
    ...
  ],
  subject: "Mail Subject",
  body: "mail body..."
}
```
