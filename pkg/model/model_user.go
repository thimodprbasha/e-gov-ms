package model

type User struct {
	Model        `bson:",inline"`
	MID_         string  `bson:"_id,omitempty" json:"-"`
	ID           string  `json:"id" bson:"id"`
	Email        string  `json:"email" bson:"email"`
	Password     string  `json:"-"  bson:"password"`
	Role         string  `json:"role" bson:"role"`
	Complain     []*ComplainIDs `json:"complain" bson:"complain"`
}

type ComplainIDs struct {
	ComplainID string `json:"complain_id" bson:"complain_id"`
}
