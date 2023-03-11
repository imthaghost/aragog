# Aragog API

## Functions

### Management
| Function | Request | Response |
|-------------------------------------|:------:|------------------------------------------------------------------------|
| [InviteUser](#) | UserReq | InviteResp |
| [RemoveUser](#) | UserReq | RemoveResp |

### Health

| Function | Request | Response |
|-------------------------------------|:------:|------------------------------------------------------------------------|
| [HealthCheck](#health-check) | HealthReq | HealthResp


## Examples

### Health Check
```go
client := pb.NewAragogProtobufClient("http://localhost:8080", &http.Client{})
// health check
health, err := client.HealthCheck(context.Background(), &pb.HealthReq{})
if err == nil {
log.Println(health.Status)
}
```
#### Objects
```protobuf
// HealthReq is an empty body request but represents
// root health check message
message HealthReq {}
```
### Health Response
```protobuf
// HealthResp represents the root response message for a health check
message  HealthResp {
  
  int32 status = 1;
  
}
```
Responses:
- 200 OK
- 500 InternalServerError: error message
