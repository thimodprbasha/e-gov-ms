package model

type Complain struct {
	Model       `bson:",inline"`
	MID_        string     `bson:"_id,omitempty" json:"-"`
	ComplainID  string     `bson:"complain_id" json:"complain_id"`
	FullName string `bson:"full_name" json:"full_name"`
	TelephoneNumber string `json:"telephone_number" bson:"telephone_number"`
	Department  string     `json:"department" bson:"department"`
	Description string     `bson:"description" json:"description"`
	ComplainFeedback ComplainFeedback `bson:"complain_feedback" json:"complain_feedback"`
	ComplainApproved ComplainApproved `bson:"complain_approved" json:"complain_approved"`
}

type ComplainDetails struct {
	NumberOfWorkers  int64    `bson:"number_of_workers" json:"number_of_workers"`
	NameOfWorkers   string `json:"name_of_workers" bson:"name_of_workers"`
	Time            string `bson:"time" json:"time"`
}

type ComplainFeedback struct {
	Pending bool `json:"pending" bson:"pending"`
	SupervisorID     string          `bson:"supervisor_id" json:"supervisor_id"`
	FeedbackApproved bool            `bson:"feedback_approved" json:"feedback_approved"`
	ComplainDetails  ComplainDetails `bson:"complain_details" json:"complain_details"`
}

type ComplainApproved struct {
	GovWorkerID       string `bson:"gov_worker_id" json:"gov_worker_id"`
	Approved        bool   `bson:"approved" bson:"approved"`
	CitizenFeedback string `bson:"citizen_feedback" json:"citizen_feedback"`
	GovWorkerFeedback string `bson:"gov_worker_feedback" json:"gov_worker_feedback"`
}
