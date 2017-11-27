	"gopkg.in/src-d/go-git-fixtures.v3"
func (s *SuiteCommit) TestParent(c *C) {
	commit, err := s.Commit.Parent(1)
	c.Assert(err, IsNil)
	c.Assert(commit.Hash.String(), Equals, "a5b8b09e2f8fcb0bb99d3ccb0958157b40890d69")
}

func (s *SuiteCommit) TestParentNotFound(c *C) {
	commit, err := s.Commit.Parent(42)
	c.Assert(err, Equals, ErrParentNotFound)
	c.Assert(commit, IsNil)
}

	c.Assert(err, IsNil)

func (s *SuiteCommit) TestStat(c *C) {
	aCommit := s.commit(c, plumbing.NewHash("6ecf0ef2c2dffb796033e5a02219af86ec6584e5"))
	fileStats, err := aCommit.Stats()
	c.Assert(err, IsNil)

	c.Assert(fileStats[0].Name, Equals, "vendor/foo.go")
	c.Assert(fileStats[0].Addition, Equals, 7)
	c.Assert(fileStats[0].Deletion, Equals, 0)
	c.Assert(fileStats[0].String(), Equals, " vendor/foo.go | 7 +++++++\n")

	// Stats for another commit.
	aCommit = s.commit(c, plumbing.NewHash("918c48b83bd081e863dbe1b80f8998f058cd8294"))
	fileStats, err = aCommit.Stats()
	c.Assert(err, IsNil)

	c.Assert(fileStats[0].Name, Equals, "go/example.go")
	c.Assert(fileStats[0].Addition, Equals, 142)
	c.Assert(fileStats[0].Deletion, Equals, 0)
	c.Assert(fileStats[0].String(), Equals, " go/example.go | 142 +++++++++++++++++++++++++++++++++++++++++++++++++++++\n")

	c.Assert(fileStats[1].Name, Equals, "php/crappy.php")
	c.Assert(fileStats[1].Addition, Equals, 259)
	c.Assert(fileStats[1].Deletion, Equals, 0)
	c.Assert(fileStats[1].String(), Equals, " php/crappy.php | 259 ++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
}

func (s *SuiteCommit) TestVerify(c *C) {
	ts := time.Unix(1511197315, 0)
	loc, _ := time.LoadLocation("Asia/Kolkata")
	commit := &Commit{
		Hash:      plumbing.NewHash("8a9cea36fe052711fbc42b86e1f99a4fa0065deb"),
		Author:    Signature{Name: "Sunny", Email: "me@darkowlzz.space", When: ts.In(loc)},
		Committer: Signature{Name: "Sunny", Email: "me@darkowlzz.space", When: ts.In(loc)},
		Message: `status: simplify template command selection
`,
		TreeHash:     plumbing.NewHash("6572ba6df4f1fb323c8aaa24ce07bca0648b161e"),
		ParentHashes: []plumbing.Hash{plumbing.NewHash("ede5f57ea1280a0065beec96d3e1a3453d010dbd")},
		PGPSignature: `
-----BEGIN PGP SIGNATURE-----

iQFHBAABCAAxFiEEoRt6IzxHaZkkUslhQyLeMqcmyU4FAloTCrsTHG1lQGRhcmtv
d2x6ei5zcGFjZQAKCRBDIt4ypybJTul5CADmVxB4kqlqRZ9fAcSU5LKva3GRXx0+
leX6vbzoyQztSWYgl7zALh4kB3a3t2C9EnnM6uehlgaORNigyMArCSY1ivWVviCT
BvldSVi8f8OvnqwbWX0I/5a8KmItthDf5WqZRFjhcRlY1AK5Bo2hUGVRq71euf8F
rE6wNhDoyBCEpftXuXbq8duD7D6qJ7QiOS4m5+ej1UCssS2WQ60yta7q57odduHY
+txqTKI8MQUpBgoTqh+V4lOkwQQxLiz7hIQ/ZYLUcnp6fan7/kY/G7YoLt9pOG1Y
vLzAWdidLH2P+EUOqlNMuVScHYWD1FZB0/L5LJ8no5pTowQd2Z+Nggxl
=0uC8
-----END PGP SIGNATURE-----
`,
	}

	armoredKeyRing := `
-----BEGIN PGP PUBLIC KEY BLOCK-----

mQENBFmtHgABCADnfThM7q8D4pgUub9jMppSpgFh3ev84g3Csc3yQUlszEOVgXmu
YiSWP1oAiWFQ8ahCydh3LT8TnEB2QvoRNiExUI5XlXFwVfKW3cpDu8gdhtufs90Q
NvpaHOgTqRf/texGEKwXi6fvS47fpyaQ9BKNdN52LeaaHzDDZkVsAFmroE+7MMvj
P4Mq8qDn2WcWnX9zheQKYrX6Cs48Tx80eehHor4f/XnuaP8DLmPQx7URdJ0Igckh
N+i91Qv2ujin8zxUwhkfus66EZS9lQ4qR9iVHs4WHOs3j7whsejd4VhajonilVHj
uqTtqHmpN/4njbIKb8q8uQkS26VQYoSYm2UvABEBAAG0GlN1bm55IDxtZUBkYXJr
b3dsenouc3BhY2U+iQFUBBMBCAA+FiEEoRt6IzxHaZkkUslhQyLeMqcmyU4FAlmt
HgACGwMFCQPCZwAFCwkIBwIGFQgJCgsCBBYCAwECHgECF4AACgkQQyLeMqcmyU7V
nAf+J5BYu26B2i+iwctOzDRFcPwCLka9cBwe5wcDvoF2qL8QRo8NPWBBH4zWHa/k
BthtGo1b89a53I2hnTwTQ0NOtAUNV+Vvu6nOHJd9Segsx3E1nM43bd2bUfGJ1eeO
jDOlOvtP4ozuV6Ej+0Ln2ouMOc87yAwbAzTfQ9axU6CKUbqy0/t2dW1jdKntGH+t
VPeFxJHL2gXjP89skCSPYA7yKqqyJRPFvC+7rde1OLdCmZi4VwghUiNbh3s1+xM3
gfr2ahsRDTN2SQzwuHu4y1EgZgPtuWfRxzHqduoRoSgfOfFr9H9Il3UMHf2Etleu
rif40YZJhge6STwsIycGh4wOiLkBDQRZrR4AAQgArpUvPdGC/W9X4AuZXrXEShvx
TqM4K2Jk9n0j+ABx87k9fm48qgtae7+TayMbb0i7kcbgnjltKbauTbyRbju/EJvN
CdIw76IPpjy6jUM37wG2QGLFo6Ku3x8/ZpNGGOZ8KMU258/EBqDlJQ/4g4kJ8D+m
9yOH0r6/Xpe/jOY2V8Jo9pdFTm+8eAsSyZF0Cl7drz603Pymq1IS2wrwQbdxQA/w
B75pQ5es7X34Ac7/9UZCwCPmZDAldnjHyw5dZgZe8XLrG84BIfbG0Hj8PjrFdF1D
Czt9bk+PbYAnLORW2oX1oedxVrNFo5UrbWgBSjA1ppbGFjwSDHFlyjuEuxqyFwAR
AQABiQE8BBgBCAAmFiEEoRt6IzxHaZkkUslhQyLeMqcmyU4FAlmtHgACGwwFCQPC
ZwAACgkQQyLeMqcmyU7ZBggArzc8UUVSjde987Vqnu/S5Cv8Qhz+UB7gAFyTW2iF
VYvB86r30H/NnfjvjCVkBE6FHCNHoxWVyDWmuxKviB7nkReHuwqniQHPgdJDcTKC
tBboeX2IYBLJbEvEJuz5NSvnvFuYkIpZHqySFaqdl/qu9XcmoPL5AmIzIFOeiNty
qT0ldkf3ru6yQQDDqBDpkfz4AzkpFnLYL59z6IbJDK2Hz7aKeSEeVOGiZLCjIZZV
uISZThYqh5zUkvF346OHLDqfDdgQ4RZriqd/DTtRJPlz2uL0QcEIjJuYCkG0UWgl
sYyf9RfOnw/KUFAQbdtvLx3ikODQC+D3KBtuKI9ISHQfgw==
=FPev
-----END PGP PUBLIC KEY BLOCK-----
`

	e, err := commit.Verify(armoredKeyRing)
	c.Assert(err, IsNil)

	_, ok := e.Identities["Sunny <me@darkowlzz.space>"]
	c.Assert(ok, Equals, true)
}