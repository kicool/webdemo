#webdemo
See [nick ye's blog](http://www.cnblogs.com/yjf512/archive/2012/09/03/2668384.html)

##Requirment
- MySQL
- mymysql 
 - $go get github.com/ziutek/mymysql/thrsafe
 - $go get github.com/ziutek/mymysql/autorc
 - $go get github.com/ziutek/mymysql/godrv

##Build
- $cd {webdemo_root}
- $go build -o webdemo src/webdemo/*.go 

##目录
- README.md
- etc	        //create SQL table
- src	       	//go codes
- template      //static files: js, css, html template
