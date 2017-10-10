package chain

import (
	"fmt"
	"gochain/config"
	"log"
	"os"
	"path/filepath"
)

type Chain struct {
	Name  string
	Chain []string
}

func (c *Chain) PrintChain() {
	x := config.Configuration{}
	x.GetConfiguration()

	c.Name = "My Chain"
	fmt.Println(c.Name)
}

func GetChain(name string) (chain Chain) {
	return Chain{}
}

func CreateChain(name string) {
	chain := Chain{}
	chain.Name = name
	createChainDir()
	createChain(name)

	return
}

func createChainDir() {
	config := config.Configuration{}
	config.GetConfiguration()

	dirPath := filepath.Join(config.ChainDirectory, "Chains")
	if exist, _ := objExists(dirPath); !exist {
		os.MkdirAll(dirPath, os.ModePerm)
	}
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
	if exist, _ := objExists(chainPath); !exist {
		newChainFile, err := os.Create(chainPath)
		if err != nil {
			log.Fatal(err)
		}
		newChainFile.Close()
	}

	return true, nil
}
