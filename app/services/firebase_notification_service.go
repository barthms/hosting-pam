package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

type NotificationSender interface {
	SendToToken(ctx context.Context, token, title, body string, data map[string]string) error
}

type NoopNotifier struct {
	reason string
}

func (n *NoopNotifier) SendToToken(ctx context.Context, token, title, body string, data map[string]string) error {
	return nil
}

type FirebaseNotifier struct {
	client *messaging.Client
}

func (n *FirebaseNotifier) SendToToken(ctx context.Context, token, title, body string, data map[string]string) error {
	if token == "" {
		return nil
	}

	msg := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: data,
		Android: &messaging.AndroidConfig{
			Priority: "high",
			Notification: &messaging.AndroidNotification{
				ChannelID: "kia_channel_high_priority",
			},
		},
	}

	_, err := n.client.Send(ctx, msg)
	return err
}

func NewFirebaseNotifierFromEnv() NotificationSender {
	if rawJSON := strings.TrimSpace(viper.GetString("FIREBASE_SERVICE_ACCOUNT_JSON")); rawJSON != "" {
		credentialsJSON, err := parseFirebaseJSON(rawJSON)
		if err != nil {
			log.Printf("[FCM] invalid FIREBASE_SERVICE_ACCOUNT_JSON: %v", err)
			return &NoopNotifier{reason: "invalid json"}
		}

		ctx := context.Background()
		app, err := firebase.NewApp(ctx, nil, option.WithCredentialsJSON(credentialsJSON))
		if err != nil {
			log.Printf("[FCM] failed to init firebase app from JSON env: %v", err)
			return &NoopNotifier{reason: "init failed"}
		}

		client, err := app.Messaging(ctx)
		if err != nil {
			log.Printf("[FCM] failed to init messaging client: %v", err)
			return &NoopNotifier{reason: "client failed"}
		}

		return &FirebaseNotifier{client: client}
	}

	if encodedJSON := strings.TrimSpace(viper.GetString("FIREBASE_SERVICE_ACCOUNT_JSON_BASE64")); encodedJSON != "" {
		rawJSON, err := base64.StdEncoding.DecodeString(encodedJSON)
		if err != nil {
			log.Printf("[FCM] failed to decode FIREBASE_SERVICE_ACCOUNT_JSON_BASE64: %v", err)
			return &NoopNotifier{reason: "decode failed"}
		}

		ctx := context.Background()
		app, err := firebase.NewApp(ctx, nil, option.WithCredentialsJSON(rawJSON))
		if err != nil {
			log.Printf("[FCM] failed to init firebase app from base64 env: %v", err)
			return &NoopNotifier{reason: "init failed"}
		}

		client, err := app.Messaging(ctx)
		if err != nil {
			log.Printf("[FCM] failed to init messaging client: %v", err)
			return &NoopNotifier{reason: "client failed"}
		}

		return &FirebaseNotifier{client: client}
	}

	serviceAccountPath := viper.GetString("FIREBASE_SERVICE_ACCOUNT_KEY_PATH")
	if serviceAccountPath == "" {
		log.Printf("[FCM] firebase credentials not configured, notifier disabled")
		return &NoopNotifier{reason: "missing env"}
	}

	resolvedPath := resolveServiceAccountPath(serviceAccountPath)
	if resolvedPath == "" {
		log.Printf("[FCM] service account file not found: %s", serviceAccountPath)
		return &NoopNotifier{reason: "file missing"}
	}

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(resolvedPath))
	if err != nil {
		log.Printf("[FCM] failed to init firebase app: %v", err)
		return &NoopNotifier{reason: "init failed"}
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Printf("[FCM] failed to init messaging client: %v", err)
		return &NoopNotifier{reason: "client failed"}
	}

	return &FirebaseNotifier{client: client}
}

func parseFirebaseJSON(raw string) ([]byte, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return nil, fmt.Errorf("empty firebase json")
	}

	if json.Valid([]byte(trimmed)) {
		return []byte(trimmed), nil
	}

	decoded, err := base64.StdEncoding.DecodeString(trimmed)
	if err != nil {
		return nil, err
	}

	if !json.Valid(decoded) {
		return nil, fmt.Errorf("decoded firebase json is invalid")
	}

	return decoded, nil
}

func resolveServiceAccountPath(rawPath string) string {
	if rawPath == "" {
		return ""
	}

	candidates := []string{rawPath}
	if !filepath.IsAbs(rawPath) {
		candidates = append(candidates, filepath.Join("cmd", rawPath))
	}

	for _, candidate := range candidates {
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}

	return ""
}
