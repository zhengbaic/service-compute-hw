package Meeting

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/weimumu/Agenda/agenda/entity/User"
)

type Meetings struct {
	Initiator    string
	Participator []string
	Title        string
	STime        time.Time
	ETime        time.Time
}

//时间处理函数
func DateToString(date time.Time) string {
	return date.Format("2006-01-02/15:04")
}
func StringToDate(date string) (time.Time, error) {
	the_time, err := time.Parse("2006-01-02/15:04", date)
	return the_time, err
}
func SmallDate(date1, date2 time.Time) bool {
	return date1.Before(date2) || date1.Equal(date2)
}
func LargeDate(date1, date2 time.Time) bool {
	return date1.After(date2) || date1.Equal(date2)
}

//返回所有会议
func returnAllMeeting() []Meetings {
	var allMeeting []Meetings
	stream, _ := ioutil.ReadFile("Meeting.json")
	json.Unmarshal(stream, &allMeeting)
	return allMeeting
}

//发起者会议
func ReturnInitiatorMeeting(userName string) []Meetings {
	allMeeting := returnAllMeeting()
	var userMeeting []Meetings
	for _, value := range allMeeting {
		if userName == value.Initiator {
			userMeeting = append(userMeeting, value)
		}
	}
	return userMeeting
}

//参与者会议
func ReturnParticipatorMeeting(userName string) []Meetings {
	allMeeting := returnAllMeeting()
	var userMeeting []Meetings
	for _, value := range allMeeting {
		if userName == value.Initiator {
			continue
		}
		for _, value1 := range value.Participator {
			if userName == value1 {
				userMeeting = append(userMeeting, value)
				break
			}
		}
	}
	return userMeeting
}

//参与者会议
func ReturnUserMeeting(userName string) []Meetings {
	allMeeting := returnAllMeeting()
	var userMeeting []Meetings
	for _, value := range allMeeting {
		if userName == value.Initiator {
			userMeeting = append(userMeeting, value)
			continue
		}
		for _, value1 := range value.Participator {
			if userName == value1 {
				userMeeting = append(userMeeting, value)
				break
			}
		}
	}
	return userMeeting
}

//检查会议时间戳是否重叠
func CheckForTime(futureMeeting, alreadyMeeting Meetings) bool {
	if SmallDate(futureMeeting.ETime, alreadyMeeting.STime) || LargeDate(futureMeeting.STime, alreadyMeeting.ETime) {
		return true
	} else {
		return false
	}
}

func CheckForParticipator(Participator []string) error {
	for _, value := range Participator {
		if !User.IsUser(value) {
			return errors.New(value + " is not a user of the system")
		}
	}
	return nil
}

func CheckForMeeting(futureMeeting Meetings) error {
	if SmallDate(futureMeeting.ETime, futureMeeting.STime) {
		return errors.New("The meeting's ETime must larger than STime")
	}
	UserMeetings := ReturnUserMeeting(futureMeeting.Initiator)
	for _, value := range UserMeetings {
		if value.Title == futureMeeting.Title {
			return errors.New("The Initiator has one Meeting called " + value.Title + " that is conflict with this meeting on Title")
		}
		if !CheckForTime(futureMeeting, value) {
			return errors.New("The Initiator has one Meeting called " + value.Title + " that is conflict with this meeting on time")
		}
	}
	for _, value := range futureMeeting.Participator {
		partMeetings := ReturnUserMeeting(value)
		for _, value1 := range partMeetings {
			if !CheckForTime(futureMeeting, value1) {
				return errors.New("The Participator " + value + " has one Meeting called" + value1.Title + "that is conflict with this meeting")
			}
		}
	}
	return nil
}

func QueryMeetingByTitle(currentUser, Title string) (Meetings, error, int) {
	meeting := returnAllMeeting()
	var gg Meetings
	for index, value := range meeting {
		if value.Title == Title && value.Initiator == currentUser {
			return value, nil, index
		}
	}
	return gg, errors.New("The User has no this meeting"), 0
}

