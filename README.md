# cpfc-dbservice



cpfcdbservice provides access to Crystal Palace FC results through a simple RESTful Web Services. 

  

## services 


1)  GET/results

    returns a json  of all results
    
    to test, run 
    
```     curl -i http://localhost:8000/results
```  

2)  GET/results:id

    returns a json object of results of game with provided id
    
    to test, run 
    
```     curl -i http://localhost:8000/results:id
```  

3)  POST/results

    given a json object, the POST service writes into into the db
    
    to test, run 
    
```     curl -i -X POST -H "Content-Type: application/json" -d "{\"Season\":\"1945/46\",\"Round\":\"15\",\"Date\":\"10-09-1946\",\"Kickofftime\":\"13:00\",
        \"AwayorHome\":\"A\",\"Oppenent\":\"Arsenal\",\"Resultshalftime\":\"1:2\",\"Resultsfulltime\":\"2:2\"}" http://localhost:8000/result
``` 


## docker container

```docker pull crocusvalley/cpfc-dbservice
```