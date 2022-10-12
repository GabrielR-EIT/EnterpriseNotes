package main

import (
	"log"

	"github.com/sony/sonyflake"
	_ "github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

// Initialize SonyFlake Function
func InitSonyFlake() string {
	var st sonyflake.Settings

	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		log.Fatal(sf)
		return "An error occurred when creating new SonyFlake\n"
	}
	return "SonyFlake was initialised successfully\n"
}

// Generate ID Function
func NextID() uint64 {
	id, err := sf.NextID()
	if err != nil {
		log.Fatal("An error occurred when generating a new ID")
	}
	return id
}
