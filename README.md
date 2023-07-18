# Food delivery rest service

An internship project where I created food delivery service that has CRUD for users, products, food venues and transactions.


Start project with:
```
go run main.go
```

User endpoints:
```
/user/register
Request body JSON:

{
    "name": "test",
    "email": "test@test.com",
    "password": "testtest",
    "phone": "123456789"
}

/user/login
Request body JSON:

{
    "email" : "test@test.com",
    "password": "testtest"
}

/user/delete
Request body JSON:

{
    "email": "test@test.com",
    "password": "testtest"
} 

/user/edit
Request body JSON:

{
    "name": "test",
    "email": "test@test.com",
    "password": "testtest",
    "phone": "123453212"
}

/users/get
Request body JSON:

{
    "name": "test",
    "email": "test",
    "phone": "test" 
}
```


