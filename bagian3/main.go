package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString(
		`<html>
<head>
	<title>Demo1</title>
</head>
<body>
	<form method='post'>
		<label for="name">Nama anda</label>
		<input name='name' />
		<label for="age">Umur anda</label>
		<input name='age' type='number' />
		<input type='submit' value='send' />
	</form>
</body>
</html>
`)
	ctx.Response.Header.Set(`content-type`, `text/html`)
}

func hasError(ctx *fasthttp.RequestCtx, err error, str string) bool {
	if err != nil {
		fmt.Println(err, str)
		str, _ := jsoniter.Marshal(map[string]interface {
		}{
			`error`: str,
		})
		fmt.Fprintln(ctx, string(str))
		return true
	}
	return false
}

func isError(ctx *fasthttp.RequestCtx, err bool, str string) bool {
	if err {
		fmt.Println(str)
		str, _ := jsoniter.Marshal(map[string]interface {
		}{
			`error`: str,
		})
		fmt.Fprintln(ctx, string(str))
		return true
	}
	return false
}

type ApiResponse struct {
	Nama                  string
	KemungkinanTahunLahir int
	JumlahKata            int
	PanjangNama           int
}

func ApiCoba(ctx *fasthttp.RequestCtx) {
	args := ctx.Request.PostArgs()
	name := ``
	age := 0
	var err error

	args.VisitAll(func(key []byte, val []byte) {
		str := string(val)
		if err != nil {
			return
		}
		switch string(key) {
		case `name`:
			name = str
			if isError(ctx, name == ``, `name must not be empty`) {
				err = fmt.Errorf(`empty name`)
				return
			}
		case `age`:
			age, err = strconv.Atoi(str)
			if hasError(ctx, err, `cannot read age: `+str) {
				return
			}
		}
	})
	if err != nil {
		return
	}
	if isError(ctx, age < 0, `age cannot be negative`) {
		return
	}

	res := ApiResponse{}
	res.Nama = name
	res.KemungkinanTahunLahir = time.Now().Year() - age
	res.JumlahKata = len(strings.Fields(name))
	res.PanjangNama = len(name)
	json, _ := jsoniter.Marshal(res)
	fmt.Fprintln(ctx, string(json))
}

func main() {
	r := router.New()
	r.GET("/", Index)
	r.POST("/", ApiCoba)

	const addr = ":8090"
	fmt.Println("Started server on address " + addr)

	log.Fatal(fasthttp.ListenAndServe(addr, r.Handler))
}
