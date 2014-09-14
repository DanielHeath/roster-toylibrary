package web

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const (
	TimeFormat     = "3:04PM"
	TimeZone       = "Australia/Melbourne"
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

func (c *CsvFileContext) Read(members io.Reader, shifts io.Reader) ([]Member, []Shift, error) {
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
	m, err := parseMembers(membersCsv)
	if err != nil {
		return nil, nil, err
	}
	s, err := parseRoster(shiftsCsv, m)
	if err != nil {
		return nil, nil, err
	}
	return m, s, nil
}

func parseMembers(r *csv.Reader) ([]Member, error) {
	memberRows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	members := make([]Member, len(memberRows))
	for i, row := range memberRows {
		hours, err := strconv.ParseFloat(row[4], 32)
		if err != nil {
			return nil, fmt.Errorf("Line %d: Hours owed: %s", err)
		}
		members[i].MemberNo = row[0]
		members[i].Name = row[1]
		members[i].Emails = strings.Split(row[2], " ")
		members[i].Mobiles = strings.Split(row[3], " ")
		members[i].HoursOwed = float32(hours)
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
		date := row[0]
		time := row[1]
		zone, err := time.LoadLocation(row[2])
		if err != nil {
			return nil, fmt.Errorf("Timezone error in roster CSV: %s", err)
		}
		t := time.ParseInLocation(TimeDateFormat, date+" "+time, zone)
		shifts[i].Start = t

		shifts[i].Duration = row[2]
		shifts[i].Location = row[3]
		shifts[i].Position = row[4]
		shifts[i].Member = row[5]
		shifts[i].EmailReminderSent = row[6]
		shifts[i].SmsReminderSent = row[7]
	}
	return members, nil
}
