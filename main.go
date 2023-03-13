package main

import (
	"fmt"
	"github.com/incubator4/vtuber-calendar/pkg/config"
	"github.com/incubator4/vtuber-calendar/pkg/route"
)

func main() {
	fmt.Println(config.GlobalConfig)
	r := route.NewServer()

	_ = r.Run()
	//cal := ics.NewCalendar()
	//cal.SetMethod(ics.MethodRequest)
	//event := cal.AddEvent(fmt.Sprintf("id@domain", 1))
	//event.SetCreatedTime(time.Now())
	//event.SetDtStampTime(time.Now())
	//event.SetModifiedAt(time.Now())
	//event.SetStartAt(time.Now())
	//event.SetEndAt(time.Now())
	//event.SetSummary("Summary")
	//event.SetLocation("Address")
	//event.SetDescription("Description")
	//event.SetURL("https://URL/")
	//event.AddRrule(fmt.Sprintf("FREQ=YEARLY;BYMONTH=%d;BYMONTHDAY=%d", time.Now().Month(), time.Now().Day()))
	//event.SetOrganizer("sender@domain", ics.WithCN("This Machine"))
	//event.AddAttendee("reciever or participant", ics.CalendarUserTypeIndividual, ics.ParticipationStatusNeedsAction, ics.ParticipationRoleReqParticipant, ics.WithRSVP(true))
	//fmt.Println(cal.Serialize())
	//var calendars []pkg.Calendar
	//pkg.DB.Table("calendar").Find(&calendars, 3)
	//fmt.Println(calendars)

}
