package merge

func getReviewersText(event Event) string {
	if len(event.Reviewers) == 0 {
		return "No Reviewers"
	}

	res := ""
	for i, rev := range event.Reviewers {
		res += rev.Name
		if i != len(event.Reviewers)-1 {
			res += ", "
		}
	}

	return res
}
