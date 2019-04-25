package utility

import (
	"cloud.google.com/go/translate"
	"context"
	"fmt"
	"github.com/autom8ter/api"
	"github.com/autom8ter/backend/clientset"
	"github.com/autom8ter/backend/config"
	"golang.org/x/text/language"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func NewEchoer(c *config.Config) *Echoer {
	cc := clientset.NewClientSet(c)
	tr, err := cc.GCP.Translate(api.Context)
	if err != nil {
		log.Fatalln(err.Error())
	}
	e := &Echoer{
		Translator: tr,
	}
	return e
}

type Echoer struct {
	Translator *translate.Client
}

func (b *Echoer) EchoSpanish(ctx context.Context, message *api.Message) (*api.Message, error) {
	resp, err := b.Translator.Translate(ctx, []string{message.Value}, language.Spanish, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to translate spanish: %s", err.Error())
	}
	return &api.Message{
		Value: resp[0].Text,
	}, nil
}

func (b *Echoer) EchoChinese(ctx context.Context, message *api.Message) (*api.Message, error) {
	resp, err := b.Translator.Translate(ctx, []string{message.Value}, language.Chinese, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to translate chinese: %s", err.Error())
	}
	return &api.Message{
		Value: resp[0].Text,
	}, nil
}

func (b *Echoer) EchoEnglish(ctx context.Context, message *api.Message) (*api.Message, error) {
	resp, err := b.Translator.Translate(ctx, []string{message.Value}, language.English, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to translate english: %s", err.Error())
	}
	return &api.Message{
		Value: resp[0].Text,
	}, nil
}

func (b *Echoer) EchoHindi(ctx context.Context, message *api.Message) (*api.Message, error) {
	resp, err := b.Translator.Translate(ctx, []string{message.Value}, language.Hindi, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to translate hindi: %s", err.Error())
	}
	return &api.Message{
		Value: resp[0].Text,
	}, nil
}

func (b *Echoer) EchoArabic(ctx context.Context, message *api.Message) (*api.Message, error) {
	resp, err := b.Translator.Translate(ctx, []string{message.Value}, language.Arabic, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to translate arabic: %s", err.Error())
	}
	return &api.Message{
		Value: resp[0].Text,
	}, nil
}

func (b *Echoer) Echo(ctx context.Context, e *api.Message) (*api.Message, error) {
	return &api.Message{
		Value: fmt.Sprintf("echoed: %s", e.Value),
	}, nil

}
