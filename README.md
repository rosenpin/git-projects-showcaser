# git-project-showcaser

Showcase your projects from Github/Gitlab in your own website automatically

## Configuration

The configuration file should be in this format (yaml):
```
ResourcesPath: resources # Path to the resources directory, should include the index.html file to work with
Port: 8080 # Port to listen on
Username: rosenpin # The username for the service
AuthCode: YOUR_AUTH_CODE # Optional auth code for services like github that limit request number
HTTPRequestTimeout: 10s # API request time out
MaxProjects: 9 # The maximum number of projects to show
IncludeForks: false # Include fork projects or not
SortMode: stars # The method to sort the projects by
GitPlatform: github  # The git service to use
ProfileURL: http://www.github.com/rosenpin # The url to the profile if you want the visitors to be able to go to your profile
ReloadInterval: 12h # At what interval to fetch from git
```

## Running it

To run the program:
``` ./showcase-projects -c $CONFIG_PATH```
