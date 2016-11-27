[![Teamcity board]()](https://www.youtube.com/watch?v=STNvU7Z1z5c)

Teamcity dashboard is a fun project that allows you to monitoring Teamcity builds in your command line. It use [GO](https://github.com/golang/go/) language and [termui](https://github.com/gizak/termui) library for building ui interface in console.

## Demo
Look demo [video](https://www.youtube.com/watch?v=STNvU7Z1z5c) to understand how Teamcity board is works.

## Usage
You should download binary file for [Mac OS](hi) or [Linux](). If there is't version for your platform, you may compile it from source code manually. After that use `teamcity-board` command with that environment variables:

| Param                      |      Description                                       | Type    |  Required | Default Value |
|:---------------------------|:-------------------------------------------------------|:--------|:----------|:--------------|
| TEAMCITY\_HOST             | Teamcity host (For example http://teamcity.ru)         | String  | yes       | no            |
| TEAMCITY\_PORT             | Teamcity port                                          | Int     | no        | 8111          |
| TEAMCITY\_PROJECT\_ID      | Project id which you want to monitor                   | String  | yes       | no            |
| TEAMCITY\_AUTH\_HEADER     | Basic auth header for base64(login:password)           | String  | yes       | no            |
| TEAMCITY\_UPDATE\_INTERVAL | Interval in seconds between data update                | Int     | no        | 15            |

You may add this params to starting command `TEAMCITY_HOST="http://teamcity" TEAMCITY_AUTH_HEADER="Basic <token>" TEAMCITY_PROJECT_ID="AndroidProjects_AvitoPro" ./teamcity-board` or set it in `~/.bashrc` file.

## Warning

It's not secure to storing basic auth information in your environment variables. This token contains your teamcity login:password "encrypted" using Base64 algorithm that easy to "decrypt".
