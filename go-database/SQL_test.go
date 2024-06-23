package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('reff', 'Reff', 'email@confidential.com', 1000000, 5.0, '2001-05-26', true)"
	// script := "UPDATE customer SET name ='Woy' WHERE id='cuy'"
	// script := "DELETE FROM customer WHERE id='cuy'"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New Customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println("Success Insert New Customer")
}

func TestQueryNext(t *testing.T) {
	//Pengecekan berulang untuk data selanjutnya
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
	}
	defer rows.Close()
}

func TestQuerySqlComplex(t *testing.T) {
	//Penanganan untuk tipe data lain pada SQL Golang
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, created_at, birth_date, married FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id, name, email sql.NullString
		var balance int32
		var rating float64
		var birth_date, created_at sql.NullTime
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &created_at, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		if name.Valid {
			fmt.Println("Name :", name.String)
		}
		if email.Valid {
			fmt.Println("email :", email.String)
		}
		fmt.Println("Balance :", balance)
		fmt.Println("rating :", rating)
		if birth_date.Valid {
			fmt.Println("birth date :", birth_date.Time)
		}
		if created_at.Valid {
			fmt.Println("created at :", created_at.Time)
		}
		fmt.Println("Married :", married)
	}
	defer rows.Close()
}

func TestSqlInjection(t *testing.T) {
	//Salah satu BUG FATAL, ketika user iseng memasukkan query SQL kedalam input
	//dan program memanggap input itu sah (harusnya tidak dieksekusi)
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "Admin'; #"
	password := "admin12345"

	script := "SELECT username FROM user WHERE username = '" + username +
		"' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, script)
	fmt.Println(script)
	if err != nil {
		panic(err)
	}
	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}

	defer rows.Close()
}
func TestSqlInjectionWithParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "Admin"
	password := "admin"

	script := `SELECT username FROM user WHERE username=? AND password=? LIMIT 1`
	rows, err := db.QueryContext(ctx, script, username, password)
	fmt.Println(script)
	if err != nil {
		panic(err)
	}
	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}

	defer rows.Close()
}

func TestExecSqlParameter(t *testing.T) {
	//Function untuk eksekusi jika data didapat dari luar atau user
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := `INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES(?, ?, ?, ?, ?, ?, ?)`
	// script := `UPDATE customer SET name ='Woy' WHERE id=?`
	// script := "DELETE FROM customer WHERE id='cuy'"
	id := "ujang"
	name := "Ujang"
	email := "ujang@mail.com"
	balance := 2000000
	rating := 5.0
	birth_date := "2000-10-10"
	married := true
	_, err := db.ExecContext(ctx, script, id, name, email, balance, rating, birth_date, married)
	// _, err := db.ExecContext(ctx, script, id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New Customer")
}

func TestAutoIncrementSql(t *testing.T) {
	//Auto Increment untuk kasus ID yang otomatis bertambah jika ada data baru
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	email := "rzeko@mail.com"
	comment := "Wah Keren bangets!"

	script := `INSERT INTO comments(email, comment) VALUES(?, ?)`
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}

	InsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Last Comment ID Inserted :", InsertId)
}

func TestPrepareStatement(t *testing.T) {
	//Prepare statement ini cocok jika perlu melakukan eksekusi berkali-kali
	//Agar tidak perlu meminta koneksi secara terus-menerus,
	//Dengan Prepare statement hanya akan meminta koneksi sekali saja dan akan terus dipakai sampai eksekusi selesai
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := `INSERT INTO comments(email, comment) VALUES(?, ?)`
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "reff" + strconv.Itoa(i) + "@mail.com"
		comment := "Comment ke-" + strconv.Itoa(i)

		res, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment ID ke-", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	script := `INSERT INTO comments(email, comment) VALUES(?, ?)`
	statement, err := tx.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 5; i++ {
		email := "reff" + strconv.Itoa(i) + "@mail.com"
		comment := "Comment ke-" + strconv.Itoa(i)

		res, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment ID ke-", id)
	}
	err = tx.Rollback()
	// err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
