package organizers

var (
	VacancyIdMap = toIdMap(VacancyUrlMap)
)

func toIdMap(s map[string]int64) map[int64]string {
	m := make(map[int64]string, len(s))
	for k, v := range s {
		m[v] = k
	}
	return m
}
