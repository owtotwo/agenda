# Agenda - Go Version for Service Computing Course

## Team member

``` info
Name   Number      Email
hwh    1533**20    owtotwo@163.com

```

(The number of team members is ... just One. And no real-name in github, pardon me)

## Usage

### get help

If you want to know more about parameters of the command, you can input

> $ agenda help

or

> $ agenda help

or

> $ agenda help

for details:

``` log
Usage:
agenda [command]

Available Commands:
help        Help about any command
meeting     Manage meetings
user        Manage user account

Flags:
-h, --help   help for agenda

Use "agenda [command] --help" for more information about a command.
```

If you want to know more details about command of user or meeting, you can input:

> $ agenda user

``` log
Usage:
agenda user [flags]
agenda user [command]

Available Commands:
create      create user account
delete      Delete user account
login       user login
logout      Sign out
show        Show user account

Flags:
-h, --help   help for user

Use "agenda user [command] --help" for more information about a command.
```

> $ agenda meeting

``` log
Usage:
agenda meeting [flags]
agenda meeting [command]

Available Commands:
clear       Clear all meetings
create      create meeting
delete      Delete meeting
leave       Leave meeting
manage      Manage meeting
show        Show meeting information

Flags:
-h, --help   help for meeting

Use "agenda meeting [command] --help" for more information about a command.
```

You can know about the details of every command by inputting

> $ agenda meeting [command] -h

or

> $ agenda user [command] -h

### register

> $ agenda user create -uAT -eowtotwo@163.com -t13250334488

And input password correctly.

### log in

> $ agenda user login -uAT

And input your password.

### log out

> $ agenda user logout

### show information of all users

> $ agenda user show

``` log
@AT:
Username       E-mail                   phone number
llguser        fgh                      dfghjkl
AT             owtotwo@163.com          13250334488
a              14125514@qq.com          13900011111
ss             saaaaaaaa                sssssss
aaa            dghjkl                   dfghjk
abcadffa       5164163423@qq.com        10086

Total number is 6
```

### delete your account

> $ agenda user delete

### create meeting

> $ agenda meeting create --name meetingA

``` log
    @AT:
    1. llguser
    2. gdfghjkuser
    3. guser
    4. UUU
    5. a
    6. ss
    7. aaa
    8. AT
    Please choose the number of them to join your meeting(seprate with space): 1 2 3
    Please input start time(format: YYYY-MM-DD/HH:MM): 2018-11-11/11:22
    Please input end time(format: YYYY-MM-DD/HH:MM): 2018-11-11/12:22
    Create meeting meetingA finished.
```

### leave meeting you participate in

> $ agenda meeting leave --name meetingA

``` log
@user:
Finish leaving meeting meetingA.
```

### manage meeting created by you

#### delete one of participators in your meeting

> $ agenda meeting manage -d --name meetingA

``` log
@AT:
Participators:
1. llguser
2. gdfghjkuser
3. guser
4. UUU
5. a
6. ss
Please input the number you want to remove: 1
llguser was removed.
```

#### add someone as a partipator in your meeting

> $ agenda meeting manage --name meetingA

``` log
@AT:
You can choose some of them to add to your meeting:
1. llguser
2. gdfghjkuser
3. guser
4. UUU
5. a
6. ss
7. aaa
8. AT
Please input the number of users you want to add(separate with blank): 1
llguser was added.
```

### delete meeting created by you

> $ agenda meeting delete --name meetingA

``` log
@AT:
Delete meetingA finished.
```

### clear all meetings created by you

> $ agenda meeting clear

``` log
@AT:
Are you sure you want to clear all of your meetings? (y/n) y
All of the meeting have been deleted.
```

### show information of all meetings you sponsored or participate in

> $ agenda meeting show -s2015-11-11/11:00 -e2019-10-31/11:00

``` log
@AT:
--·--·--·--·--·--·--·--·--·--·--
Theme: meetingname
Sponsor: AT
Start time: 2018-11-11/11:22
End time: 2018-11-11/12:22
Participator: llguser, gdfghjkuser, guser
--·--·--·--·--·--·--·--·--·--·--
```

## data

All user data of our program include _user.json_, _meeting.json_, _curUser.txt_, _agenda.log_ is put in _HOME/.agenda_.

## code

My code consists by `cmd` and `entity`. `entity` is responsible for low-level storage and logical processes. `cmd` is used for processing user input. `service` is back-end supporting the `cmd`.

## IG nb！！！