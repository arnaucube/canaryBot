# canaryBot
Bot to check if services are alive.
Current bots:
- Matrix (Riot)


#### Config
File `config.json`
```json
{
  "matrix": {
    "room_id": "!zzzzz:mmmmmm.ooo",
    "user": "aaaaa",
    "password": "xxxxx",
    "server": "https://sssss.ooo"
  },
  "services": [{
      "name": "name01",
      "url": "http://127.0.0.1:80",
      "statusCode": 200
    },
    {
      "name": "service02",
      "url": "http://127.0.0.1:7000/api",
      "statusCode": 200
    }
  ],
  "sleepTime": 30,
  "retry": 5
}
```

- sleepTime: time between each request to the services
- retry: after X times failing, restart the counter

### Run
```
./canaryBot
```
