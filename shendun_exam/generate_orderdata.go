package shendun_exam

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
	"time"
)

const IdNum = 1000

// const OrderNum = 1000000
const OrderNum = 100

var idSlice []int
var db *sql.DB

func generateId() []int {
	generatedIds := make(map[int]bool)
	userIds := make([]int, IdNum)
	for i := 0; i < IdNum; {
		id := rand.Intn(1000000)
		if _, exists := generatedIds[id]; !exists {
			generatedIds[id] = true
			userIds[i] = id
			i++
		}
	}
	return userIds
}

func GenerateOrder() {
	var err error
	err = initDdAndTable()
	checkErr(err)
	idSlice = generateId()
	for i := 0; i < OrderNum; i++ {
		err = insertToSql(idSlice[rand.Intn(IdNum)], generateWeight())
		checkErr(err)
	}
	_, err = db.Exec("CREATE INDEX uid_idx ON express_order (uid);")
	checkErr(err)
	db.Close()
}

func insertToSql(uid int, weight float64) error {
	if db == nil {
		return errors.New("db not init")
	}
	insertExpressOrder, err := db.Prepare("insert into express_order(uid,weight) values (?,?);")
	res, err := insertExpressOrder.Exec(uid, weight)
	if err != nil {
		return err
	}
	_, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func initDdAndTable() error {
	var err error
	db, err = sql.Open("sqlite3", "./order.sqlite.db")
	if err != nil {
		return err
	}
	sqlTable := `
    CREATE TABLE IF NOT EXISTS express_order(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        uid INTEGER NOT NULL,
        weight DOUBLE NOT NULL ,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    `
	_, err = db.Exec(sqlTable)
	if err != nil {
		return err
	}
	return nil
}

func generateWeight() (re float64) {
	//有各事件出现具体概率的话，根据累积概率分布就可以获取发生事件。
	// 计算累积概率
	probabilities := make([]float64, 100)
	sum := 0.0
	for i := 0; i < 100; i++ {
		sum += 1.0 / (float64(i+1) * 1.0)
		probabilities[i] = sum
	}
	// 生成随机数
	r := rand.Float64() * sum
	// 找到对应的整数,然后取其范围内的double
	for i, p := range probabilities {
		if r <= p {
			re = float64(i+1) - rand.Float64()
			break
		}
	}
	return
}

func QueryOrderById(uid int) (string, error) {
	var err error
	var weightTotal float64
	var rt string
	var count int
	var totalFee int
	if db == nil {
		db, err = sql.Open("sqlite3", "./order.sqlite.db")
		if err != nil {
			return "", err
		}
	}
	rows, err := db.Query("SELECT id, weight,created_at FROM express_order where uid=?;", uid)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var weight float64
		var createdAt time.Time
		err := rows.Scan(&id, &weight, &createdAt)
		if err != nil {
			return "", err
		}
		weightTotal += weight
		fee := GetCourierFee(weight)
		totalFee += fee
		count += 1
		rt += fmt.Sprintf("No.%v, order id: %v, weight: %v, created time: %v, fee:%v\n", count, id, weight, createdAt, fee)
	}
	rt += fmt.Sprintf("total count:%v, total weight: %v, total fee: %v\n", count, weightTotal, totalFee)
	return rt, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
