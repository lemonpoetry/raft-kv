
源代码在master分支中，请在分支切换中进行master选定

shardkv架构图
![](https://github.com/lemonpoetry/raft-kv/blob/master/shardkv%E6%A1%86%E6%9E%B6%E5%9B%BE.png)

框架测试结果：（-race检测多线程中访问数据时的数据竞态，一般出现在漏锁上，框架搭建初期经常忘）
![image](https://github.com/user-attachments/assets/f64dbae7-15bd-42bf-a068-54e05c3fb1c4)





进行30个并发任务，进行100轮测试，检验框架的正确性
![image](https://github.com/user-attachments/assets/89809bd6-c24b-4a9f-8e63-87fd8f9fe13e)


