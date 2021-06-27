# fun_size_url

A tiny url clone using go and gorilla mux.

The frontend is written in html, javascript and css, and then served via the golang api route.

The url hash redirects are stored in a cassandra database setup on startup using docker compose.
The go web app then waits for this process to be completed before starting up so may take a while the first time running.

Everything can be run using the command: 

```docker-compose up```

Potential future improvements:
- Add user database table so the program doesn't have to search through all rows everytime someone redirects
- Add testing (this started as a project to learn docker with multiple containers so very much tested along the way but TDD would be better)
- Handle incorrect URLs and hashes better (at the moment they just reroute to the root)

![Fun Size URL](https://github.com/rhodriguerrier/fun_size_url/blob/main/fun_size_url_example.PNG?raw=true)
