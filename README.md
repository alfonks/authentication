***Diagram Flow***
`deall-alfon/diagram`

***Tech Stack used***: Go, MongoDB, Redis

***To run the program***

```
1. go mod vendor && go mod tidy -v
2. docker compose up
3. go run main.go (do this in seperate tab from docker compose)
```

***Preparing the MongoDB***
```
1. Open mongodb
2. Create database with the name deall-test
3. Create a document with the name deall-user
4. Import the collection in schema/deall-user.json
5. (optional) I add an index based on the user.email field, but i cant seem to export
the existing indexes
```

***Credential Mongo DB***
```
User: ruri
Password: Q1w2E3r4T5y6
```

***Credential normal user:***
```
Email: asdf@gmail.com
Password: asdf
```

***Credentital Admin User***
```
Email: qwerty@gmail.com
Password: qwerty
```

To view documentation, it is attached in json format `deall.postman_collection.json` and you can import it to the postman apps

***What you can do?***
```
1. User Login (with the attached credential) -> also used to get pair key
2. Sign up (normal user)
3. Sign up (admin user) -> need admin role (login with admin credential)
4. Generate refresh token -> used to create a new token pair
```

***Explaintaion User Login***
```
This endpoint return a pair of access key and refresh key. The access key 
have a lifetime of 15 minutes and refresh key 365 days (you can change this life time
by modifying the config json cfg/deall-alfon-secret-config.json) 
```

***Explaination User Sign Up***
```
Nothing much here, normal sign up flow
```

***Explaination User Sign Up Admin***
```
For this flow, you need extra verification. The only one who can access this endpoinst is
another admin, so normal user can't make an admin account. To do so, hit the login endpoint
to get the access key, then place the access key into the header with the field name Jwt-Access-Key
and then you can create another admin account. The verification is done through middleware.
```

***Explaination Generate New Token***
```
In case you have an expired access key, you can get a new one by hitting this endpoint.
Add the refresh token you got when hitting the login endpoint, then put the token into the header
with field name Jwt-Refresh-Token and then you will get a new token pair.
```