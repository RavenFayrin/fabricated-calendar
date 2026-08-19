package calendar

// import "fmt"

// type Date struct {
// 	Calendar *CalendarSystem
// 	Year     int
// 	Month    *Month
// 	Day      int
// }

// func NewDate(cal *CalendarSystem, year int, month interface{}, day int) (*Date, error) {
// 	m, err := cal.GetMonth(month)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Date{
// 		Calendar: cal,
// 		Year:     year,
// 		Month:    m,
// 		Day:      day,
// 	}, nil
// }

// func (d *Date) Equals(other *Date) bool {
// 	if other == nil {
// 		return false
// 	}

// 	return d.Calendar == other.Calendar &&
// 		d.Year == other.Year &&
// 		d.Month == other.Month &&
// 		d.Day == other.Day
// }

// func (cs *CalendarSystem) GetMonth(month interface{}) (*Month, error) {

// 	switch type_value := month.(type) {

// 	case string:
// 		for _, m := range cs.Months {
// 			if m.Name == type_value {
// 				return m, nil
// 			}
// 		}

// 	case int:
// 		for _, m := range cs.Months {
// 			if m.Order == type_value {
// 				return m, nil
// 			}
// 		}
// 	}

// 	return nil, fmt.Errorf("invalid month")
// }

// func (d *Date) AbsoluteDay() int {
// 	return d.Calendar.DateToAbsoluteDay(d)
// }

// func (d *Date) Before(other *Date) bool {
// 	return d.AbsoluteDay() < other.AbsoluteDay()
// }
