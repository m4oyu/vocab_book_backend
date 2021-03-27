package controller

import (
	"cloud.google.com/go/translate"
	"context"
	"fmt"
	"golang.org/x/text/language"
)

type TranslateHandler interface {
	TranslateText(targetLanguage, text string) (string, error)
}

type translateHandler struct {

}

func NewTranslateHandler() TranslateHandler {
	return &translateHandler{}
}

func (t *translateHandler)TranslateText(targetLanguage, text string) (string, error) {
	// user get
	// request receive
	// throw translate request to gcp api
	// insert into db
	// return response


	// text := "The Go Gopher is cute"
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}

	// 翻訳情報をDBに格納

	return resp[0].Text, nil
}