# repl
通过给定的正则表达式和字符串，对标准输入的内容进行替换。  
说明：repl regexp string  
  或：repl regexp  
  
示例一：  
  标准输入：Hello,123  
  repl [0-9] A  
  标准输出：Hello,AAA  
示例二：  
  标准输入：Hello,123  
  repl [0-9]  
  标准输出：Hello,  
