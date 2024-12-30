package bst_err_msg

import "fmt"

func KeyAlreadyExist(key string) string {
	return fmt.Sprintf("Key is already exist in BST: %s", key)
}

func KeyDoNotExist(key string) string {
	return fmt.Sprintf("Key do not exist in BST: %s", key)
}
