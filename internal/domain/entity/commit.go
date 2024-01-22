package entity

type CommitEntity struct {
	CommitType string
	Scope      string
	ShortDesc  string
	LongDesc   string
	Issue      string
}

func (c CommitEntity) Formatter() string {
	var commit string
	commit = commit + c.CommitType
	if c.Scope != "" {
		commit = commit + "(" + c.Scope + ")"
	}
	commit = commit + ": " + c.ShortDesc
	if c.LongDesc != "" {
		commit = commit + "\n\n" + c.LongDesc
	}
	if c.Issue != "" {
		commit = commit + "\n\n" + "#" + c.Issue
	}
	return commit
}
