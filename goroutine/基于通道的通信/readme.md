## 基于通道的通信
- 先 done <- true 或 close(done)，后 <-done，可保证前者们在前执行，后者后执行
- 先 <-done1，后 done <- true，可保证前者在前执行，后者后执行