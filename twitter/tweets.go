// Tweet represents a Twitter Tweet, previously called a status.
package twitter

// https://dev.twitter.com/overview/api/tweets
type Tweet struct {
	CreatedAt            string                 `json:"created_at"`
	FavoriteCount        int                    `json:"favorite_count"`
	Favorited            bool                   `json:"favorited"`
	FilterLevel          string                 `json:"filter_level"`
	ID                   int64                  `json:"id"`
	IDStr                string                 `json:"id_str"`
	InReplyToScreenName  string                 `json:"in_reply_to_screen_name"`
	InReplyToStatusID    int64                  `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr string                 `json:"in_reply_to_status_id_str"`
	InReplyToUserID      int64                  `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   string                 `json:"in_reply_to_user_id_str"`
	Lang                 string                 `json:"lang"`
	PossiblySensitive    bool                   `json:"possibly_sensitive"`
	QuoteCount           int                    `json:"quote_count"`
	ReplyCount           int                    `json:"reply_count"`
	RetweetCount         int                    `json:"retweet_count"`
	Retweeted            bool                   `json:"retweeted"`
	RetweetedStatus      *Tweet                 `json:"retweeted_status"`
	Source               string                 `json:"source"`
	Scopes               map[string]interface{} `json:"scopes"`
	Text                 string                 `json:"text"`
	FullText             string                 `json:"full_text"`
	Truncated            bool                   `json:"truncated"`
	User                 *User                  `json:"user"`
	WithheldCopyright    bool                   `json:"withheld_copyright"`
	WithheldInCountries  []string               `json:"withheld_in_countries"`
	WithheldScope        string                 `json:"withheld_scope"`
	QuotedStatusID       int64                  `json:"quoted_status_id"`
	QuotedStatusIDStr    string                 `json:"quoted_status_id_str"`
	QuotedStatus         *Tweet                 `json:"quoted_status"`
}