func AddOneMeeting(futureMeeting Meetings) error {
	err := CheckForParticipator(futureMeeting.Participator)
	if err != nil {
		return err
	}
	err1 := CheckForMeeting(futureMeeting)
	if err1 != nil {
		return err1
	}
	AllMeetings := returnAllMeeting()
	for _, value := range AllMeetings {
		if value.Title == futureMeeting.Title {
			return errors.New("The meeting's title has been used")
		}
	}
	AllMeetings = append(AllMeetings, futureMeeting)
	result, _ := json.Marshal(AllMeetings)
	file, _ := os.OpenFile("Meeting.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
	file.Write(result)
	fmt.Println("Create a meeting successfully")
	return nil
}

func AddParticipators(currentUser, Title string, NewParticipator []string) error {
	err := CheckForParticipator(NewParticipator)
	if err != nil {
		return err
	}
	InitiatorMeeting, err1, index := QueryMeetingByTitle(currentUser, Title)
	if err1 != nil {
		return err1
	}
	for _, value := range NewParticipator {
		if currentUser == value {
			return errors.New("The NewParticipators has one that is Initiator")
		}
		for _, value1 := range InitiatorMeeting.Participator {
			if value == value1 {
				return errors.New(value + " has entered this meeting")
			}
		}
		partMeetings := ReturnUserMeeting(value)
		for _, value1 := range partMeetings {
			if !CheckForTime(value1, InitiatorMeeting) {
				return errors.New("The Participator " + value + "has a meeting called " + value1.Title + " that is conflict with your meeting on time")
			}
		}
	}
	AllMeetings := returnAllMeeting()
	AllMeetings[index].Participator = append(AllMeetings[index].Participator[0:], NewParticipator[0:]...)
	result, _ := json.Marshal(AllMeetings)
	file, _ := os.OpenFile("Meeting.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
	file.Write(result)
	return nil
}

func DeleteParticipators(currentUser, Title string, NewParticipator []string) error {
	err := CheckForParticipator(NewParticipator)
	if err != nil {
		return err
	}
	_, err1, index := QueryMeetingByTitle(currentUser, Title)
	if err1 != nil {
		return err1
	}
	AllMeetings := returnAllMeeting()
	for _, value := range NewParticipator {
		for index1, value1 := range AllMeetings[index].Participator {
			if value == value1 {
				AllMeetings[index].Participator = append(AllMeetings[index].Participator[0:index1], AllMeetings[index].Participator[index1+1:]...)
				if len(AllMeetings[index].Participator) == 0 {
					AllMeetings = append(AllMeetings[0:index], AllMeetings[index+1:]...)
				}
				break
			}
			if index1 == len(AllMeetings[index].Participator)-1 {
				return errors.New("The meeting has no this participator called " + value)
			}
		}
	}
	result, _ := json.Marshal(AllMeetings)
	file, _ := os.OpenFile("Meeting.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
	file.Write(result)
	return nil
}

func DeleteMeetingByTitle(currentUser, Title string) error {
	AllMeetings := returnAllMeeting()
	for index, value := range AllMeetings {
		if value.Initiator == currentUser && value.Title == Title {
			AllMeetings = append(AllMeetings[0:index], AllMeetings[index+1:]...)
			result, _ := json.Marshal(AllMeetings)
			file, _ := os.OpenFile("Meeting.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
			file.Write(result)
			return nil
		}
	}
	return errors.New("The User has no this meeting")
}

func QuitMeetingByTitle(currentUser, Title string) error {
	AllMeetings := returnAllMeeting()
	for index, value := range AllMeetings {
		if value.Title == Title {
			for index1, value1 := range value.Participator {
				if currentUser == value1 {
					AllMeetings[index].Participator = append(AllMeetings[index].Participator[0:index1], AllMeetings[index].Participator[index1+1:]...)
					if len(AllMeetings[index].Participator) == 0 {
						AllMeetings = append(AllMeetings[0:index], AllMeetings[index+1:]...)
					}
					result, _ := json.Marshal(AllMeetings)
					file, _ := os.OpenFile("Meeting.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
					file.Write(result)
					return nil
				}
				if index1 == len(AllMeetings[index].Participator)-1 {
					return errors.New("You doesn't take part in this meeting")
				}
			}
		}
	}
	return errors.New("No meeting called " + Title)
}

func QueryMeetingByTime(currentUser string, STime, ETime time.Time) []Meetings {
	meetings := ReturnUserMeeting(currentUser)
	var result []Meetings
	for _, value := range meetings {
		if !(value.ETime.Before(STime) || value.STime.After(ETime)) {
			result = append(result, value)
		}
	}
	return result
}

func ClearAllUserMeeting(currentUser string) {
	AllMeetings := returnAllMeeting()
	for index, value := range AllMeetings {
		if value.Initiator == currentUser {
			AllMeetings = append(AllMeetings[0:index], AllMeetings[index+1:]...)
		}
	}
	result, _ := json.Marshal(AllMeetings)
	file, _ := os.OpenFile("Meeting.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
	file.Write(result)
}
