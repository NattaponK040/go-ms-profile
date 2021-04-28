package message

import "go.mongodb.org/mongo-driver/bson/primitive"

type Profile struct {
	Id                   primitive.ObjectID `bson:"_id"`
	Email                string             `bson:"email"`
	EmailVerified        bool               `bson:"emailVerified"`
	EmailVerifiedMessage string             `bson:"emailVerifiedMessage"`
	EmailReceive         bool               `bson:"emailReceive"`
	Uid                  string             `bson:"uid"`
	AndroidVersion       string             `bson:"androidVersion"`
	IOSVersion           string             `bson:"iosVersion"`
	JoinTime             struct {
		ISO  string `bson:"iso"`
		Full string `bson:"full"`
	} `bson:"joinTime""`
	Preference struct {
		ShareToFacebook bool `bson:"shareToFacebook"`
	} `bson:"preference"`
	SocialNetworkConnect struct {
		Facebook struct {
			FbId      string `bson:"fbId"`
			Connected bool   `bson:"connected"`
		} `bson:"facebook"`
		Google struct {
			GId       string `bson:"gId"`
			Connected bool   `bson:"connected"`
		} `bson:"google"`
	} `bson:"socialNetworkConnect"`
	Profile struct {
		Name        string `bson:"name"`
		Email       string `bson:"email"`
		AboutMe     string `bson:"aboutMe"`
		PhoneNumber string `bson:"phoneNumber"`
		Gender      struct {
			Id   int    `bson:"id"`
			Name string `bson:"name"`
		} `bson:"gender"`
		ProfilePicture struct {
			Url          string `bson:"url"`
			ThumbnailUrl string `bson:"thumbnailUrl"`
			SmallUrl     string `bson:"smallUrl"`
			Uploaded     bool   `bson:"uploaded"`
		} `bson:"ProfilePicture"`
		City struct {
			Id   string `bson:"id"`
			Name string `bson:"name"`
		} `bson:"city"`
		Age       int `bson:"age"`
		Followers struct {
			Count int        `bson:"count"`
			list  []struct{} `bson:"list"`
		} `bson:"Followers"`
		Following struct {
			Count int        `bson:"count"`
			list  []struct{} `bson:"list"`
		} `bson:"Following"`
	} `bson:"profile"`
}

