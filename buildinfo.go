package buildinfo

// See docs for --workspace_status_command
var (
  _GIT_COMMIT_ID = ""
)

func GitCommitID() string {
  return _GIT_COMMIT_ID
}
