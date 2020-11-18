package sender

import (
	"bytes"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/wmrodrigues/twitter-sender/internal/structs"
	"log"
	"os"
	"strings"
	"text/template"
)

const TweetLimitLen = 150

type Twitter struct {
	settings    structs.Settings
	recipients  []structs.Recipient
	set         structs.RecipientSet
	accessToken string
}

// NewTwitter creates a new instance of Twitter struct
func NewTwitter(settingsFile structs.Settings) *Twitter {
	return &Twitter{settings: settingsFile}
}

// SetRecipients set the recipients attribute
func (t *Twitter) SetRecipients(r []structs.Recipient) {
	t.recipients = r
}

func getMessageTemplate() *template.Template {
	wd, err := os.Getwd()
	if err != nil {
		err = fmt.Errorf("error getting working dir for template file, %s", err.Error())
		log.Fatal(err)
	}

	templateFilePath := fmt.Sprintf("%s/configs/message.template", wd)
	_template := template.Must(template.ParseFiles(templateFilePath))
	return _template
}

func (t *Twitter) buildMessage(set structs.RecipientSet) string {
	_template := getMessageTemplate()
	var content bytes.Buffer
	err := _template.Execute(&content, set)

	if err != nil {
		err = fmt.Errorf("error loading content on template, %s", err.Error())
		log.Fatal(err)
	}

	return content.String()
}

func (t *Twitter) tweet(set structs.RecipientSet) {
	message := t.buildMessage(set)

	config := oauth1.NewConfig(t.settings.Twitter.Key, t.settings.Twitter.Secret)
	token := oauth1.NewToken(t.settings.Twitter.AccessToken, t.settings.Twitter.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	tweet, resp, err := client.Statuses.Update(message, nil)

	if err != nil {
		err = fmt.Errorf("error sending tweet, %s", err.Error())
		log.Println(err)
	}

	log.Println(message)
	log.Println("**** tweet ****", tweet.ID)
	log.Println("**** response ****", resp.StatusCode, resp.Status)
}

func (t *Twitter) SendTweets() {
	set := structs.RecipientSet{}

	maleMnemonic := t.settings.Message.MaleMnemonic
	femaleMnemonic := t.settings.Message.FemaleMnemonic

	maleRecipients := make([]string, 0)
	femaleRecipients := make([]string, 0)

	for _, item := range t.recipients {
		if item.Treatment == maleMnemonic && item.Username != "" {
			maleRecipients = append(maleRecipients, item.Username)
			set.HasCongressmen = true
		}
		// just making sure we're considering only settings mnemonics
		if item.Treatment == femaleMnemonic && item.Username != "" {
			femaleRecipients = append(femaleRecipients, item.Username)
			set.HasCongresswomen = true
		}

		_len := strings.Join(femaleRecipients, ", ") + strings.Join(maleRecipients, ", ")

		if len(maleRecipients) + len(femaleRecipients) == t.settings.Message.Qty || len(_len) > TweetLimitLen {
			set.Congresswomen = strings.Join(femaleRecipients, ", ")
			set.Congressmen = strings.Join(maleRecipients, ", ")

			t.tweet(set)

			femaleRecipients = nil
			maleRecipients = nil
			femaleRecipients = make([]string, 0)
			maleRecipients = make([]string, 0)
			set = structs.RecipientSet{}
		}
	}

}