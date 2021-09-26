package main

import (
  "crypto/aes"
  "encoding/hex"
  "fmt"
)

const (
	KEY = "thisis32bitlongpassphraseimusing"
)

type Message struct {
        Email    string `json:"email"`
        Username string `json:"username"`
        Message  string `json:"message"`
}

func EncryptAES(plaintext string) string {             // шифрование сообщений для хранения в бд
  c, _ := aes.NewCipher([]byte(KEY))
  out := make([]byte, len(plaintext))
  c.Encrypt(out, []byte(plaintext))

  return hex.EncodeToString(out)
}

func DecryptAES(ct string) string {             // расшифровка сообщений из бд
  ciphertext, _ := hex.DecodeString(ct)
  c, _ := aes.NewCipher([]byte(KEY))
  pt := make([]byte, len(ciphertext))
  c.Decrypt(pt, ciphertext)

  return string(pt[:])
}

func create_chat(first_user_id int, second_user_id int) {
  db := connect_db()
  defer db.Close()

  res, _ := db.Query(fmt.Sprintf("INSERT INTO chats (first_user_id, second_user_id) VALUES (%s, %s)", first_user_id, second_user_id))
  defer res.Close()
}

func add_message(body string, chat_id int, user_id int) {
  db := connect_db()
  defer db.Close()

  res, _ := db.Query(fmt.Sprintf("INSERT INTO messages (body, chat_id, user_id) VALUES (\"%s\", %s, %s)", EncryptAES(body), chat_id, user_id))
  defer res.Close()
}

func get_messages_from_chat(chat_id int) {

}
