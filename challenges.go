package fortnite

type Challenges struct {
	Items []Challenge `json:"items"`
}

type Challenge struct {
	Metadata []Metadata `json:"metadata"`
}

type Metadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func GetChallenges(s *Service) (result Challenges, err error) {
	_, err = s.Do("challenges", result)
	return
}
