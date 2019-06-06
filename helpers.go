package courtsdk

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

//Debug returns if the current env is safe to debug.
func Debug() string {
	return strings.ToUpper(os.Getenv("DEBUG"))
}

//DebugPrint - log only in dev environment.
func DebugPrint(v ...interface{}) {
	if Debug() == "CONFIG" || Debug() == "ALL" {
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
	collector := colly.NewCollector(colly.Async(EngineConfig["IsAsync"].(bool)))
	if Debug() == "REQUEST" || Debug() == "ALL" {
		collector = colly.NewCollector(colly.Async(EngineConfig["IsAsync"].(bool)),
			colly.Debugger(&debug.LogDebugger{}))
	}
	transport := &http.Transport{}
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	collector.WithTransport(transport)
	return collector
}

//GetNewContext - return a new context with default timeout and context cancelation function.
func GetNewContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), EngineConfig["RequestTimeout"].(time.Duration)*time.Second)
}

//GenerateMD5 - returns the MD5 hash of the given pointer value.
func GenerateMD5(data *string) string {
	hasher := md5.New()
	hasher.Write([]byte(*data))
	return hex.EncodeToString(hasher.Sum(nil))
}

//RemoveUnusedChars remove unused chars from string
func RemoveUnusedChars(data string) string {
	pattern := regexp.MustCompile(`(\s*<!--.*-->\s*)|(\s+)`)
	return pattern.ReplaceAllString(data, " ")
}

//HasMaxFailures check if reached the max failures
func HasMaxFailures(failures *int) bool {
	return *failures >= EngineConfig["MaxFailures"].(int)
}

//GetElasticMapping - returns the default Elasticsearch mapping
func GetElasticMapping() string {
	return `{
		"settings": {
			"index": {
				"number_of_shards": 10,
				"number_of_replicas": 0
			}
		},
		"mappings": {
			"_doc": {
				"_all": {
					"type": "text",
					"index": "true",
					"analyzer": "brazilian"
				},
				"properties": {
					"court": {
						"type": "text"
					},
					"document_type": {
						"type": "keyword"
					},
					"document_id": {
						"type": "keyword"
					},
					"is_enabled": {
						"type": "boolean"
					},
					"checksum": {
						"type": "text"
					},
					"full_document_link": {
						"type": "text"
					},
					"content": {
						"type": "text",
						"index": "true",
						"analyzer": "brazilian"
					},
					"judged_at": {
						"type": "date"
					},
					"inserted_at": {
						"type": "date"
					},
					"updated_at": {
						"type": "date"
					}
				}
			}
		}
	}`
}
