package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Date struct {
	Date string
}
type DateInterface struct {
	Date map[string]int
}

var date Date

func rightjust(s string, n int, fill string) string {
	return strings.Repeat(fill, n) + s
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainPage)
	mux.HandleFunc("/info", infoPage)
	mux.HandleFunc("/process", processForm)
	fmt.Println("Server start listening at :8000")
	log.Fatal(http.ListenAndServe(":8000", mux))

}

func mainPage(w http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := tmpl.Execute(w, err); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func processForm(w http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	str := request.FormValue("dateofbirth")
	if str == "" {
		http.Error(w, "NO DATE", 400)
		return
	}
	date.Date = getInfo(str)
	http.Redirect(w, request, "/info", http.StatusSeeOther)
}

func infoPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/info.html")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	err = tmpl.Execute(w, map[string]interface{}{
		"items":       viewInfo(date.Date),
		"information": viewDescription(vI(date.Date)),
	})
	if err != nil {
		log.Fatal(err)
		return
	}
}

func getInfo(date string) string {
	date_splited := strings.Split(date, "-")

	day_splited := strings.Split(date_splited[2], "")
	month_splited := strings.Split(date_splited[1], "")
	year_splited := strings.Split(date_splited[0], "")

	sum_day_month := 0
	sum_year := 0

	for i := 0; i < 2; i++ {
		value, _ := strconv.Atoi(day_splited[i])
		sum_day_month += value
	}
	for i := 0; i < 2; i++ {
		value, _ := strconv.Atoi(month_splited[i])
		sum_day_month += value
	}
	for i := 0; i < 4; i++ {
		value, _ := strconv.Atoi(year_splited[i])
		sum_year += value
	}

	first_work := sum_day_month + sum_year

	fw_str := strconv.FormatInt(int64(first_work), 10)
	if len(fw_str) != 2 {
		fw_str = rightjust(fw_str, 1, "0")
	}
	fw_str_list := strings.Split(fw_str, "")
	second_work := 0
	for i := 0; i < 2; i++ {
		value, _ := strconv.Atoi(fw_str_list[i])
		second_work += value
	}
	tmp, _ := strconv.Atoi(day_splited[0])
	third_work := second_work - 2*tmp
	fourth_work := (third_work / 10) + third_work%10

	sd_str := strconv.FormatInt(int64(second_work), 10)
	if len(sd_str) != 2 {
		sd_str = rightjust(sd_str, 1, "0")
	}
	trd_str := strconv.FormatInt(int64(third_work), 10)
	if len(trd_str) != 2 {
		trd_str = rightjust(trd_str, 1, "0")
	}
	frt_str := strconv.FormatInt(int64(fourth_work), 10)
	if len(frt_str) != 2 {
		frt_str = rightjust(frt_str, 1, "0")
	}

	total := fw_str + sd_str + trd_str + frt_str + date

	return total
}

func viewInfo(total string) map[string]string {
	ans := make(map[string]string)
	for i := 1; i < 10; i++ {
		i_key := strconv.FormatInt(int64(i), 10)
		cnt := strings.Count(total, i_key)
		ans[i_key] = strings.Repeat(i_key, cnt)
	}
	return ans
}

func vI(total string) map[string]int {
	ans := make(map[string]int)
	for i := 1; i < 10; i++ {
		i_key := strconv.FormatInt(int64(i), 10)
		cnt := strings.Count(total, i_key)
		ans[i_key] = cnt
	}
	return ans
}

