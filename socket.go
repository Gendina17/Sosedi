package main

import (
  "log"
  "net/http"
)

func handleConnections(w http.ResponseWriter, r *http.Request) {

        ws, err := upgrader.Upgrade(w, r, nil)   // Обновите начальный запрос GET до веб-сокета
        if err != nil {
                log.Fatal(err)
        }

        defer ws.Close()

         clients[ws] = true

         for {                                 // Бесконечный цикл, который считывает новое сообщение в джейсоне, превращает в структуру и отправляет в канал
                var msg Message

                err := ws.ReadJSON(&msg)
                if err != nil {
                        log.Printf("error: %v", err)  // Если ошибка, то удаляем пользователя из подписчиков (пока так)
                        delete(clients, ws)
                        break
                }

                broadcast <- msg
        }

      }

func handleMessages() {                    // По сути приемник: тот читает сообщения из клиента и передает сюда который отправляет всем
  for {
    msg := <-broadcast
                                          //  В бесчконечном цикле берем сообщение из канала, переводим в джеймон и рассылаем всем пока что клиентам
    for client := range clients {

      err := client.WriteJSON(msg)

      if err != nil {
        log.Printf("error: %v", err)
        client.Close()
        delete(clients, client)
      }

    }
  }
}
