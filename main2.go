package main

//import ("net/http")

import (
        "database/sql"
        //"fmt"
        _ "github.com/go-sql-driver/mysql"
        "net/http" //http定義
)

//User is struct
type User struct {
        ID   int
        Name string
}

func main() { // メイン処理
    http.HandleFunc("/", HelloHandler) // ハンドラを登録 --- (*1)
    http.ListenAndServe(":8888", nil) // サーバーを起動 --- (*2)
}

// HelloHandler サーバーの処理内容を記述 --- (*3)
func HelloHandler(w http.ResponseWriter, r *http.Request) {
   // w.Write([]byte("Hello, World!")) // --- (*4)

        db, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/sample_db")
        if err != nil {
                panic(err.Error())
        }
        defer db.Close()

        rows, err := db.Query("SELECT * FROM users")
        if err != nil {
                panic(err.Error())
        }
        defer rows.Close()

        for rows.Next() {
                var user User
                err := rows.Scan(&user.ID, &user.Name)
                if err != nil {
                        panic(err.Error())
                }
               // fmt.Println(user.ID, user.Name)
                w.Write([]byte(user.Name))
        }

        err = rows.Err()
        if err != nil {
                panic(err.Error())
        }
}
