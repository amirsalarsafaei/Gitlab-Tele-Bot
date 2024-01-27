package merge

func getReviewersText(reviewers Reviewers) string {
	if len(reviewers) == 0 {
		return "No Reviewers"
	}

	res := ""
	for i, rev := range reviewers {
		res += rev.Name
		if i != len(reviewers)-1 {
			res += ", "
		}
	}

	return res
}
