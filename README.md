# Web-scraping-tor


### setup 
Install
```
brew install chromedriver
export MAX_DRIVER=10        # browser count
```

setup
```.sh
make setup
```

### start

```.sh
make start-tor  # start all tor service
make start      # start web scraping
```

### stop

```.sh
make stop-tor # stop all tor service
```