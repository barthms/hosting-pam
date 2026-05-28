package services

import (
	"context"
	"log"
	"os"
	"path/filepath"

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
	serviceAccountPath := viper.GetString("FIREBASE_SERVICE_ACCOUNT_KEY_PATH")
	if serviceAccountPath == "" {
		log.Printf("[FCM] FIREBASE_SERVICE_ACCOUNT_KEY_PATH not set, notifier disabled")
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
