## Agenda Command Design

### User

##### Register

* agenda register --user=username --pass=password --email=123456@qq.com  --mobilephone=15687878122
* agenda register -u username -p password -e 123456@qq.com  -m 15687878122

##### Login

* agenda login --username=username --pass=password
* agenda login -u username -p password
* agenda logout

##### Query

* agenda queryu --username=username(default is all registed user)
* agenda queryu -u username

##### Delete

* agenda delu

  ​

#### Meeting

##### Create Meeting

* agenda cm --title=something --participator=alice,job,gg  --stime=2017-06-30 12:00 --etime=2017-06-30 12:30
* agenda cm -t something -p alice,job,gg  -s 2017-06-30/12:00 -e 2017-06-30/12:30

##### Add participator to specific meeting according to the title

* agenda addp --title=something --participator=alice,job
* agenda addp -t something -p alice,job
* agenda delp --title=something --participator=alice,job
* agenda delp -t something -p alice,job

##### Query meeting in specific time(User's meeting)

* agenda querym --stime=2017-06-30/12:00 --etime=2017-06-30/12:30
* agenda querym -s 2017-06-30/12:00 -e 2017-06-30/12:30

##### Delete Meeting(the owner of the meeting)

* agenda delm --title=something
* agenda delm -t something

##### Quit Meeting(the participator of the meeting)

* agenda quitm --title=something
* agenda quitm -t something

##### Clear Meeting(the owner of the meeting)

* agenda clearm