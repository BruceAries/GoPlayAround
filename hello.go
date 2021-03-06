package main 

import ( 
	"fmt"
//	"github.com/bruceAries/stringutil"
	"net"
	"os"
	"log"
)

func main() {
//  fmt.Println( stringutil.Reverse( "Hello, Go" ) )
  
  //Network related trial
  addr := "baidu.com:80"
  conn, err := net.Dial( "tcp", addr )
  if err != nil {
	fmt.Println( "Connect to %q failed.", addr )
  }
  
  fmt.Fprintf( conn, "GET / HTTP/1.0\r\n\r\n" )


  //OS related trial
  {
	f, err := os.Open( "hello.go" )
	if err != nil {
		fmt.Println( err )

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal( err )
		}	else {
			fmt.Println( "rooted path: ", dir )
		}
		log.Fatal(1)
	}
	
	
	d, err := f.Stat()
	if err != nil {
		f.Close()
		log.Fatal( err )
	}
	
	fmt.Println( d.Name(), d.Size(), d.ModTime() )
  }
  

//this is from github remote server
//this is from remoete again
//this is from local mac
//this is from remoete again
//this is from remoete again
// this is from local mac 1
