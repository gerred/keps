package sigs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"net/http"

	"github.com/go-yaml/yaml"

	"github.com/calebamiles/keps/pkg/sigs"
)

var _ = Describe("the SIGs helper package", func() {
	Describe("All()", func() {
		It("contains each SIG specified in sigs.yaml of kubernetes/community@master", func() {
			upstreamList := fetchUpstreamSIGNames()
			compiledList := sigs.All()

			Expect(len(upstreamList)).To(Equal(len(compiledList)), "compiled list of SIGs has different length than upstream")

			for _, s := range upstreamList {
				Expect(compiledList).To(ContainElement(s))
			}
		})
	})

	Describe("Exists()", func() {
		It("returns whether a specific SIG exists", func() {
			upstreamList := fetchUpstreamSIGNames()

			for _, s := range upstreamList {
				Expect(sigs.Exists(s)).To(BeTrue())
			}
		})
	})

	Describe("ExtractNameFromPath()", func() {
		Context("when the SIG exists", func() {
			It("extracts the SIG name from a given path", func() {
				p := "content/sig-node/kubelet"
				Expect(sigs.ExtractNameFromPath(p)).To(Equal("sig-node"))
			})
		})

		Context("when the SIG does not exist", func() {
			It("returns an empty string", func() {
				p := "content/sig-not-real/kubelet"
				Expect(sigs.ExtractNameFromPath(p)).To(BeEmpty())
			})
		})

		Context("when there are multiple SIG names in a given path", func() {
			It("returns an empty string", func() {
				p := "content/sig-node/sig-architecture/kublet"
				Expect(sigs.ExtractNameFromPath(p)).To(BeEmpty())
			})
		})
	})

	Describe("ExtractSubprojectNameFromPath()", func() {
		Context("when the subproject is immediately nested under the SIG", func() {
			It("extracts the subproject name from a given path", func() {
				p := "content/sig-node/kubelet"
				Expect(sigs.ExtractSubprojectNameFromPath(p)).To(Equal("kubelet"))
			})

		})

		Context("when the subproject is not immediately nested under the SIG", func() {
			It("returns an empty string", func() {
				p := "content/sig-node/some-dir/kubelet"
				Expect(sigs.ExtractSubprojectNameFromPath(p)).To(BeEmpty())
			})
		})

		Context("when the subproject is not owned by the SIG it is nested under", func() {
			It("returns an empty string", func() {
				p := "content/sig-networking/kubelet"
				Expect(sigs.ExtractSubprojectNameFromPath(p)).To(BeEmpty())
			})
		})
	})
})

func fetchUpstreamSIGNames() []string {
	resp, err := http.Get(upstreamSIGListURL)
	defer resp.Body.Close()

	Expect(err).ToNot(HaveOccurred(), "downloading canonical SIG list")

	respBytes, err := ioutil.ReadAll(resp.Body)
	Expect(err).ToNot(HaveOccurred(), "reading HTTP response")

	sl := &upstreamSIGList{}
	err = yaml.Unmarshal(respBytes, sl)
	Expect(err).ToNot(HaveOccurred(), "unmarshalling SIG list")
	Expect(sl.SIGs).ToNot(BeEmpty(), "unmarshalled SIG list is empty")

	names := []string{}
	for i := range sl.SIGs {
		names = append(names, sl.SIGs[i].Name)
	}

	return names
}

type upstreamSIGList struct {
	SIGs []upstreamSIGEntry `yaml:"sigs"`
}

type upstreamSIGEntry struct {
	Name string `yaml:"dir"` // we actually want to look at what the SIG is called on disk
}

const upstreamSIGListURL = "https://raw.githubusercontent.com/kubernetes/community/master/sigs.yaml"
