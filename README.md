# Release monitor

**This package is not completed**

Simple commandline app for watching releases a running tasks when is a new release published.

## Motivation

From time to time I need to be notificated when some library/package/apps is updated to a new version. For example, I manage [ARM build of docker/registry](https://github.com/Budry/docker-registry-arm) and I need hook for run automatic build when official app is updated to a new version. Because GitHub doest not have any simple mechanism for that (has only atom feed but without any notifications) I create this simple app.

## How to use it

The best way how to use this app is [as docker image](https://hub.docker.com/r/budry/release-monitor/).

```shell
docker run \
    -v /path/to/config/directory:/etc/release-monitor \
    -v /path/to/data/directory:/var/lib/release-monitor \
    budry/release-monitor
```

Required is only the first volume, application needs to have configuration file. The second volume is used for store last process release notification. When you don't use it and restart container application, resend all notification from beginning.

### Configuration file

Configuration file must have name `config.json` and must contain valid JSON. File cannot be empty, minimal content is `{}`

#### Example configuration

```json
{
    "interval": "@daily",
    "monitors": [
        {
            "name": "Docker registry",
            "url": "https://github.com/Budry/docker-registry-arm",
            "adapter": "github",
            "commands": [
                "curl --header \"Content-Type: application/json\" --request POST --data %%RELEASE%% <CI server endpoint>"
            ]
        }
    ]
}
```

This example send HTTP request with release info in json format to my CI server and run automatic build. 

#### Format

* `root`
    * `interval` - CRON string. You can use format described [here](https://godoc.org/github.com/robfig/cron)
    * `monitors` - list of monitors
* `monitor`
    * `name` - Your monitor name. Name is used for remember last processed release
    * `url` - Source URL address for feed
    * `adapter` - Which adapter should be used for parse
    * `commands` - list of strings. Each command will be run and you can use %%RELEASE%% placeholder for insert release info. All commands are execute for each new release

#### Release info

Each release has simple format

```json
{
    "version": "",
    "date": "" 
}
```

## TODO

- [ ] Custom dokcer images support
- [ ] Adapter for Bitbucket and GitLab
- [ ] Tests