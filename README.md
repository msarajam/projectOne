# ProjectOne

This is for Symantec Assignment

this program will make http call to the [server](https://gist.githubusercontent.com/joelbirchler/66cf8045fcbb6515557347c05d789b4a/raw/9a196385b44d4288431eef74896c0512bad3defe/exoplanets) , with the Get request , the application  will receive a JSON format data

after analizing and organazing that JSON file, it will do three task :

* Getting The number of orphan planets (no star).

* The name (planet identifier) of the planet orbiting the hottest star.

* A timeline of the number of planets discovered per year grouped by size.


### Install Go in MAC and run the project :
download golang from :
https://golang.org/dl/

setup the enviorment :

```export PATH=$PATH:/usr/local/go/bin ```

run this command in mac to create necessary folders :
```
mkdir go
mkdir go/bin go/pkg go/src
```

then go to ```~/go/src```

clone the project

https://github.com/msarajam/projectOne.git

go to ```projectOne``` directory

and then Run :
```go run ./...```


### Or the easy way for windows and linux :

clone the repo and :

run ```projectOne.exe``` for windows

run ```./projectOne``` for linux


### output :
```
2019/10/18 09:34:20 Number of orphan planets (no star) :  2
2019/10/18 09:34:20 The planet Identifier of the planet orbiting the hottest star is 'V391 Peg b'
2019/10/18 09:34:20  in 1781 we discovered 1 small planets, 0 medium planets, and 0 large planets. 
2019/10/18 09:34:20  in 2004 we discovered 2 small planets, 5 medium planets, and 0 large planets. 
2019/10/18 09:34:20  in 2014 we discovered 830 small planets, 30 medium planets, and 0 large planets. 
2019/10/18 09:34:20  in 1846 we discovered 1 small planets, 0 medium planets, and 0 large planets. 
2019/10/18 09:34:20  in 2001 we discovered 1 small planets, 0 medium planets, and 0 large planets. 
2019/10/18 09:34:20  in 2005 we discovered 1 small planets, 3 medium planets, and 0 large planets. 
2019/10/18 09:34:20  in 2009 we discovered 4 small planets, 6 medium planets, and 0 large planets. 
2019/10/18 09:34:20  in 2012 we discovered 52 small planets, 21 medium planets, and 0 large planets. 
2019/10/18 09:34:20  in 2015 we discovered 104 small planets, 30 medium planets, and 0 large planets. 
2019/10/18 09:34:20  in 1930 we discovered 1 small planets, 0 medium planets, and 0 large planets. 
2019/10/18 09:34:20  in 2002 we discovered 0 small planets, 1 medium planets, and 0 large planets. 
2019/10/18 09:34:20  in 2008 we discovered 1 small planets, 21 medium planets, and 1 large planets. 
2019/10/18 09:34:20  in 2010 we discovered 15 small planets, 39 medium planets, and 0 large planets. 
2019/10/18 09:34:20  in 2011 we discovered 32 small planets, 48 medium planets, and 1 large planets. 
2019/10/18 09:34:20  in 2016 we discovered 1267 small planets, 26 medium planets, and 0 large planets. 
2019/10/18 09:34:20  in 1999 we discovered 0 small planets, 1 medium planets, and 0 large planets. 
2019/10/18 09:34:20  in 2006 we discovered 1 small planets, 6 medium planets, and 0 large planets. 
2019/10/18 09:34:20  in 2007 we discovered 4 small planets, 16 medium planets, and 0 large planets. 
2019/10/18 09:34:20  in 2013 we discovered 58 small planets, 30 medium planets, and 2 large planets. 
```