func viewDescription(dict map[string]int) map[string]string {
	ans := make(map[string]string)
	for key, value := range dict {
		switch key {
		case "1":
			switch value {
			case 0:
				ans[key] = "Встречается только у людей, рожденных после 2000 года. Считает, что, мир крутится вокруг него и все ему чем-то обязаны. Важно еще в раннем детстве избавлять его от ощущения своей неповторимости и особенности, предпочтительнее воспитание в коллективе"
			case 1:
				ans[key] = "Перед нами типичный эгоист. Свои интересы ставит превыше всего"
			case 2:
				ans[key] = "Близок к типичному эгоисту. Самодовольство и самовосхваление"
			case 3:
				ans[key] = "Золотая середина. Уравновешенный человек"
			case 4:
				ans[key] = "Сильный характер. Решительность и действенность"
			case 5:
				ans[key] = "Диктатор"
			case 6:
				ans[key] = "Достаточно жесток, сложный характер. Но способен совершить благородный поступок"
			}
		case "2":
			switch value {
			case 0:
				ans[key] = "Это энергетические вампиры в хорошем смысле слова. Они открыты для общения, подпитываются от окружающих новыми идеями и впечатлениями"
			case 1:
				ans[key] = "Достаточно своей энергии, однако для поддержания баланса рекомендуются занятия спортом"
			case 2:
				ans[key] = "Золотая середина. Баланс и гармония. Своей энергией с радостью делится с окружающими"
			case 3:
				ans[key] = "Из него может получиться неплохой экстрасенс. Есть смысл развивать интуицию"
			case 4:
				ans[key] = "Пользуется популярностью у противоположного пола за счет своей активности и непосредственности. Особенно привлекателен для людей с тремя шестерками в квадрате"
			}

		case "3":
			switch value {
			case 0:
				ans[key] = "Он пунктуален и очень любит чистоту"
			case 1:
				ans[key] = "Уборка по настроению. У него может быть как идеальный порядок, так и хаос. Причем как в доме, так и в голове"
			case 2:
				ans[key] = "Перед нами ученый. Или тот, кто им точно мог бы стать. Точные науки интересуют и даются легко"
			case 3:
				ans[key] = "Желание есть, а возможностей мало. Как бы не хотел углубиться в изучение формул, не получается"
			}

		case "4":
			switch value {
			case 0:
				ans[key] = "Слабый, подвержен болезням,  особенно если в матрице много двоек"
			case 1:
				ans[key] = "Здоровье среднестатистического человека. Болеет, но не часто"
			case 2:
				ans[key] = "Крепкий орешек. Он не простудится на сквозняке, да и любая хворь быстро пройдет. Запас здоровья сказывается на сексуальной активности"
			case 3:
				ans[key] = "Третья четверка является дополнительным бонусом ко всему тому, о чем мы написано в предыдущем пункте"
			}

		case "5":
			switch value {
			case 0:
				ans[key] = "Его мозг всегда в работе. Он что-то придумывает, что-то доказывает, в том числе и самому себе. Методом проб и ошибок достигает поставленной цели"
			case 1:
				ans[key] = "Интуиция развита. В жизни все дается легко"

			case 2:
				ans[key] = "Хорошая интуиция. Главное к ней прислушиваться. Часто из обладателей двух пятерок получаются грамотные следователи и юристы"
			case 3:
				ans[key] = "Совершить ошибку для них редкость, ведь они почти ясновидящие. Судьба постоянно посылает им знаки свыше. А они умело пользуются подсказками "
			case 4:
				ans[key] = "У этих людей есть все шансы сделать экстрасенсорику своей профессией. Они способны видеть будущее. Им снятся вещие сны, а также быстро приходят ответы на интересующие вопросы "
			}

		case "6":
			switch value {
			case 0:
				ans[key] = "Предназначение человека — физический труд. Нередко про таких говорят «золотые руки». Из них получаются неплохие ремесленники, плотники, маляры. Единственное, занятие физическим трудом не всегда приносит им удовольствие "
			case 1:
				ans[key] = "Обладателю одной шестерки делать что-то своими руками необходимо, однако ему предоставляется возможность найти себя и в другой сфере деятельности"
			case 2:
				ans[key] = "В этом случае заниматься физическим трудом не обязательно, но обладатели двух шестерок любят мастерить что-то своими руками. Однако в данном случае это скорее хобби, чем профессия"
			case 3:
				ans[key] = "На них можно рассчитывать. Добросовестно выполняют свою работу. Но им необходима подпитка. Чаще всего энергию берут у партнера с большим количеством двоек в матрице"
			case 4:
				ans[key] = " Настоящий работяга. Ему сложно сидеть без дела. Всегда должен чем-то заниматься, при этом получает от этого огромное удовольствие. Если в квадрате еще и девятки, есть смысл получить высшее образование, уравновесив чрезмерную активность"
			}

		case "7":
			switch value {
			case 0:
				ans[key] = "Жизнь будет сложной, всего придется добиваться собственным трудом, учиться на своих ошибках. Не исключено отрешение от мирской жизни и уход в религию"
			case 1:
				ans[key] = "Жизнь протекает легко. Но ярких ее проявлений и крутых поворотов в судьбе ждать не стоит"

			case 2:
				ans[key] = "Они талантливы. Есть возможность стать известным художником или музыкантом. Однако, если не развиваться в профессии или творчестве, плюс может легко поменяться на минус. Ведь этим людям с легкостью дается не только хорошее, но и плохое"
			case 3:
				ans[key] = "Насыщенная яркими событиями жизнь. Нередко увлекаются экстремальными видами спорта. Риск не пугает их, а скорее манит. Однако есть смысл периодически сбавлять обороты и давать себе передышку"
			case 4:
				ans[key] = "Чаще всего обладатели четырех семерок не приспособлены к жизни. Увы, их жизненный путь недолог. И даже если они остаются на земле, их постоянно преследуют болезни"
			}

		case "8":
			switch value {
			case 0:
				ans[key] = "Не ждите от него пожертвований. Он с удовольствием будет принимать помощь и поддержку со стороны. А вот ставить интересы других выше своих собственных вряд ли сможет"
			case 1:
				ans[key] = "На него можно положиться. Всегда готов оказать поддержку"

			case 2:
				ans[key] = "Ощущение, что желание помогать заложено у него в генах. Стоит только попросить, как он бросит свои дела и уделит вам внимание"
			case 3:
				ans[key] = "Обладатель трех восьмерок охватывает заботой не только своих близких и друзей, но и всех нуждающихся. Неравнодушен к страданию народа, готов пойти на жертвы, чтобы добиться справедливости "
			case 4:
				ans[key] = "Перед нами талантливый психолог, которому не надоедает выслушивать, проявлять заботу и милосердие. В то же время обладает тягой к точным наукам"
			}

		case "9":
			switch value {
			case 0:
				ans[key] = "Отсутствие девятки может быть только у появившихся на свет после 2000 года. С самого рождения надо начинать развивать память и логику, чтобы уже к школе восполнить пустующий квадрат "
			case 1:
				ans[key] = "Заставляйте свой мозг постоянно работать. Разгадывайте кроссворды, учите стихи. Если второй девятки нет в матрице, ее надо отрабатывать"
			case 2:
				ans[key] = "У обладателей двух девяток с умственными способностями все в порядке. Есть все данные для успешного развития. Но необходимо постоянно работать, просто так ничего не дается "
			case 3:
				ans[key] = "А вот здесь запаса ума хватает на то, чтобы прилагать минимум усилий для хорошей учебы и успешной карьеры. Таким людям все дается легко"
			case 4:
				ans[key] = "Им открывается истина. Они практически на сто процентов защищены от провала в любом начинании. Возможно потому, что для них все очень просто, они жестоки, неприятны в общении, смотрят на всех свысока и часто бывают неадекватны."
			}

		}
	}

	return ans
}
