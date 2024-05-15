package schedule

type UploadScheduleDTO struct {
	Weeks []WeekDTO
}

type WeekDTO struct {
	IsEven bool
	Days   []DayDTO
}

type DayDTO struct {
	DayNumber int
	Subjects  []SubjectDTO
}

type SubjectDTO struct {
	Name        string
	Cabinet     string
	Teacher     string
	Description string
	StartTime   string
	EndTime     string
}
