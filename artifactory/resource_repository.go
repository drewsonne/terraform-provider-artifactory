package artifactory

import (
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
)

var types,
	packageTypes,
	checksumPolicyTypes,
	snapshotVersionBehaviors,
	remoteRepoChecksumPolicyTypes,
	vcsGitProviders,
	vcsType,
	pomRepositoryReferencesCleanupPolicy []string

func init() {
	types = []string{"local", "remote", "virtual"}
	packageTypes = strings.Split("maven|gradle|ivy|sbt|nuget|gems|npm|bower|debian|composer|pypi|docker|vagrant|gitlfs|yum|conan|generic", "|")
	checksumPolicyTypes = []string{"client-checksums", "server-generated-checksums"}
	snapshotVersionBehaviors = []string{"unique", "non-unique", "deployer"}
	remoteRepoChecksumPolicyTypes = []string{"", "generate-if-absent", "fail", "ignore-and-generate", "pass-thru"}
	vcsType = []string{"", "git"}
	vcsGitProviders = []string{"", "github", "bitbucket", "stash", "artifactory", "custom"}
	pomRepositoryReferencesCleanupPolicy = []string{"discard_active_reference", "discard_any_reference", "nothing"}
}

type retryFunc func() error

func retry(f retryFunc) error {
	return resource.Retry(5*time.Minute, func() *resource.RetryError {
		err := f()
		if err != nil {
			return resource.RetryableError(err)
		}
		return nil
	})

}
