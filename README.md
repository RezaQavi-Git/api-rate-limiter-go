# Api Rate Limiter Go
This project include an api rate limiter using Golang. The responsibility of this project is listening on api calls and check it limit reach or not, if a config limit reach, user requests should deny until limit remove if not, should accept user request and update user history.

## Rate Limit Mechanisms
### Fixed Window Counter
This algorithm divides time into fixed intervals (e.g., 1 minute, 1 hour). 
It counts the number of user requests within each interval. If the count exceeds the predefined limit within the interval, subsequent requests are blocked until the next interval begins, at which point the counter resets.
- Boundary Problem
### Sliding Window Counter
This approach improves upon the fixed window by dividing time into smaller sub-intervals (e.g., seconds within a minute). 
For each incoming request, it calculates the sum of requests in the current sub-interval and the proportionate requests from the previous sub-interval, creating a "sliding" effect. If the total exceeds the limit, the request is blocked.
- Multiple Redis operations
### Sliding Log
This method maintains a log of timestamps for each user's requests. 
Upon receiving a new request, the system removes outdated timestamps (those outside the defined time window) and checks the count of remaining timestamps. 
If the count exceeds the limit, the request is blocked; otherwise, it's allowed, and the new timestamp is added to the log.
- High memory usage
- Slower performance
### Token Bucket
In this mechanism, each user is assigned a bucket with a fixed capacity of tokens (representing the request limit). Tokens are added to the bucket at a constant rate. Each incoming request consumes a token. If a request arrives when the bucket is empty (no tokens available), it is blocked.
- If refill rate is too high, users may still spam the system.
### Leakey Bucket
Similar to the token bucket, but with a key difference: requests are added to a queue (the bucket), and they are processed at a fixed rate, regardless of the arrival rate. If the queue is full, incoming requests are dropped.
- Requests may be delayed
- Queue memory usage

## Components
1. `RateLimiter`: Base component of application, include `HttpMiddleware`, get incoming request, extract data, check limit and deny or accept request.
2. `RateLimitConfig`: you can define a config with below schema, and this config will apply over all api calls.
   ```yaml
   rate_limit_configs:
   - pattern: "^/api/v2/.*"
     limit: 2
     duration: 60
   - pattern: "^/api/v3/.*"
     limit: 1
     duration: 60

## How to Use
To use this minimal api rate limiter, you need to set up a redis client and also pass `RateLimitConfig` list to `NewRateLimiter` method, and then you can use middleware anywhere in code.
```go
appConfig := configs.Load()
rc := tools.NewRedisClient(appConfig.RedisConfigs)

rateLimiter := internal.NewRateLimiter(appConfig.RateLimitConfigs, rc)
```
more detail on [main.go](./main.go)