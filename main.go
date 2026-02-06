package main

import (
	"context"
	"log"

	sdk "github.com/apito-io/go-apito-plugin-sdk"
)

func main() {
	log.Printf("🎯 [hc-logging-plugin] Starting Logging plugin...")
	plugin := sdk.Init("hc-logging-plugin", "1.0.0", "apito-plugin-key")
	statusType := sdk.NewObjectType("LoggingStatus", "Logging plugin status").
		AddStringField("status", "Plugin status", false).
		AddStringField("version", "Plugin version", false).
		Build()
	plugin.RegisterQuery("getLoggingStatus",
		sdk.ComplexObjectField("Get logging plugin status", statusType),
		func(ctx context.Context, rawArgs map[string]interface{}) (interface{}, error) {
			return map[string]interface{}{"status": "ready", "version": "1.0.0"}, nil
		})
	log.Printf("🚀 [hc-logging-plugin] Plugin ready")
	plugin.Serve()
}
