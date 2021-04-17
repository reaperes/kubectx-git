package git

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_git(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "git test")
}

var _ = Describe("git", func() {
	var (
		ts *httptest.Server
	)

	BeforeSuite(func() {
		ts = createMockGithubServer()
	})

	AfterSuite(func() {
		ts.Close()
	})

	Describe("Fetch", func() {
		Context("file exists", func() {
			It("fetch should not throw error", func() {
				git := createGit("secret")
				err := git.Fetch(fmt.Sprintf("%s/config.yaml", ts.URL))
				Expect(err).NotTo(HaveOccurred())
			})

			Context("invalid access token", func() {
				It("fetch failed", func() {
					git := createGit("invalid")
					err := git.Fetch(fmt.Sprintf("%s/config.yaml", ts.URL))
					Expect(err).Should(HaveOccurred())
				})
			})
		})

		Context("file not exists", func() {
			It("fetch should throw error", func() {
				git := createGit("secret")
				err := git.Fetch(fmt.Sprintf("%s/not-exists.yaml", ts.URL))
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})

func createGit(accessToken string) Git {
	if accessToken == "" {
		Fail("ACCESS_TOKEN is not defined")
	}
	return NewGit(accessToken)
}

func createMockGithubServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" && r.URL.Path == "/config.yaml" {
			auth := r.Header.Get("Authorization")
			// aW52YWxpZDo=: base64 encoded string("secret:")
			if auth != "Basic c2VjcmV0Og==" {
				w.WriteHeader(404)
				return
			}
			w.Header().Set("content-type", "text/plain; charset=utf-8")
			w.Write([]byte(`apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN5RENDQWJDZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRJeE1EUXhOREF3TWpnek5sb1hEVE14TURReE1qQXdNamd6Tmxvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTmhICjViSlNhNWR2RlVnSkFrUVlnbTdVUG9FWGNZTXJYY0xaT0RVNlZZcWJVWHBhMUUzUElzc3pLTlhlVUxCbFRic1AKYVJDbHVXSkJuUmU2RzR5d2NEMXJwa0QrTlUwK2Q2ajlRdjlQK0ozeGpyemlwbzlTd1ZmVkFDeTZiQ2F0eGlHUgoyWVY3ZjlQQ0MvNCtya2R6WTNyRC80VXk5L0VlMzdhVkxPNkZrSmhFbjhhdHdUeGdlK3dJOU1VZ1BmKzdpcWdRCld6WkMxT1RSVXpMVjZ0YnR3NDVEc2VBa05ZNFZXVU9pWjhPcFYwVjc5bWpJT1Q2SHpvR0JyQlVjSHQvQkpNTVMKY0RiN3JFaElYSTQ2OTJmWEJlZGNwamc5Q3lySHdxMU1BdWdiUG5CdjdaUUJPMDVmQnZCbXQ2MG9iWHU2d0FzSApiZ0VsdnhldlpYTnRKeVMxN3FjQ0F3RUFBYU1qTUNFd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFOS01hbU40bHBIYzVRUEpyRmwxUW95QTFFeUEKeW5FdVVzd0NsZmU1K0dib1hSeEVCMldFSmhTT2N0SHZNS2t4TVVFTzQzSzgycUNnUkxjbGNpcGx3Zk9YVU5qLwp3RzBGcE8zMXkrQWVmVTNMODAxSlN0bld3cktVbytwVURTaG1kWVZvY3EwRTRqaTdNdFpJOHB4TndzeWRGVEE1CllLSTVZU1dOeDVMbmRrV2xNWEQ1V09IeWFBVlhVUGdSSUpreVFSekNYOXl4NTM3R2RNVEVxSWRDa0xNM1RMSUsKcXFyd1lwNTdjbjQ3NTB4MHJKeU56aGtsLzlmdkhwaFloRURGQ2d0NGwyT3lFenNYdWRSeDlIek1CSm9CYkVtNwozamM0ZS9MMjZsUm0yakpOWlVoVzRlbkc0bnM2QmlFMmpsT016S0YwSkZRWUdFcVUrVXpzYTA1T3lrcz0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
    server: https://127.0.0.1:55657
  name: kind-kind
contexts:
- context:
    cluster: kind-kind
    namespace: default
    user: kind-kind
  name: kind-kind
current-context: kind-kind
kind: Config
preferences: {}
users:
- name: kind-kind
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM4akNDQWRxZ0F3SUJBZ0lJUE9BeGNleFB0MEV3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TVRBME1UUXdNREk0TXpaYUZ3MHlNakEwTVRRd01ESTRNemxhTURReApGekFWQmdOVkJBb1REbk41YzNSbGJUcHRZWE4wWlhKek1Sa3dGd1lEVlFRREV4QnJkV0psY201bGRHVnpMV0ZrCmJXbHVNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQXhWTUxGTnJXM2I5Tkl1SG4KOTI3VXhtaTF2Y1pHZGZJciszdW9XZjNJQU1LZ3Z0a3hrbVpsY3BQVE9zcGhlZkl6QUJGSDNDSmUvUW5adkk1dgowTVlhbzBMOXlJS1JPcEpJVDN2NGVxdGdGK0xyejdJNDJtT2htRG84UGVKNkN6SCtoU0VXUWNYbmtEcitsNmV6CkhmUFdoWHFFcjNoRXEyTXdia3pTWXFqdHBZMFk5K1NXNUNtSVNFSzRoK2NzT3lGZkVwb3NPMUlPNlBlR3VEdUgKWXVZVVY3NUVqSzV6bktBUmRvaWg2UHJCeU0wNVNHNm4xVTZaSHBZN21Ua3ErQ0VKMTY1NlgwNlpoaFZyVm5YMgo1Mk1mLzVSZmsxcGpTWGNzT3V6dEdBaWNKMlp6TlhMUWs1WGw2RzJwN3pSaU1mckhCVzJqcTZNdEo4SlR6QjdKCmFzZDlKd0lEQVFBQm95Y3dKVEFPQmdOVkhROEJBZjhFQkFNQ0JhQXdFd1lEVlIwbEJBd3dDZ1lJS3dZQkJRVUgKQXdJd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFNc1FndytDVW05WGFLYmpnUHpmM3EvdFJhUU9kUTNQSEh1UQpKcWU1NG93YkF1Z2lBUlR0aWRoMW1NYTlhWHVWVUNPM0d1enJmMXl2N3NEZ0RIeFh6aE1zWVNpZWRGUGNTOUh4CkF0Ykk5L3VsdjlCVVBCTi9CMmhpU1orWU1ENEhQbTdjd1YrbVEzTi9zT0lrd3l5dkVlWHBzRFp5UWc5czhNMy8KQVNORGQyRlhzTEsxa0pmM0tMd0lpM1NxVDloc2N4M1NNeXJsdCtyN2pTR3ZpVStPQmZXYVc4YTZwSldtdEVXZAoyRkZJVDZaNFNBeVcxcklQMG8rQzBwRWRXM3VrRkpxK0hRK3JKVFFMTFFHVm5DTldRamNkeEtKT3FrcnlsditICkFXNXZnODZGWlhDVkNlZ1p1d2lmQWx5TXFkd2Rjd3VpcW5KVUFGME9xVHFFYnhNSEhnMD0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb2dJQkFBS0NBUUVBeFZNTEZOclczYjlOSXVIbjkyN1V4bWkxdmNaR2RmSXIrM3VvV2YzSUFNS2d2dGt4CmttWmxjcFBUT3NwaGVmSXpBQkZIM0NKZS9Rblp2STV2ME1ZYW8wTDl5SUtST3BKSVQzdjRlcXRnRitMcno3STQKMm1PaG1EbzhQZUo2Q3pIK2hTRVdRY1hua0RyK2w2ZXpIZlBXaFhxRXIzaEVxMk13Ymt6U1lxanRwWTBZOStTVwo1Q21JU0VLNGgrY3NPeUZmRXBvc08xSU82UGVHdUR1SFl1WVVWNzVFaks1em5LQVJkb2loNlByQnlNMDVTRzZuCjFVNlpIcFk3bVRrcStDRUoxNjU2WDA2WmhoVnJWblgyNTJNZi81UmZrMXBqU1hjc091enRHQWljSjJaek5YTFEKazVYbDZHMnA3elJpTWZySEJXMmpxNk10SjhKVHpCN0phc2Q5SndJREFRQUJBb0lCQUU4M1pXMTc1V0trV0EwMgo3KzhQbVhBRnZXQndadjBXdWIxK0NUb0hmZkdBTVJRdEVZK2FlQU9sZ05sTFFoSzR2dXk2QTBWR3J5ZWFlc1VOCjBhbll2ZnpvK2dVekZhYVQ3MStwZnptUDcwWG5uWStHRnZqbG9vd0FaUXJiRHUvTHBFaEIzak9OaGNjTFNBWU8KRndSaFRhL01YZFFyempXWDNtdUpmN1NINk0vYXZzTVNMcm43WjhtUS9pK1JlanRXY3hEZmlaVCtRSUlDK0FPVQpZcWhuR1VLRE8xWWZoUTRmUXNYNDd1NS90bStVbnpuY2VIeExGdTNDVk5TNnczRnhBZXFPeFM1VndRaGgrYllPClorZzZyaWFaZzB0OGNycEVHNEtuWU1lQ2ZQMlk5UHBkTjVWbTdZZlp0bWpLa0xYK0dUSXpBVUlhdHV4NW5JTmwKWHhRTHBURUNnWUVBNGNXd0dneWFZcmRKeWdQdVl6SFNuNUVhd3VrdWNKclJJWVY0YnZDR0JKZEprWXRiVGxGeAp5S1Frb1dqejVYR29sUUlxaDdEcHYxZkZPSGVyR001bnYrWlYwL0U0MERrVm56TC9iWWxGdzdhRnpzYk0yV0pZCmFXZEUzRGprblpYQlVwSWRwU1lGODU5SkVKY1RhZWl3d3l3ay9ML0toaStGWXRzdmdGaUJXdXNDZ1lFQTM3NVAKVTByd0dhQ3dUdStOTVB3NVFvN1hjZjYyNndqYlpTZHVyS3NHeC9KMnlHL3R2d09RV3daUVFxdXJnVXJwYW96eQo1UnlZTWNtSFZISDgvaXpmT1NJZ1pKckd1YkU2cytrQTBsdGljTmVxUkxNWjhVblFpR3RJU2JUY1RFRG5veHM1CndDZ3BkYUtTQUk2Q3d6dHM3MGErOU5SV3FzQzdiR01xN1RSL1g3VUNnWUEyT3dRSDNjc2Z0eU1VVVVscnJrWUgKYWhWaGlCMU5rd0owNk5oNjNXOXpudHRmQ0hoUUlhUVJLOHhZc1JzVW0rNkFqRnFtNlVuY1dqclhTM2RmcUFTbgp4YTRNRUw4eTFPTnFzQmRHdWxoMW9Gd1h4UXpqa29ubUY1WWt4ODJ5Ukl5QlJ2T3ovYVFrVnJoNE1iSEtHTWlNCnVRZlJaa2hCWHh6TkdCVWE0U2VCTFFLQmdFbWxsdTdQeCtCbnFDRVRjT0lpNDZZbzVubTdZZkpUWkFRVHlyWkQKUldRalJ6NEt1Yk1hTlRZQkNnSW9CN2Z6TkluQ0EyR3UyOW5uZ0FnbnpTTE5HbHp3QXNHYXdMYjJ3MS9jM2t4ZgprRE9jaVlzN2VOcVhkWEN4LzRWalp2QWluUnh2SmI4K2VRY2pqL05tOVZ2Vi83RnpFLy82dE54WHZGbWMrdEJCCmEzdlJBb0dBQ1AyQjcwOFdzanMwZHo0ZkRuQVppMUo5NmI5RU56azU5aDdkYW1MZ3J6U2dMY1QxaGFvZzJ3RXQKVVpZQjNVQytGd1FGelQreG5QZGYrZG9yY1BSVVg4d0lYUzhWbWg0QmhTMVBKalJNSGtsaVQvNHNHdlA2ODBiQgpSRU40UDdXT1AvMkJwc21LcjQxUTc5RDN2MjhaM1lCTFdsSGdmVStiTUduOUZvS2RBbjg9Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
`))
			return
		}
		w.WriteHeader(404)
		return
	}))
	return ts
}
