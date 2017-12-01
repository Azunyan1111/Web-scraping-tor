# Web-scraping-tor

deploy branch is can you use.

master branch is develop.

### setup 
Install
```
brew install phantomjs
export MAX_DRIVER=10        # browser count
```

rpi phantomjs mode
```
sudo apt-get install libfontconfig1 libfreetype6 libpng12-0
curl -o /tmp/phantomjs -sSL https://github.com/fg2it/phantomjs-on-raspberry/releases/download/v2.1.1-wheezy-jessie/phantomjs
sudo mv /tmp/phantomjs /usr/local/bin/phantomjs
sudo chmod a+x /usr/local/bin/phantomjs
```

tor
```
sudo apt-get install tor
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
make stop # stop all tor service
```

### high speed access!!

![2017-11-29 18 32 26](https://user-images.githubusercontent.com/13769176/33369080-dd8b0db2-d536-11e7-8ad2-c8753afc62e3.png)
