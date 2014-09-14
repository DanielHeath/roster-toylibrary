package web

import (
	"encoding/csv"
	"fmt"
)

const (
	TimeFormat     = "3:04PM"
	DateFormat     = "Jan 02 2006"
	TimeDateFormat = DateFormat + " " + TimeFormat
)

type CsvFileContext struct {
}

type Member struct {
	MemberNo  string
	Name      string
	Mobiles   []string
	Emails    []string
	HoursOwed float32
}

type Shift struct {
	Start             time.Time
	Duration          time.Duration
	Location          string
	Position          string
	Member            *Member
	EmailReminderSent time.Time
	SmsReminderSent   time.Time
}

var rosterHeaders = []string{
	"date",
	"time",
	"timezone",
	"duration",
	"location",
	"role",
	"member_id",
	"email_reminder_sent_date",
	"email_reminder_sent_date",
}

var memberHeaders = []string{
	"memberNo",
	"name",
	"mobiles",
	"emails",
	"hours_owed",
}

func checkHeadingRow(r *csv.Reader, description string, expected []string) error {
	headerLine, err := r.Read()
	if err != nil {
		return err
	}
	for i := range headerLine {
		if headerLine[i] != expected[i] {
			return fmt.Errorf("Column %d of %d CSV is %s, needs to be %s", i, description, headerLine[i], rosterHeaders[i])
		}
	}
	return nil
}

func (c *CsvFileContext) Read(members io.Reader, shifts io.Reader) ([]Members, []Shifts, error) {
	membersCsv := csv.NewReader(members)
	shiftsCsv := csv.NewReader(shifts)

	err := checkHeadingRow(membersCsv, "members", memberHeaders)
	if err != nil {
		return nil, nil, err
	}
	err = checkHeadingRow(shiftsCsv, "roster", rosterHeaders)
	if err != nil {
		return nil, nil, err
	}

}

func parseMembers(r *csv.Reader) ([]Member, error) {
	memberRows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	members := make([]Member, len(memberRows))
	for i, row := range memberRows {
		members[i].MemberNo = row[0]
		members[i].Name = row[1]
		members[i].Emails = row[2]
		members[i].Mobiles = row[3]
		members[i].HoursOwed = row[4]
	}
	return members, nil
}

func parseRoster(r *csv.Reader, members []Member) ([]Shift, error) {
	shiftRows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	shifts := make([]Shift, len(shiftRows))
	for i, row := range shiftRows {
		shifts[i].Start = row[0]
		shifts[i].Duration = row[0]
		shifts[i].Location = row[0]
		shifts[i].Position = row[0]
		shifts[i].Member = row[0]
		shifts[i].EmailReminderSent = row[0]
		shifts[i].SmsReminderSent = row[0]
	}
	return members, nil
}
