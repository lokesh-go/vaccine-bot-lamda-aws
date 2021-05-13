package models

type Get struct {
	ChatID  int64  `bson:"_id" mapstructure:"_id"`
	Pincode string `bson:"pincode" mapstructure:"pincode"`
	Name    string `bson:"name" mapstructure:"name"`
}
