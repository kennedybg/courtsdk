package courtsdk

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

//DebugEnv returns if the current env is safe to debug.
func DebugEnv() bool {
	env := os.Getenv("environment")
	return env != "prod" && env != "stg"
}

//DebugPrint - log only in dev environment.
func DebugPrint(v ...interface{}) {
	if DebugEnv() {
		log.Println(v...)
	}
}

//GetEnvInt return the given env int var
func GetEnvInt(envVar string, Default int) int {
	num, err := strconv.Atoi(os.Getenv(envVar))
	if err != nil {
		DebugPrint("[WARNING] ENV VAR ->", envVar, "not found, using default value ->", strconv.Itoa(Default), ".")
		return Default
	}
	return num
}

//GetEnvString return the given env string var
func GetEnvString(envVar string, Default string) string {
	str := os.Getenv(envVar)
	if str == "" {
		DebugPrint("[WARNING] ENV VAR ->", envVar, "not found, using default value ->", Default, ".")
		return Default
	}
	return str
}

//GetDefaultcollector - return the default collector (colly)
func GetDefaultcollector() *colly.Collector {
	collector := colly.NewCollector(colly.Async(EngineConfig["isAsync"].(bool)))
	if DebugEnv() {
		collector = colly.NewCollector(colly.Async(EngineConfig["isAsync"].(bool)),
			colly.Debugger(&debug.LogDebugger{}))
	}
	transport := &http.Transport{}
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	collector.WithTransport(transport)
	return collector
}
