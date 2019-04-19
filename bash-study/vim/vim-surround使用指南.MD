普通模式
--------
|    | 命令                   | 说明 + 示例                                                                          |
|---:|----------------------|:---------------------------------------------------------------------------------|
|    | ds                   | 删除括号                                                                             |
|  例 | `ds` `"`             | `"Hello world!"` =><br> `Hello world!`                                           |
|    | cs                   | 替换括号                                                                             |
|  例 | `cs` `"(`            | `"Hello world!"` =><br> `(Hello world!)`                                         |
|    | cS                   | 替换括号，括号内文本做新一行                                                                   |
|  例 | `cS` `"{`            | `"Hello world!"` =><br> `{` <br> &nbsp;&nbsp;&nbsp;&nbsp;`Hello world!` <br> `}` |
|    | ys                   | 添加括号(配合vim光标移动)                                                                  |
|  例 | `ys` `w` `[`         | `Hello world!` =><br> `[Hello] world!`                                           |
|  例 | `ys` `w` `<em` Enter | `Hello world!` =><br> `<em>Hello<em> world!`                                     |
|    | yS                   | 添加括号，括号内文本做新一行                                                                   |
|  例 | `yS` `w` `[`         | `Hello world!` =><br> `[` <br> &nbsp;&nbsp;&nbsp;&nbsp; `Hello` <br> `] world!`  |
|    | yss                  | 整行括起来                                                                            |
|  例 | `yss` `(`            | `Hello world!` =><br> `( Hello world! )`                                         |
|    | ySS                  | 整行括起来，括号内文本做新一行                                                                  |
|  例 | `ySS` `{`            | `Hello world!` =><br> `{` <br> &nbsp;&nbsp; `Hello world! ` <br> `}`             |
|  例 | `ySS` `<div` Enter   | `Hello world!` =><br> `<div>` <br> &nbsp;&nbsp; `Hello world! ` <br> `<div>`     |
|    | ySs                  | 与ySs功能相同                                                                         |


可视模式
--------
|    | 命令             | 说明 + 示例                                                           |
|---:|:---------------|:----------------------------------------------------------------------|
|    | S              | 选中的括起来                                                              |
|  例 | 选中world: `S(`  | `Hello world!` =><br> `Hello (world)!`                                |
|    | gS             | 选中的括起来，括号内文本做新一行                                             |
|  例 | 选中world: `gS{` | `Hello world!` =><br> `Hello {` <br> &nbsp;&nbsp; ` world` <br> `}! ` |


* copy  fron [vim-surround使用指南.MD](https://gist.github.com/wilon/ac1fc66f4a79e7b0c161c80877c75c94)
