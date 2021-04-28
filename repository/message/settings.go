package message

type Settings struct {
	AccountPrivacy bool   `bson:"accountPrivacy"`
	Language       string `bson:"language"`
	Notification   struct {
		Message struct {
			Ding             bool `bson:"ding"`
			PushNotification bool `bson:"pushNotification"`
		} `bson:"message"`
		FeedBack struct {
			Links            bool `bson:"links"`
			Shared           bool `bson:"shared"`
			Comment          bool `bson:"comment"`
			PushNotification bool `bson:"pushNotification"`
		} `bson:"feedBack"`
		Event struct {
			Followers        bool `bson:"followers`
			Following        bool `bson:"following"`
			GroupInvitations bool `bson:"groupInvitations"`
			News             bool `bson:"news"`
			Bookmarks        bool `bson:"bookmarks"`
		} `bson:"event"`
		Locale   string `bson:"locale"`
		Birthday struct {
			ISO  string `bson:"iso"`
			Full string `bson:"full"`
		}
	} `bson:"notification"`
}
