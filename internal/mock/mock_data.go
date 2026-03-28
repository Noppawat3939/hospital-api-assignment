package mock

type MockPatient struct {
	FirstNameTH  string  `json:"first_name_th"`
	MiddleNameTH *string `json:"middle_name_th"`
	LastNameTH   string  `json:"last_name_th"`
	FirstNameEN  string  `json:"first_name_en"`
	MiddleNameEN *string `json:"middle_name_en"`
	LastNameEN   string  `json:"last_name_en"`
	DateOfBirth  string  `json:"date_of_birth"`
	PatientHN    string  `json:"patient_hn"`
	NationalID   string  `json:"national_id"`
	PassportID   string  `json:"passport_id"`
	PhoneNumber  *string `json:"phone_number"`
	Email        *string `json:"email"`
	Gender       string  `json:"gender"`
}

var MockPatients = []MockPatient{
	{
		FirstNameTH: "สมหญิง",
		LastNameTH:  "สวยงาม",
		FirstNameEN: "Somying",
		LastNameEN:  "Suayngam",
		DateOfBirth: "1995-05-10",
		PatientHN:   "HN100001",
		NationalID:  "1111111111111",
		PassportID:  "1111111111111",
		Gender:      "F",
	},
	{
		FirstNameTH: "สมชาย",
		LastNameTH:  "ใจดี",
		FirstNameEN: "Somchai",
		LastNameEN:  "Jaidee",
		DateOfBirth: "1990-01-01",
		PatientHN:   "HN100002",
		NationalID:  "2222222222222",
		PassportID:  "2222222222222",
		Gender:      "M",
	},
	{
		FirstNameTH: "สมหมาย",
		LastNameTH:  "แข็งแรง",
		FirstNameEN: "Sommai",
		LastNameEN:  "Khaengraeng",
		DateOfBirth: "1985-03-15",
		PatientHN:   "HN100003",
		NationalID:  "3333333333333",
		PassportID:  "3333333333333",
		Gender:      "M",
	},
	{
		FirstNameTH: "สมใจ",
		LastNameTH:  "สดใส",
		FirstNameEN: "Somjai",
		LastNameEN:  "Sodsai",
		DateOfBirth: "1992-07-20",
		PatientHN:   "HN100004",
		NationalID:  "4444444444444",
		PassportID:  "4444444444444",
		Gender:      "F",
	},
	{
		FirstNameTH: "สมปอง",
		LastNameTH:  "มีสุข",
		FirstNameEN: "Sompong",
		LastNameEN:  "Meesuk",
		DateOfBirth: "1980-12-05",
		PatientHN:   "HN100005",
		NationalID:  "5555555555555",
		PassportID:  "5555555555555",
		Gender:      "M",
	},
	{
		FirstNameTH: "อรทัย",
		LastNameTH:  "งามดี",
		FirstNameEN: "Onthay",
		LastNameEN:  "Ngamdee",
		DateOfBirth: "1993-04-11",
		PatientHN:   "HN100006",
		NationalID:  "6666666666666",
		PassportID:  "6666666666666",
		Gender:      "F",
	},
	{
		FirstNameTH: "วีระ",
		LastNameTH:  "ชัย",
		FirstNameEN: "Weera",
		LastNameEN:  "Chai",
		DateOfBirth: "1987-09-09",
		PatientHN:   "HN100007",
		NationalID:  "7777777777777",
		PassportID:  "7777777777777",
		Gender:      "M",
	},
	{
		FirstNameTH: "กมล",
		LastNameTH:  "ดีมาก",
		FirstNameEN: "Kamon",
		LastNameEN:  "Deemark",
		DateOfBirth: "1998-02-14",
		PatientHN:   "HN100008",
		NationalID:  "8888888888888",
		PassportID:  "8888888888888",
		Gender:      "F",
	},
	{
		FirstNameTH: "ธนา",
		LastNameTH:  "รุ่งเรือง",
		FirstNameEN: "Tana",
		LastNameEN:  "Rungreang",
		DateOfBirth: "1983-06-30",
		PatientHN:   "HN100009",
		NationalID:  "9999999999999",
		PassportID:  "9999999999999",
		Gender:      "M",
	},
	{
		FirstNameTH: "จันทร์เพ็ญ",
		LastNameTH:  "สว่าง",
		FirstNameEN: "Janpen",
		LastNameEN:  "Sawang",
		DateOfBirth: "1996-10-21",
		PatientHN:   "HN100010",
		NationalID:  "1010101010101",
		PassportID:  "1010101010101",
		Gender:      "F",
	},
}
