package courtsdk

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"strconv"
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
