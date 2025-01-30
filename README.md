# Api Rate Limiter Go
This project include a api rate limiter using Golang. The responsibility of this project is listening on api calls and check it limit reach or not, if a config limit reach, user requests should deny until limit remove if not, should accept user request and update user history.

## Components
1. Config: you can define a config with below schema, and this config will apply over all api calls.
   ```json 
   {
    "pattern": String,
    "limit": Int,
    "duration": Int,
   }
2. 