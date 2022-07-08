package badgerbrowserweb

import (
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/gin-gonic/gin"
)

var Db *badger.DB

func Index(c *gin.Context) {

	c.Redirect(301, "/web/html/layout.html")

}

func DeleteKey(c *gin.Context) {

	if c.PostForm("key") == "" {
		c.String(200, "no key | n")
	}

	Db.Update(func(tx *badger.Txn) error {
		err := tx.Delete([]byte(c.PostForm("key")))

		if err != nil {

			c.String(200, "error Deleting KV | n")
			return fmt.Errorf("delete kv: %s", err)
		}

		return nil
	})

	c.String(200, "ok")

}

func Set(c *gin.Context) {

	if c.PostForm("key") == "" {
		c.String(200, "no key | n")
	}

	Db.Update(func(tx *badger.Txn) error {
		err := tx.Set([]byte(c.PostForm("key")), []byte(c.PostForm("value")))

		if err != nil {

			c.String(200, "error writing KV | n")
			return fmt.Errorf("create kv: %s", err)
		}

		return nil
	})

	c.String(200, "ok")

}

func Get(c *gin.Context) {

	res := []string{"nok", ""}

	if c.PostForm("key") == "" {

		res[1] = "no key | n"
		c.JSON(200, res)
	}

	Db.View(func(tx *badger.Txn) error {

		v, err := tx.Get([]byte(c.PostForm("key")))
		if err != nil {
			res[1] = "no such key | n"
			c.JSON(200, res)
			return fmt.Errorf("get kv: %s", err)
		}

		res[0] = "ok"
		val, err := v.ValueCopy(nil)
		res[1] = string(val)

		fmt.Printf("Key: %s\n", v)
		return nil

	})

	c.JSON(200, res)

}

type Result struct {
	Result string
	M      map[string]string
}

func PrefixScan(c *gin.Context) {

	res := Result{Result: "nok"}

	res.M = make(map[string]string)

	count := 0

	Db.View(func(tx *badger.Txn) error {
		prefix := []byte(c.PostForm("key"))
		it := tx.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			_ = item.Value(func(val []byte) error {
				res.M[string(k)] = string(val)
				return nil
			})
			count++
			if count > 2000 {
				break
			}
		}

		res.Result = "ok"

		return nil
	})

	c.JSON(200, res)

}
