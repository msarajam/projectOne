# ProjectOne

This is for Symantec Assignment

this program will make http call to the [server](https://gist.githubusercontent.com/joelbirchler/66cf8045fcbb6515557347c05d789b4a/raw/9a196385b44d4288431eef74896c0512bad3defe/exoplanets) , with the Get request , the application  will receive a JSON format data

after analizing and organazing that JSON file, it will do three task :


..* Getting The number of orphan planets (no star).

..* The name (planet identifier) of the planet orbiting the hottest star.

..* A timeline of the number of planets discovered per year grouped by size.




### Install Go in MAC :
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


### Or the easy way :

clone the repo and :

run ```projectOne.exe``` for windows

run ```./projectOne``` for linux
