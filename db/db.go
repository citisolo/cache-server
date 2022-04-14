package db

import "github.com/hashicorp/go-memdb"

type UserContext struct {
	Id string   `json:"id"`
	Token string `json:"token"`
}

type Cache interface {
	InitDB()
	Insert(data *UserContext) error
	Get(key string) *UserContext
}

type CacheManager struct {
	cache *memdb.MemDB
}

const (
	tableName = "usercontext"
	index = "id"
)


func  NewDB() *memdb.MemDB {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			tableName: &memdb.TableSchema{
				Name: tableName,
				Indexes: map[string]*memdb.IndexSchema{
					index: &memdb.IndexSchema{
						Name: index,
						Unique: true,
						Indexer: &memdb.StringFieldIndex{Field: "Id"},
					},
					"token":&memdb.IndexSchema{
						Name: "token",
						Unique: false,
						Indexer: &memdb.StringFieldIndex{Field: "Token"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	return db
}

func (c* CacheManager) InitDB() {
	c.cache = NewDB()
}

func (c* CacheManager) Insert(data *UserContext) error {
   transaction := c.cache.Txn(true)

   if err := transaction.Insert(tableName, data); err != nil {
	   return err
   }

   transaction.Commit()

   return nil
}

func (c* CacheManager) Get(key string) *UserContext{
  transaction := c.cache.Txn(false)
  defer transaction.Abort()

  raw, err := transaction.First(tableName, index,  key)
  if err != nil {
	  return nil
  }
  var data *UserContext
  if raw != nil {
	data = raw.(*UserContext)
  } else {
	  return nil
  }

  return data
}