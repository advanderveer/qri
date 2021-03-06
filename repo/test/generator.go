package test

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	cfgtest "github.com/qri-io/qri/config/test"
	"github.com/qri-io/qri/repo/gen"
)

// NewTestCrypto returns a mocked cryptographic generator for tests
func NewTestCrypto() gen.CryptoGenerator {
	return &testCryptoGenerator{}
}

type testCryptoGenerator struct {
	count int
}

var defaultCryptoGenerator gen.CryptoGenerator

func init() {
	defaultCryptoGenerator = NewTestCrypto()
}

func (g *testCryptoGenerator) GeneratePrivateKeyAndPeerID() (string, string) {
	info := cfgtest.GetTestPeerInfo(g.count)
	g.count++
	return info.EncodedPrivKey, info.EncodedPeerID
}

func (g *testCryptoGenerator) GenerateNickname(peerID string) string {
	return "testnick"
}

func (g *testCryptoGenerator) GenerateEmptyIpfsRepo(repoPath, configPath string) error {
	unzipFile(TestdataPath("empty_ipfs_repo.zip"), repoPath)
	return nil
}

func unzipFile(sourceZip, destDir string) {
	r, err := zip.OpenReader(sourceZip)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			panic(err)
		}
		defer rc.Close()

		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				panic(err)
			}
			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				panic(err)
			}
			_, err = io.Copy(outFile, rc)
			outFile.Close()
		}
	}
}
