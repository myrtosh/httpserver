HTTP service with three HTTP endpoints
============
### step by step
1. clone this repository: ```git clone https://github.com/myrtosh/httpservice```
2. run: ```cd httpservice```
3. run: ```make docker```

------
### The endpoints:
1. point a browser at http://localhost:8080/helloworld which returns Hello Stranger!
2. point a browser at http://localhost:8080/helloworld?name=AlfredENeumann (or any filtered value) which returns Hello Alfred E Neumann
3. point a browser at http://localhost:8080/versionz which returns a JSON with git hash and name of the project