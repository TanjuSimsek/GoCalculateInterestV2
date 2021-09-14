package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
)

type dataStruct struct {
	period string
}

type sample struct {
	data interface{}
}

func (s *sample) func1() {
	obj := map[string]interface{}{}
	// obj["foo1"] = "bar1"
	s.data = obj
}

func (s *sample) addField(key, value string) {
	v, _ := s.data.(map[string]interface{})
	v[key] = value
	s.data = v
}
func main() {

	s := &sample{}
	s.func1()
	//Son ödeme tarihi ve ödeme tarihi alınarak dönem parametresi alıp faiz hesabı yapılacak.
	//bu parametreler dışarıdan alınacak ve double bir değer dönecek.
	//amaç şu bir faturanın son ödeme tarihi var birde ödeme tarihi var.Mesela Faturanın son ödeme tarihi 30-04-2021
	//ödeme tarihi 16-07-2021 aradaki fark 77 gün.Burada 30 + 30 +17 gün devir olacak.
	// var expiredDate = "30-07-2020"
	var expiredDate = "31-10-2016"
	var deliverDate float64

	//son ödeme tarihine 1 ay ekle ve  devret

	// var paymentDate = "15-09-2020"
	var paymentDate = "01-02-2017"
	// var donem = "06-2018"
	//***yenı kod ---

	// var donem = "10-2016"
	// var interestRate = getData(donem)

	//***yenı kod ---
	// i, _ := strconv.ParseFloat(interestRate, 64)
	// log.Println(interestRate)
	// layout := "2006-01-02"
	layout := "02-01-2006"

	t, err := time.Parse(layout, expiredDate)

	var tt, mm = addOneMounth(t)

	var donem = ConvertPeriod(tt)
	var interestRate = getData(donem)
	//gelen son ödeme tarihinin başlangıç ve bitiş tarihini buldum
	log.Println(tt)
	log.Println(mm)
	t2, err := time.Parse(layout, paymentDate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
	//nekadarlık bir gecikme oldu onu buldum
	var expDay = calculateDate2(t, t2)
	if expDay < 28 {

		deliverDate = expDay
	} else {

		deliverDate = calculateDate(tt, mm)
	}
	log.Println(expDay)

	var firstControl = float64(expDay - deliverDate)
	log.Println(firstControl)
	log.Println(deliverDate)
	var resultCalculate = calculateIntrest(2850.63, 940.72, interestRate, deliverDate, firstControl)

	log.Println("---------ilk ay hesaplanan değer")
	log.Println(resultCalculate)
	log.Println("---------ilk ay hesaplanan değer")
	float64ToString := fmt.Sprint(resultCalculate)
	s.addField(tt.String(), float64ToString)
	var control = int(expDay - deliverDate)
	// var resultMountByMount = calculateOtherMount(resultCalculate, int(deliverDate), int(expDay))
	// log.Println(resultMountByMount)
	for control > 0 {
		// var firstTt = tt
		var r, tt2, d = controlAndCalculateOtherMount(control, tt, resultCalculate, deliverDate)
		float64ToString := fmt.Sprint(d)
		s.addField(tt2.String(), float64ToString)
		log.Println("------ikinci değer----")
		log.Println(d)
		log.Println("------ikinci değer----")
		control = r
		if control > 0 {

			// var secondTt = tt
			control, tt, resultCalculate = controlAndCalculateOtherMount(r, tt2, d, deliverDate)
			float64ToString := fmt.Sprint(resultCalculate)
			s.addField(tt.String(), float64ToString)
			log.Println("tttt")
			// control = control

		}

	}
	var sum float64 = 0
	for _, val := range s.data.(map[string]interface{}) {

		// log.Println(key)
		// log.Println(val)
		// log.Println(donem)
		var number = foo(val)
		sum += number
	}
	log.Println("Sonuc:")
	log.Println(sum)
	log.Println("-----------------------")
	log.Println(s)

	// getData("asd")
}
func ConvertPeriod(timePeriod time.Time) string {

	formatted := fmt.Sprintf("%02d-%02d",
		timePeriod.Month(), timePeriod.Year())
	// layout2 := "01-2006"
	// t3, _ := time.Parse(layout2, formatted)

	return formatted

}

func foo(veri interface{}) float64 {
	f64, _ := strconv.ParseFloat(veri.(string), 64)
	return f64 + 1
}

func calculateOtherMount(firstMountCalculate float64, dayOfMount int, delayMount int) float64 {

	var result = (firstMountCalculate / float64(dayOfMount)) * float64(delayMount)

	return result

}
func controlAndCalculateOtherMount(delayTimeMount int, firstMont time.Time, resultCalculate float64, baseDeliverDate float64) (int, time.Time, float64) {

	var tt, mm = addOneMounth(firstMont)
	var deliverDate = calculateDate(tt, mm)
	var resultMountByMount float64

	if delayTimeMount-int(deliverDate) > 0 {

		resultMountByMount = calculateOtherMount(resultCalculate, int(baseDeliverDate), int(deliverDate))

	} else {

		resultMountByMount = calculateOtherMount(resultCalculate, int(baseDeliverDate), delayTimeMount)
	}

	var control = int(float64(delayTimeMount) - deliverDate)

	return control, tt, resultMountByMount

}

func addOneMounth(date time.Time) (time.Time, time.Time) {

	log.Println(date)
	var after = date.AddDate(0, 1, 0)
	// var app = Bod(after)
	year, month, _ := after.Date()
	firstDayOfThisMonth := time.Date(year, month, 1, 0, 0, 0, 0, date.Location())

	endOfThisMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, date.Location())

	log.Println(firstDayOfThisMonth)
	log.Println(endOfThisMonth)

	return firstDayOfThisMonth, endOfThisMonth

}
func Bod(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

func getData(date string) string {

	log.Println(date)
	var result string = ""
	resp, _ := soup.Get("https://www.tcmb.gov.tr/wps/wcm/connect/TR/TCMB+TR/Main+Menu/Temel+Faaliyetler/Para+Politikasi/Reeskont+ve+Avans+Faiz+Oranlari")
	doc := soup.HTMLParse(resp)
	rows := doc.Find("table").FindAll("tr")[1:]

	for _, row := range rows {
		columns := row.FindAll("td")

		tarih := columns[0].Text()
		// iskonto := columns[1].Text()
		faiz := columns[2].Text()
		res2 := strings.ReplaceAll(faiz, "\u00a0", "")
		res3 := strings.Replace(tarih, ".", "-", -1)
		layout := "02-01-2006"

		t, err := time.Parse(layout, res3)
		if err != nil {
			fmt.Println(err)
		}

		// year, month, _ := t.Date()
		// log.Println(t)
		// log.Println(year)
		// log.Println(int(month))

		formatted := fmt.Sprintf("%02d-%02d",
			t.Month(), t.Year())
		// log.Println(formatted)
		layout2 := "01-2006"
		t2, _ := time.Parse(layout2, date)
		t3, _ := time.Parse(layout2, formatted)
		if t2.Unix() > t3.Unix() {

			log.Println(faiz)

			result = res2

		}

		// fmt.Println(tarih + "\t" + iskonto + "\t" + faiz)
		log.Println(result)
	}
	return result
}
func calculateDate2(expiredDate time.Time, paymentDate time.Time) float64 {

	// var result = -1.0
	log.Println(expiredDate)
	log.Println(paymentDate)
	days := paymentDate.Sub(expiredDate).Hours() / 24
	// if days >= 0 {

	// 	result = (days + 1)

	// }
	return days

}
func calculateDate(expiredDate time.Time, paymentDate time.Time) float64 {

	var result = -1.0
	log.Println(expiredDate)
	log.Println(paymentDate)
	days := paymentDate.Sub(expiredDate).Hours() / 24
	if days >= 0 {

		result = (days + 1)

	}
	return result

}
func calculateIntrest(aylik_ucret float64, devlete_iletilecek_vergiler float64, interestedRate string, date float64, firstControl float64) float64 {

	interestRateRep := strings.Replace(interestedRate, ",", ".", -1)

	f, _ := strconv.ParseFloat(interestRateRep, 64)
	var dayOfInterest = ((f + 5) / 360) / 100
	var a = ((aylik_ucret * date) * dayOfInterest)
	var b = devlete_iletilecek_vergiler * date * dayOfInterest

	var result = (a + b) * 1.28

	return result

}
