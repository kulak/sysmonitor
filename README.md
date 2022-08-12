# Readme

This application generates HTML report files that can be viewed over webbrowser to see the status of the remote system.

Report file name is based on the day of the year and can be exposed using web server like `nginx`.

The application can be scheduled to run every day using `systemd`.

## Report Example

![report](doc/screenshot-1.png)

Nginx index of reports served by nginx ![index of reports](doc/screenshot-2.png)

## Configuration Example

See example [sysmonitor.yaml](doc/sysmonitor.yaml) file.

## Commands

### Build

    go build .
