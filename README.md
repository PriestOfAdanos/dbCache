# dbCache

This repo uses golang-lru lib to provide caching to a dummy "database". By default it will run 10000 requests, but 9900 is just re-requests of previously chached data. If run, You should see the output along the lines of "100 retreived from  sql, 9900 from cache"

Methods:
 - fetchUserById: fetches user first from cache and if it fails than from databse
 - postgresql - databse that returns user name and id if given id (it is always user named Adam with different id's tho)
 
 Requirements:
  - go version go1.19.3 
  - fmt 
  - golang-lru 
