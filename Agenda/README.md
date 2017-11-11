# Agenda
Golang_Agenda

#### How To Use

​	You can download agenda.exe to you PC. Then go into the path that include the agenda.exe and input something like agenda register ... (See more details in cmd-design.cmd)

#### Test

* agenda register -u pansir -p pansir -e 3791879@qq.com -m 1568983980 ![register.jpg](/image/register.jpg) 

  We can see in User.json

  ![register1](/image/register1.jpg)

  ​

* agenda login -u pansir -p pansir ![login](/image/login.jpg)

  We can see in curUser.txt ![login1.jpg](/image/login1.jpg)

* agenda logout ![logout](/image/logout.jpg)

  We can see in curUser.txt ![logout](/image/logout1.jpg)

* agenda queryu ![queryu](/image/queryu.jpg)

*  agenda cm -t golang -p weimumu123,weimumu -s 2017-09-09/21:00 -e 2017-09-09/21:33 ![cm](/image/cm.jpg)

  We can see in Meeting.json ![cm](/image/cm1.jpg)

* agenda addp -t golang -p zhengpin,xiejieqi ![addp](/image/addp.jpg)

  We can see in Meeting.json two participator were add to Meeting.json ![addp](/image/addp1.jpg)

*  agenda delp -t golang -p weimumu,zhengpin ![delp](/image/delp.jpg)

  We can see in Meeting.json two participator were deleted ![delp](image/delp1.jpg)

* agenda querym --stime=2017-09-09/21:00 --etime=2017-09-09/21:10 ![querym](image/querym.jpg) 

* agenda delm -t golang ![delm](/image/delm.jpg) 

  We can see in Meeting.json this meeting is deleted ![delm](/image/delm1.jpg)

* agenda login -u weimumu -p weimumu123 and agenda quitm -t golang ![quitm](/image/quitm.jpg)

  We can see in Meeting.json, weimumu was deleted from the meeing golang ![quitm](/image/quitm1.jpg)

* agenda clearm ![clearm](/image/clearm.jpg)

  We can see in Meeting.json ![clearm](/image/clearm1.jpg)

* agenda delu ![delu](/image/delu.jpg) 

  We can see in Meeting.json all meeting that initiated by the user are deleted ![delu](/image/delu1.jpg)

* about log. We can see in Operate.log ![log](/image/log.jpg)

  [info] means operate successfully

  [error] means operate will cause error

#### Thank!

