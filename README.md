# Takehome 
The entry into the program is in cmd folder. To run the program from the home directory run: 
 ```go run cmd/main.go``` 

The program is hooked up to listen on port 8080, so the following endpoints are:

```
http://localhost:8080/receipts/process

http://localhost:8080/receipts/:id/points
```


## Implementation
To make things simpler for the API, I made some assumptions for the API: 

1. ```purchaseDate``` is required to be in the format of ```2024-01-01```
2. ```purchaseTime``` is required to be in format of ```13:01```
3. ```total``` and ```price``` are in the float format

*The total and price do not account for float overflow and floating point precision, so points calculation will be off in those scenarios*

If the above format fails in the payload, there will be an BadRequest error.

To store receipts, I made a hashmap to account for fast retrievals. The "id" for the receipt is the current length of the hashmap.
E.g.
If you create a receipt after starting the application, it's ID will be 0. The next receipt created will have id 1, etc.


