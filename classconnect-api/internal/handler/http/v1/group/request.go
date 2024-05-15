package group

type CreateGroupRequest struct {
	Name string `json:"name" binding:"required,min=3"`
}

type JoinToGroupRequest struct {
	ID   uint64 `json:"sub_id" binding:"required"`
	Code string `json:"code" binding:"required,min=4,max=4"`
}

type LeaveFromGroupRequest struct {
	ID uint64 `json:"sub_id" binding:"required"`
}
