package time

import "time"

// golang 格式化格式书写,不是yyyy-MM-dd这种写法
// 到秒的格式
const Format_LocalDateTimeSecond string = "2006-01-02 15:04:05"
// 到分的格式
const Format_LocalDateTimeMinute string = "2006-01-02 15:04"
//到天的格式
const Format_LocalDate = "2006-01-02"
//到天的数字表示格式如 20190605
const Format_LocalDateIntString = "20060102"



type LocalTimeSecond time.Time
type LocalTimeMinute time.Time
type LocalTimeDate time.Time


func (l LocalTimeSecond) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(Format_LocalDateTimeSecond)+2)
	b = append(b, '"')
	b = time.Time(l).AppendFormat(b, Format_LocalDateTimeSecond)
	b = append(b, '"')
	return b, nil
}

func (l *LocalTimeSecond) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"`+Format_LocalDateTimeSecond+`"`, string(b), time.Local)
	*l = LocalTimeSecond(now)
	return err
}


// --------
func (l LocalTimeMinute) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(Format_LocalDateTimeMinute)+2)
	b = append(b, '"')
	b = time.Time(l).AppendFormat(b, Format_LocalDateTimeMinute)
	b = append(b, '"')
	return b, nil
}

func (l *LocalTimeMinute) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"`+Format_LocalDateTimeMinute+`"`, string(b), time.Local)
	*l = LocalTimeMinute(now)
	return err
}


// --------

func (l LocalTimeDate) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(Format_LocalDate)+2)
	b = append(b, '"')
	b = time.Time(l).AppendFormat(b, Format_LocalDate)
	b = append(b, '"')
	return b, nil
}

func (l *LocalTimeDate) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"`+Format_LocalDate+`"`, string(b), time.Local)
	*l = LocalTimeDate(now)
	return err
}




