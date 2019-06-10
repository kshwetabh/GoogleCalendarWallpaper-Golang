
# Sample
![sample](https://raw.githubusercontent.com/kshwetabh/GoogleCalendarWallpaper-Golang/master/output/wallpaper.png)

# Install Dependencies:
    - Get the Google Calendar API Go client library and OAuth2 package using the following commands:
        go get -u google.golang.org/api/calendar/v3
        go get -u golang.org/x/oauth2/google
    - Image manipulation library
        go get github.com/fogleman/gg

# Configuration Details
    "SourceImageName": "canyon.jpg",
    "OutputFileName": "wallpaper.png",
    "MarginRight": 400,
    "MarginTop": 20,
    "Width": 1920,
    "Height": 981,
    "PrintDate": true,
    "TitleFont": "./Fonts/Roboto/Roboto-Bold.ttf",
    "TitleFontSize": 18,
    "ItemFontSize": 14,
    "ItemFont": "./Fonts/Roboto/Roboto-Medium.ttf",
    "TitleText": "Today's Calendar",
    "ItemPadding": 25,
    "GoogleCalendarID": "aaaa@group.calendar.google.com"

# Canyon wallpaper taken from:
[pexels.com](https://www.pexels.com/photo/canyon-arizona-1672813/)