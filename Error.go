package timeparse

import ("fmt"
        "strconv"
	    )

type Time struct{
	hour,minute,second int
}
type TimeParseError struct{
	msg string
	input string
}

func (t *TimeParseError) Error() string{
	return fmt.Sprintf("%v : %v",t.msg,t.input)
}

func ParseTime(input string)(Time,error){
	components := string.split(input,":")
	if len(components) != 3{
		return Time{},&TimeParseError{"Invalid Number of time components",input}
	}else {
		hour,err := strconv.Atoi(components[0])
		if err != nil {
			return Time{},&TimeParseError(fmt.Sprintf("Error Parsing hour %v",err),input)
		}

		minute,err := strconv.Atoi(components[1])
		if err != nil {
			return Time{},&TimeParseError(fmt.Sprintf("Error Parsing minute %v",err),input)
		}

		seconds,err := strconv.Atoi(components[2])
		if err != nil {
			return Time{},&TimeParseError(fmt.Sprintf("Error Parsing seconds %v",err),input)
		}

		if hour > 23 || hour < 0 {
			return Time{0, 0, 0}, &TimeParseError{"hour out of range: 0 <= hour <= 23", fmt.Sprintf("%v", hour)}
		}
		if minute > 59 || minute < 0 {
			return Time{0, 0, 0}, &TimeParseError{"minute out of range: 0 <= minute <= 59", fmt.Sprintf("%v", minute)}
		}
		if second > 59 || second < 0 {
			return Time{0, 0, 0}, &TimeParseError{"second out of range: 0 <= second <= 59", fmt.Sprintf("%v", second)}
		}

		return Time{hour, minute, second}, nil
	}
}