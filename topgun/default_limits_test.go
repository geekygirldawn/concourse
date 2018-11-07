package topgun_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("an ATC with default resource limits set", func() {
	BeforeEach(func() {
		Deploy("deployments/concourse.yml",
			"-o", "operations/default_limits.yml",
			"-v", "default_task_cpu_limit=512",
			"-v", "default_task_memory_limit=1GB",
		)
	})

	It("respects the default resource limits, overridding when specified", func() {
		buildSession := fly.Spawn("execute", "-c", "tasks/tiny.yml")
		<-buildSession.Exited

		hijackSession := fly.Spawn(
			"hijack",
			"-b", "1",
			"--", "sh", "-c",
			"cat /sys/fs/cgroup/memory/memory.memsw.limit_in_bytes; cat /sys/fs/cgroup/cpu/cpu.shares",
		)
		<-hijackSession.Exited

		Expect(hijackSession.ExitCode()).To(Equal(0))
		Expect(hijackSession).To(gbytes.Say("1073741824\n512"))

		buildSession = fly.Spawn("execute", "-c", "tasks/limits.yml")
		<-buildSession.Exited

		hijackSession = fly.Spawn(
			"hijack",
			"-b", "2",
			"--", "sh", "-c",
			"cat /sys/fs/cgroup/memory/memory.memsw.limit_in_bytes; cat /sys/fs/cgroup/cpu/cpu.shares",
		)
		<-hijackSession.Exited

		Expect(hijackSession.ExitCode()).To(Equal(0))
		Expect(hijackSession).To(gbytes.Say("104857600\n256"))
	})
})
