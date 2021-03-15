package connection

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinxing-go/mysql"
)

var DB *mysql.MySQl
