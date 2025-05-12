package main

import "backend/pkg/storage"

func init() {
	if err := storage.InitPostgres("postgres://xl:xl@192.168.1.208:5432/xl"); err != nil {
		panic(err)
	}
}

func main() {

}
