# git-project-showcaser

Showcase your projects from Github/Gitlab in your own website automatically

## Configuration

The configuration file should be in this format (yaml):
```
ResourcesPath: resources
Port: 8080
Username: rosenpin
AuthCode: YOUR_AUTH_CODE
HTTPRequestTimeout: 10s
MaxProjects: 9
IncludeForks: false
SortMode: stars
GitPlatform: github
ProfileURL: http://www.github.com/rosenpin 
```
