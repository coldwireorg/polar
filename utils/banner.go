package utils

import "fmt"

func Banner(version string) {
	b := fmt.Sprintf(`    ___      _            
   / _ \___ | | __ _ _ __ 
  / /_)/ _ \| |/ _  | '__|
 / ___/ (_) | | (_| | |   
 \/    \___/|_|\__,_|_|   

 > version %s
 `, version)

	fmt.Println(b)
}
