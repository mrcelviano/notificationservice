# notificationservice
SocialTech

Сервис отвечает за уведомления пользователей о регистрации

Сервис слушает порт 8081

## GRPC

#### SendNotification
Request structure
```
message User {
    int64 ID = 1;
    string Email = 2;
    string Name = 3;
}
```

Response structure
```
message SendNotificationResponse {
  int64 TaskID = 1;
}
```
