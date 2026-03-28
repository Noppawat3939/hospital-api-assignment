package mock

func FindPatientByID(id string) (*MockPatient, bool) {
	var result *MockPatient

	for _, p := range MockPatients {
		if p.PassportID == id || p.NationalID == id {
			result = &p
			break
		}
	}

	if result != nil {
		return result, true
	}

	return nil, false
}
