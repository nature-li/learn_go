package main

import (
	"sync"
	"sync/atomic"
	"log"
	"time"
	"math/rand"
)

const (
	maxCoroutine = 5
)

var (
	idCounter int32
)

type dbConnection struct {
	ID int32
}

func (db *dbConnection) Close() error {
	log.Println("关闭连接", db.ID)
	return nil
}

func createConnection() interface{} {
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{id}
}

func dbQuery(query int, pool *sync.Pool) {
	conn := pool.Get().(*dbConnection)
	defer pool.Put(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("第%d个查询，使用的是ID为%d的数据库连接", query, conn.ID)
}

func main_in_pool() {
	var wg sync.WaitGroup
	wg.Add(maxCoroutine)

	p := &sync.Pool{
		New: createConnection,
	}

	for query := 0; query < maxCoroutine; query++ {
		go func(q int) {
			dbQuery(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("开始关闭资源池")
}
