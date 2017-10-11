package chain

import (
	"errors"
	"fmt"
	"gochain/config"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Chain : the shape of what is returned
// should have a name corresponding to the file it's stored in
// and the actual chain consisting of links
type Chain struct {
	Name  string
	Chain []link
}

type link struct {
	Date   string
	Symbol rune
}

func (c *Chain) PrintChain() {
	x := config.Configuration{}
	x.GetConfiguration()

	c.Name = "My Chain"
	fmt.Println(c.Name)
}

// GetChain : returns a chain with a given name
func GetChain(name string) (chain Chain) {
	return Chain{}
}

// CreateChain : create a chain with a given name
// will create a chain directory if none exists
// directory will default to user's home directory
// or if the ~/gochain.json exists will use the
// defaultDirectory value from there
func CreateChain(name string) error {
	chain := Chain{}
	chain.Name = name
	createChainDir()
	_, err := createChain(name)

	return err
}

func createChainDir() error {
	config := config.Configuration{}
	config.GetConfiguration()

	dirPath := filepath.Join(config.ChainDirectory, "Chains")
	if exist, err := objExists(dirPath); !exist {
		if err == nil {
			os.MkdirAll(dirPath, os.ModePerm)
		} else {
			return err
		}
	} else {
		return errors.New("directory creation failed")
	}

	return nil
}

func objExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func createChain(name string) (bool, error) {
	config := config.Configuration{}
	config.GetConfiguration()
	chainPath := filepath.Join(config.ChainDirectory, "Chains", name+".chain")
	if exist, existErr := objExists(chainPath); !exist {
		if existErr == nil {
			newChainFile, err := os.Create(chainPath)
			if err != nil {
				return false, err
			}
			newChainFile.Close()
		} else {
			return false, existErr
		}
	} else {
		return false, errors.New("a chain with name already exists")
	}

	return true, nil
}

// GetAllChains : returns list of all chains in chain dir
func GetAllChains() (list []string, readError error) {
	config := config.Configuration{}
	config.GetConfiguration()
	chainPath := filepath.Join(config.ChainDirectory, "Chains")
	files, err := ioutil.ReadDir(chainPath)
	if err != nil {
		return nil, err
	}
	fileNames := make([]string, len(files), len(files))
	for i, f := range files {
		basename := f.Name()
		name := strings.TrimSuffix(basename, filepath.Ext(basename))
		fileNames[i] = name
	}

	return fileNames, nil
}
