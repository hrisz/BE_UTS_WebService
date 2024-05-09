package _anidata

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AnimationData struct {
	ID		 	primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
	UniversalID string				`bson:"uid,omitempty" json:"uid,omitempty"`
	Title	 	string		  	   	`bson:"title,omitempty" json:"title,omitempty"`
	Producer 	string				`bson:"producer,omitempty" json:"producer,omitempty"`
	Games	 	string			   	`bson:"games,omitempty" json:"games,omitempty"`
	Dates    	string				`bson:"dates,omitempty" json:"dates,omitempty"`
	URLVid   	string			   	`bson:"urlvid,omitempty" json:"urlvid,omitempty"`
}

type StatisticData struct {
	ID		 	primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
	UniversalID string				`bson:"uid,omitempty" json:"uid,omitempty"`
	Animation	AnimationData		`bson:"animation,omitempty" json:"animation,omitempty"`
	Views    	int					`bson:"views,omitempty" json:"views,omitempty"`
	Likes    	int					`bson:"likes,omitempty" json:"likes,omitempty"`
	Comments 	int					`bson:"comments,omitempty" json:"comments,omitempty"`
}