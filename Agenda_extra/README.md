# DB
    This time I implement simple DB operation by using "gorm"

## API
### AddUser
(post)http://localhost:8080/service/userinfo?username=1&password=2



	{"OK": true,"Message": "add user"}

### GetAUser

(Get)http://localhost:8080/service/userinfo?userid=1



	{"OK": true,"Message": {"UID":1, "Username": "1", "Password": "2", "CreateTime": "2017-11-30T20:52:30+08:00"}}

### GetAllUser

(Get)http://localhost:8080/service/getallusers
	
	{
		"OK":true,
		"Data":[{
			"ID":1,
			"Username":"1",
			"Password":"2",
			"SignUpDate":"2017-11-30T20:52:30+08:00"
			},
			{"ID":2,
			"Username":"zbc",
			"Password":"zbc",
			"SignUpDate":"2017-11-30T20:54:48+08:00"
			}]
	}

#Bye 