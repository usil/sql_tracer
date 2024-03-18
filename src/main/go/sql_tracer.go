package main
import (
    "fmt"
    _ "github.com/sijms/go-ora/v2"
    "database/sql"
    "flag"
    "strconv"
    "os"
    "time"
    "strings"
)
func main() {

    port := flag.Int("p", 1521, "port")
    host := flag.String("h", "localhost", "host")
    user := flag.String("u", "foo", "user")
    //@TODO: password is moved temporary to env 
    //password := flag.String("ps", "changeme", "password")
    password := os.Getenv("sql_tracer_database_password")
    databaseName := flag.String("n", "acme", "service name")
    alias := flag.String("a", "acme_db", "a human alias to the host from where is running")
    reportLocation := flag.String("r", "/tmp/report.csv", "a csv f-null location to append the metrics")
    intervalInSeconds := flag.Int("i", 15, "interval of execution in seconds")

    //@TODO: get ip
    // localIp,localIpError  := getLocalIp()
    // if localIpError == nil {
    //     fmt.Println(localIpError)
    // }
    localIp := "192.168.0.1"

    flag.Parse()
    fmt.Println("ip", localIp)
    fmt.Println("alias", *alias)
    fmt.Println("host", *host)
    fmt.Println("port", *port)
    fmt.Println("user", *user)
    fmt.Println("password", "****")
    fmt.Println("database name", *databaseName)
    fmt.Println("reportLocation", *reportLocation)

    connectionStringForLog := fmt.Sprintf("oracle://%s:%s@%s:%s/%s", *user, "*****", *host, strconv.Itoa(*port), *databaseName)
    fmt.Println("connection string", connectionStringForLog)

    connectionString := fmt.Sprintf("oracle://%s:%s@%s:%s/%s", *user, password, *host,  strconv.Itoa(*port), *databaseName)   
    
    for range time.Tick(time.Second * time.Duration(*intervalInSeconds)) {
        start := time.Now()

        conn, errCon := sql.Open("oracle", connectionString)
        if errCon != nil {
            fmt.Println(errCon)
            saveReport(localIp, *alias, start, time.Now(), "failed", "error code 6660: "+strings.ReplaceAll(errCon.Error(), ",", ";"), *reportLocation)
            conn.Close()
            continue
        }
        errCon = conn.Ping()
        if errCon != nil {
            fmt.Println(errCon)
            saveReport(localIp, *alias, start, time.Now(), "failed", "error code 6661: "+strings.ReplaceAll(errCon.Error(), ",", ";"), *reportLocation)
            conn.Close()
            continue
        }
        fmt.Println("\nconnection status: success")
        rows, errQuery := conn.Query("SELECT 1 from dual")
        rows.Close()
        conn.Close()
        if errQuery != nil {
            fmt.Println(errQuery)
            saveReport(localIp, *alias, start, time.Now(), "failed", "error code 6662: "+strings.ReplaceAll(errQuery.Error(), ",", ";"), *reportLocation)
            continue
        }
        var (
            id int64
        )
        for rows.Next() {
            rows.Scan(&id)
        }
        fmt.Println("query validation result:", id)
        saveReport(localIp, *alias, start, time.Now(), "success", "", *reportLocation)
    }
}

//error string should not have the [,] char because [,] is the default column delimiter in csv
func saveReport(ip string, alias string, start time.Time, end time.Time, status string, errAsString string, reportLocation string){

    b, err := os.ReadFile(reportLocation)
    if err != nil {
        fmt.Println("report doesn't exist")
        str := "client_ip, client_alias, start, end, status, elapsed_millis, error_trace\n"
        //if not exist, initialize it with empty string
        err = os.WriteFile(reportLocation, []byte(str), 0644)
        if err != nil {
            fmt.Println(err)
        }
    }
    fmt.Println("previous file size in bytes:", len(b))

    f, err := os.OpenFile(reportLocation, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        fmt.Println("report doesn't exist , error code 66650")
        os.Exit(1)
    }
    
    defer f.Close()
    
    //append new line
    startMillis := start.UnixNano()/ int64(time.Millisecond)
    endMillis := end.UnixNano()/ int64(time.Millisecond)
    diff := endMillis - startMillis
    newLine := fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s\n",
        ip, alias, start.Format(time.RFC3339), end.Format(time.RFC3339), status, strconv.Itoa(int(diff)), errAsString)
    fmt.Println("new line", newLine)
    if _, err = f.WriteString(newLine); err != nil {    
        fmt.Println(err)
        os.Exit(1)
    }
}

//#TODO
// func getLocalIp(){
    // conn, err := net.Dial("udp", "8.8.8.8:80")
    // if err == nil {
    //     fmt.Println("Failed to get the local ip")
    //     localAddr := "unknown"
    // }else{
    //     defer conn.Close()
    //     localAddr := conn.LocalAddr().(*net.UDPAddr)    
    // }
    // return localAddr    
// }


