package utility

import (
	"cloud.google.com/go/translate"
	"context"
	"github.com/autom8ter/api/common"
	"github.com/autom8ter/backend/clientset"
	"github.com/autom8ter/backend/config"
	"golang.org/x/text/language"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func NewEchoer(c *config.Config) *Echoer {
	cc := clientset.NewClientSet(c)
	tr, err := cc.GCP.Translate(context.TODO())
	if err != nil {
		log.Fatalln(err.Error())
	}
	e := &Echoer{
		translator: tr,
	}
	return e
}

type Echoer struct {
	translator *translate.Client
}

func (b *Echoer) EchoSpanish(ctx context.Context, message *common.String) (*common.String, error) {
	resp, err := b.translator.Translate(ctx, []string{message.Text}, language.Spanish, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to translate spanish: %s", err.Error())
	}
	return common.ToString(resp[0].Text), nil
}

func (b *Echoer) EchoChinese(ctx context.Context, message *common.String) (*common.String, error) {
	resp, err := b.translator.Translate(ctx, []string{message.Text}, language.Chinese, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to translate chinese: %s", err.Error())
	}
	return common.ToString(resp[0].Text), nil
}

func (b *Echoer) EchoEnglish(ctx context.Context, message *common.String) (*common.String, error) {
	resp, err := b.translator.Translate(ctx, []string{message.Text}, language.English, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to translate english: %s", err.Error())
	}
	return common.ToString(resp[0].Text), nil

}

func (b *Echoer) EchoHindi(ctx context.Context, message *common.String) (*common.String, error) {
	resp, err := b.translator.Translate(ctx, []string{message.Text}, language.Hindi, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to translate hindi: %s", err.Error())
	}
	return common.ToString(resp[0].Text), nil

}

func (b *Echoer) EchoArabic(ctx context.Context, message *common.String) (*common.String, error) {
	resp, err := b.translator.Translate(ctx, []string{message.Text}, language.Arabic, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to translate arabic: %s", err.Error())
	}
	return common.ToString(resp[0].Text), nil

}

func (b *Echoer) Echo(ctx context.Context, message *common.String) (*common.String, error) {
	return common.ToString(message.Text), nil

}
