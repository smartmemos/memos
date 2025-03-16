package model

type Profile struct {
	Owner       string
	Mode        string
	Version     string
	InstanceUrl string
}

type GetProfileRequest struct {
}
