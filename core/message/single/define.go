package single

type Options struct {
	Safe                   int
	EnableIdTrans          int
	EnableDuplicateCheck   int
	DuplicateCheckInterval int
}

type OptionFunc func(options *Options)

func WithOptionSafe(safe int) OptionFunc {
	return func(options *Options) {
		options.Safe = safe
	}
}

func WithOptionEnableIdTrans(enableIdTrans int) OptionFunc {
	return func(options *Options) {
		options.EnableIdTrans = enableIdTrans
	}
}

func WithOptionEnableDuplicateCheck(enableDuplicateCheck int) OptionFunc {
	return func(options *Options) {
		options.EnableDuplicateCheck = enableDuplicateCheck
	}
}

func WithOptionDuplicateCheckInterval(duplicateCheckInterval int) OptionFunc {
	return func(options *Options) {
		options.DuplicateCheckInterval = duplicateCheckInterval
	}
}

type MessageResp struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
}

type CommonBody struct {
	ToUser                 string `json:"touser"`
	ToParty                string `json:"toparty"`
	ToTag                  string `json:"totag"`
	MsgType                string `json:"msgtype"`
	AgentId                int    `json:"agentid"`
	Safe                   int    `json:"safe"`
	EnableIdTrans          int    `json:"enable_id_trans"`
	EnableDuplicateCheck   int    `json:"enable_duplicate_check"`
	DuplicateCheckInterval int    `json:"duplicate_check_interval"`
}

type TextMessage struct {
	CommonBody
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

type ImageMessage struct {
	CommonBody
	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

type VoiceMessage struct {
	CommonBody
	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

type VideoMessage struct {
	CommonBody
	Video struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"video"`
}

type FileMessage struct {
	CommonBody
	File struct {
		MediaId string `json:"media_id"`
	} `json:"file"`
}

type TextCardMessage struct {
	CommonBody
	TextCard struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		BtnTxt      string `json:"btntxt"`
	} `json:"textcard"`
}

type NewsMessage struct {
	CommonBody
	News struct {
		Articles []Article `json:"articles"`
	} `json:"news"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
}

type MpNewsMessage struct {
	CommonBody
	MpNews struct {
		Articles []MpArticle `json:"articles"`
	} `json:"mpnews"`
}

type MpArticle struct {
	Title            string `json:"title"`
	ThumbMediaId     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	ContentSourceUrl string `json:"content_source_url"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
}

type MarkdownMessage struct {
	CommonBody
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}
