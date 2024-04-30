### create new account
___

```curl -X POST http://localhost/newAccount -d '{"name":"baby","phone":"6151234567","email":"user1@gmail.com","psword":"secret"}'```

🤫 to create an admin user, add the admin role

```curl -X POST http://localhost/newAccount -d '{"name":"someAdminWorker","phone":"2222222222","email":"user22@gmail.com","psword":"secret","role":"admin"}'```

### get all users
___
```curl -X GET http://localhost:8080/getAllUsers```

### delete user
___

📝 you will need to get the user's UUID before deleting the desired user

```curl -X POST http://localhost/deleteUser -d '{"uuid":"57dcd3f0-4e16-4776-9e5b-da564ca9196c"}'```