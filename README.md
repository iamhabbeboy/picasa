# wallpaper

Introducing the Random Wallpaper Tool for Mac OSX and Linux *(coming soon)*

Are you tired of staring at the same old wallpaper on your Mac every day? Do you want to add some variety and spice up your desktop? Look no further than our Random Wallpaper Tool!

A simple CLI tool that downloads random pictures from [unsplash](unsplash.com/) and use it as a wallpaper. 

```
> wallpaper set --interval 5
```

- 5m - 5 minutes
- 10m - 10 minutes
- 60m - 60 minutes
- 1h  - 1 hour
- 1d - 1 day

### Configuration
Due to unsplash policy, the secret key used for this project might stop working at some point, to configure yours; Just head to unsplash and create an app, get the secret key and access point.

```
> wallpaper config --access_token xklkejwerlkjk --secret_key lwkejrklwjerlkjwelkrj
```


You can set time interval for wallpaper change
