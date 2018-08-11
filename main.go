package main

//import "github.com/gin-gonic/gin"
import _ "github.com/go-sql-driver/mysql"
import "database/sql"
import "fmt"

func main() {
  fmt.Println("hello world")

  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sampledb")
  checkErr(err)

	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	checkErr(err)

	stmt, err := db.Prepare("CREATE TABLE person (id int NOT NULL AUTO_INCREMENT, first_name varchar(40), last_name varchar(40), PRIMARY KEY (id));")
	checkErr(err)

	_, err = stmt.Exec()
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("Table successfully migrated....")
	}

  //router := SetupRouter()
  //router.Run(":8081")


}

func checkErr(err error) {
  if err != nil {
    fmt.Print(err.Error())
  }
}

// func SetupRouter() *gin.Engine {
//   router := gin.Default()
//
//   v1 := router.Group("api/v1")
//   {
//     v1.GET("/instructions", ins.getInstructions())
//   }
//   return router
// }
