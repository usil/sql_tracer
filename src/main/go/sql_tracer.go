package main
import (
    "fmt"
    _ "github.com/sijms/go-ora/v2"
    "database/sql"
    "flag"
    "strconv"
    "os"
)
func main() {

    fmt.Println("Parsing shell args")

    port := flag.Int("p", 1521, "port")
    host := flag.String("h", "localhost", "host")
    user := flag.String("u", "foo", "user")
    password := flag.String("ps", "changeme", "password")
    serviceName := flag.String("s", "acme", "service name")

    flag.Parse()
    fmt.Println("host", *host)
    fmt.Println("port", *port)
    fmt.Println("user", *user)
    fmt.Println("password", "****")
    fmt.Println("service name", *serviceName)

    _connectionString := fmt.Sprintf("oracle://%s:%s@%s:%s/%s", *user, *password, *host,  "*****", *serviceName)
    connectionString := fmt.Sprintf("oracle://%s:%s@%s:%s/%s", *user, *password, *host,  strconv.Itoa(*port), *serviceName)
    fmt.Println("connectionString", _connectionString)
    
    conn, errCon := sql.Open("oracle", connectionString)
    if errCon != nil {
        fmt.Println(errCon)
        os.Exit(1)
    }
    errCon = conn.Ping()
    if errCon != nil {
        fmt.Println(errCon)
        os.Exit(1)
    }
    fmt.Println("connected")
    rows, errQuery := conn.Query("SELECT 1 from dual")
    defer rows.Close()
    if errQuery != nil {
        fmt.Println(errQuery)
    }
    var (
        id int64
    )
    for rows.Next() {
        rows.Scan(&id)
    }
    fmt.Println("query validation result:", id)

}



