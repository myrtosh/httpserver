# httpserver
Type make in the terminal
This http service has three endpoints
-->First endpoint: go to the browser and type 'http://localhost:8080/helloworld'
this returns: Hello Stranger!
-->Second endpoint:go to the browser and type 'http://localhost:8080/helloworld?name=WhateverNameYouWant'
this returns:Whatever Name You Want //(split by camelcase)
-->Third endpoint: go to the browser and type 'http://localhost:8080/versionz'
this returns: {"GitHash":"lastcommtithash","NameOfTheProject":"httpservice"}