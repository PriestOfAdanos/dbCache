# dbCache

This repo uses golang-lru lib to provide caching to a dummy "database".

Methods:
 - fetchUserById: fetches user first from cache and if it fails than from databse
 - postgresql - databse that returns user name and id if given id (it is always user named Adam with different id's tho)
 
 Requirements:
  - go version go1.19.3 
  - fmt 
  - golang-lru 
