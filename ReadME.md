# Get started
- Run `docker-compose up`

# Running APIs
- I tend to use httpie for quick testing of apis. https://httpie.io/

From your command line:
```
http POST localhost:8000/swipe
http POST localhost:8000/user/create
http POST localhost:8000/user/login
http localhost:8000/profiles
```

# Other useful commands
To force the DB to reset -> `docker-compose up --force-recreate db`
