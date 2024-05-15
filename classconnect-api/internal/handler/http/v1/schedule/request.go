package schedule

import "github.com/tclutin/classconnect-api/internal/domain/schedule"

type GetScheduleForDayRequest struct {
	ID uint64 `json:"sub_id" binding:"required"`
}

type UploadScheduleRequest struct {
	Weeks []Week `json:"weeks" binding:"required"`
}

type Week struct {
	IsEven bool  `json:"is_even" binding:"required"`
	Days   []Day `json:"days" binding:"required"`
}

type Day struct {
	DayNumber int       `json:"day_number" binding:"required"`
	Subjects  []Subject `json:"subjects" binding:"required"`
}

type Subject struct {
	Name        string `json:"name" binding:"required"`
	Cabinet     string `json:"cabinet" binding:"required"`
	Teacher     string `json:"teacher" binding:"required"`
	Description string `json:"description" binding:"required"`
	StartTime   string `json:"start_time" binding:"required"`
	EndTime     string `json:"end_time" binding:"required"`
}

func (u UploadScheduleRequest) ToDTO() schedule.UploadScheduleDTO {
	var weeksDTO []schedule.WeekDTO
	for _, week := range u.Weeks {
		weeksDTO = append(weeksDTO, week.ToDTO())
	}
	return schedule.UploadScheduleDTO{
		Weeks: weeksDTO,
	}
}

func (w Week) ToDTO() schedule.WeekDTO {
	var daysDTO []schedule.DayDTO
	for _, day := range w.Days {
		daysDTO = append(daysDTO, day.ToDTO())
	}
	return schedule.WeekDTO{
		IsEven: w.IsEven,
		Days:   daysDTO,
	}
}

func (d Day) ToDTO() schedule.DayDTO {
	var subjectsDTO []schedule.SubjectDTO
	for _, subject := range d.Subjects {
		subjectsDTO = append(subjectsDTO, subject.ToDTO())
	}
	return schedule.DayDTO{
		DayNumber: d.DayNumber,
		Subjects:  subjectsDTO,
	}
}

func (s Subject) ToDTO() schedule.SubjectDTO {
	return schedule.SubjectDTO{
		Name:        s.Name,
		Cabinet:     s.Cabinet,
		Teacher:     s.Teacher,
		Description: s.Description,
		StartTime:   s.StartTime,
		EndTime:     s.EndTime,
	}
}
