## goroutine 顺序控制
- sync的unlock 和 lock 控制，unlock在前，lock在后
   
- 同步原语，done <-1 在前，<-done 在后