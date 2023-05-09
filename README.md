# Picasa for Unix - rwallpaper

Introducing the Random Wallpaper Tool for Mac OSX and Linux *(coming soon)*

Are you tired of staring at the same wallpaper on your Mac every day? Do you want to add some variety and spice up your desktop? Look no further than Random Wallpaper!

A simple CLI tool that downloads random pictures from [unsplash](unsplash.com/) and use it as a wallpaper. 

> NB: The random change of the wallpaper rely solely on cronjob, which means the tool will request access to edit your crontab.

Download release and copy to `/usr/local/bin/` folder so it can be available globally using

```
 > cp picasa /usr/local/bin
```

```
> picasa -h
```

```
> picasa set interval 5m
```

```
 > picasa set -i 24h
```

- 5m - 5 minutes
- 10m - 10 minutes
- 60m - 60 minutes
- 1h  - 1 hour
- 1d - 1 day

### Configuration
Due to unsplash policy, the secret key used for this project might stop working at some point, to configure yours; Just head to unsplash and create an app, get the secret key and access point.

```
> picasa config --access_token xklkejwerlkjk --secret_key lwkejrklwjerlkjwelkrj
```

Change image query
The default image query type is `nature` which means that you only get nature related images, but you can decide to change the image query with:
```
> picasa config -q water
```
You will only see the effect after the cronjob execute to download next set of images.
